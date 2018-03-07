package grpc

import (
	"context"
	"io"

	"github.com/kristoiv/wtf"
	"github.com/kristoiv/wtf/grpc/internal"

	"google.golang.org/grpc"
)

type TodoListServiceHandler struct {
	TodoListService wtf.TodoListService
}

func (h *TodoListServiceHandler) Add(ctx context.Context, r *internal.AddRequest) (*internal.AddReturns, error) {
	item, err := h.TodoListService.Add(r.GetTitle())
	if err != nil {
		return nil, err
	}
	retItem, err := internal.MarshalItem(item)
	if err != nil {
		return nil, err
	}
	return &internal.AddReturns{Item: retItem}, nil
}

func (h *TodoListServiceHandler) SetChecked(ctx context.Context, r *internal.SetCheckedRequest) (*internal.SetCheckedResponse, error) {
	if err := h.TodoListService.SetChecked(wtf.ItemID(r.GetId()), r.GetChecked()); err != nil {
		return nil, err
	}
	return &internal.SetCheckedResponse{}, nil
}

func (h *TodoListServiceHandler) Remove(ctx context.Context, r *internal.RemoveRequest) (*internal.RemoveResponse, error) {
	if err := h.TodoListService.Remove(wtf.ItemID(r.GetId())); err != nil {
		return nil, err
	}
	return &internal.RemoveResponse{}, nil
}

func (h *TodoListServiceHandler) Items(r *internal.ItemsRequest, s internal.Grpc_ItemsServer) error {
	items, err := h.TodoListService.Items()
	if err != nil {
		return err
	}
	for _, item := range items {
		out, err := internal.MarshalItem(&item)
		if err != nil {
			return err
		}
		if err := s.Send(&internal.ItemStreamReturns{Item: out}); err != nil {
			return err
		}
	}
	return nil
}

var _ wtf.TodoListService = &TodoListService{}

type TodoListService struct {
	Addr *string
}

func (s *TodoListService) Add(title string) (*wtf.Item, error) {
	conn, err := s.dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := internal.NewGrpcClient(conn)
	res, err := client.Add(context.Background(), &internal.AddRequest{Title: title})
	if err != nil {
		return nil, err
	}
	item := wtf.Item{}
	err = internal.UnmarshalItem(res.GetItem(), &item)
	return &item, err
}

func (s *TodoListService) SetChecked(id wtf.ItemID, checked bool) error {
	conn, err := s.dial()
	if err != nil {
		return err
	}
	defer conn.Close()
	client := internal.NewGrpcClient(conn)
	if _, err := client.SetChecked(context.Background(), &internal.SetCheckedRequest{Id: string(id), Checked: checked}); err != nil {
		return err
	}
	return nil
}

func (s *TodoListService) Remove(id wtf.ItemID) error {
	conn, err := s.dial()
	if err != nil {
		return err
	}
	defer conn.Close()
	client := internal.NewGrpcClient(conn)
	if _, err := client.Remove(context.Background(), &internal.RemoveRequest{Id: string(id)}); err != nil {
		return err
	}
	return nil
}

func (s *TodoListService) Items() ([]wtf.Item, error) {
	conn, err := s.dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := internal.NewGrpcClient(conn)
	stream, err := client.Items(context.Background(), &internal.ItemsRequest{Index: 0, Count: -1})
	if err != nil {
		return nil, err
	}

	items := []wtf.Item{}
	var item internal.ItemStreamReturns
	for {
		if err := stream.RecvMsg(&item); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		wtfItem := wtf.Item{}
		if err := internal.UnmarshalItem(item.GetItem(), &wtfItem); err != nil {
			return nil, err
		}

		items = append(items, wtfItem)
	}

	return items, nil
}

func (s *TodoListService) dial() (*grpc.ClientConn, error) {
	return grpc.Dial(*s.Addr, grpc.WithInsecure())
}
