package grpc

import (
	"context"
	"io"

	"github.com/kristoiv/wtf"

	"google.golang.org/grpc"
)

type TodoListServiceHandler struct {
	TodoListService wtf.TodoListService
}

func (h *TodoListServiceHandler) Add(ctx context.Context, r *AddRequest) (*AddReturns, error) {
	item, err := h.TodoListService.Add(r.GetTitle())
	if err != nil {
		return nil, err
	}
	retItem, err := MarshalItem(item)
	if err != nil {
		return nil, err
	}
	return &AddReturns{Item: retItem}, nil
}

func (h *TodoListServiceHandler) SetChecked(ctx context.Context, r *SetCheckedRequest) (*SetCheckedResponse, error) {
	if err := h.TodoListService.SetChecked(wtf.ItemID(r.GetId()), r.GetChecked()); err != nil {
		return nil, err
	}
	return &SetCheckedResponse{}, nil
}

func (h *TodoListServiceHandler) Remove(ctx context.Context, r *RemoveRequest) (*RemoveResponse, error) {
	if err := h.TodoListService.Remove(wtf.ItemID(r.GetId())); err != nil {
		return nil, err
	}
	return &RemoveResponse{}, nil
}

func (h *TodoListServiceHandler) Items(r *ItemsRequest, s Grpc_ItemsServer) error {
	items, err := h.TodoListService.Items()
	if err != nil {
		return err
	}
	for _, item := range items {
		out, err := MarshalItem(&item)
		if err != nil {
			return err
		}
		if err := s.Send(&ItemStreamReturns{Item: out}); err != nil {
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
	client := NewGrpcClient(conn)
	res, err := client.Add(context.Background(), &AddRequest{Title: title})
	if err != nil {
		return nil, err
	}
	item := wtf.Item{}
	err = UnmarshalItem(res.GetItem(), &item)
	return &item, err
}

func (s *TodoListService) SetChecked(id wtf.ItemID, checked bool) error {
	conn, err := s.dial()
	if err != nil {
		return err
	}
	defer conn.Close()
	client := NewGrpcClient(conn)
	if _, err := client.SetChecked(context.Background(), &SetCheckedRequest{Id: string(id), Checked: checked}); err != nil {
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
	client := NewGrpcClient(conn)
	if _, err := client.Remove(context.Background(), &RemoveRequest{Id: string(id)}); err != nil {
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

	client := NewGrpcClient(conn)
	stream, err := client.Items(context.Background(), &ItemsRequest{Index: 0, Count: -1})
	if err != nil {
		return nil, err
	}

	items := []wtf.Item{}
	var item ItemStreamReturns
	for {
		if err := stream.RecvMsg(&item); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		wtfItem := wtf.Item{}
		if err := UnmarshalItem(item.GetItem(), &wtfItem); err != nil {
			return nil, err
		}

		items = append(items, wtfItem)
	}

	return items, nil
}

func (s *TodoListService) dial() (*grpc.ClientConn, error) {
	return grpc.Dial(*s.Addr, grpc.WithInsecure())
}
