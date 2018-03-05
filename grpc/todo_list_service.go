package grpc

import (
	"context"
	"io"

	"github.com/kristoiv/wtf"
	"github.com/kristoiv/wtf/data"

	"google.golang.org/grpc"
)

type TodoListServiceHandler struct {
	TodoListService wtf.TodoListService
}

func (h *TodoListServiceHandler) Add(ctx context.Context, r *data.AddRequest) (*data.AddReturns, error) {
	item, err := h.TodoListService.Add(r.GetTitle())
	if err != nil {
		return nil, err
	}
	retItem, err := data.MarshalItem(item)
	if err != nil {
		return nil, err
	}
	return &data.AddReturns{Item: retItem}, nil
}

func (h *TodoListServiceHandler) SetChecked(ctx context.Context, r *data.SetCheckedRequest) (*data.SetCheckedResponse, error) {
	if err := h.TodoListService.SetChecked(wtf.ItemID(r.GetId()), r.GetChecked()); err != nil {
		return nil, err
	}
	return &data.SetCheckedResponse{}, nil
}

func (h *TodoListServiceHandler) Remove(ctx context.Context, r *data.RemoveRequest) (*data.RemoveResponse, error) {
	if err := h.TodoListService.Remove(wtf.ItemID(r.GetId())); err != nil {
		return nil, err
	}
	return &data.RemoveResponse{}, nil
}

func (h *TodoListServiceHandler) Items(r *data.ItemsRequest, s data.Grpc_ItemsServer) error {
	items, err := h.TodoListService.Items()
	if err != nil {
		return err
	}
	for _, item := range items {
		out, err := data.MarshalItem(&item)
		if err != nil {
			return err
		}
		if err := s.Send(&data.ItemStreamReturns{Item: out}); err != nil {
			return err
		}
	}
	return nil
}

type TodoListService struct {
	Addr *string
}

func NewTodoListService() *TodoListService {
	h := &TodoListService{}
	return h
}

func (s *TodoListService) Add(title string) (*wtf.Item, error) {
	conn, err := s.dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := data.NewGrpcClient(conn)
	res, err := client.Add(context.Background(), &data.AddRequest{Title: title})
	if err != nil {
		return nil, err
	}
	item := wtf.Item{}
	err = data.UnmarshalItem(res.GetItem(), &item)
	return &item, err
}

func (s *TodoListService) SetChecked(id wtf.ItemID, checked bool) error {
	conn, err := s.dial()
	if err != nil {
		return err
	}
	defer conn.Close()
	client := data.NewGrpcClient(conn)
	if _, err := client.SetChecked(context.Background(), &data.SetCheckedRequest{Id: string(id), Checked: checked}); err != nil {
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
	client := data.NewGrpcClient(conn)
	if _, err := client.Remove(context.Background(), &data.RemoveRequest{Id: string(id)}); err != nil {
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

	client := data.NewGrpcClient(conn)
	stream, err := client.Items(context.Background(), &data.ItemsRequest{Index: 0, Count: -1})
	if err != nil {
		return nil, err
	}

	items := []wtf.Item{}
	var item data.ItemStreamReturns
	for {
		if err := stream.RecvMsg(&item); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		wtfItem := wtf.Item{}
		if err := data.UnmarshalItem(item.GetItem(), &wtfItem); err != nil {
			return nil, err
		}

		items = append(items, wtfItem)
	}

	return items, nil
}

func (s *TodoListService) dial() (*grpc.ClientConn, error) {
	return grpc.Dial(*s.Addr, grpc.WithInsecure())
}
