package grpc

import (
	"context"
	"io"
	"log"

	"github.com/kristoiv/wtf"
	"github.com/kristoiv/wtf/models"

	"google.golang.org/grpc"
)

type TodoListServiceHandler struct {
	TodoListService wtf.TodoListService
}

func (h *TodoListServiceHandler) Add(ctx context.Context, r *models.AddRequest) (*models.AddReturns, error) {
	item, err := h.TodoListService.Add(r.GetTitle())
	if err != nil {
		return nil, err
	}
	log.Printf("Added todo item %s: %q\n", item.ID, item.Title)
	retItem, err := models.MarshalItem(item)
	if err != nil {
		return nil, err
	}
	return &models.AddReturns{Item: retItem}, nil
}

func (h *TodoListServiceHandler) SetChecked(ctx context.Context, r *models.SetCheckedRequest) (*models.SetCheckedResponse, error) {
	log.Printf("Setting checked on todo item %s to: %t\n", r.GetId(), r.GetChecked())
	if err := h.TodoListService.SetChecked(wtf.ItemID(r.GetId()), r.GetChecked()); err != nil {
		return nil, err
	}
	return &models.SetCheckedResponse{}, nil
}

func (h *TodoListServiceHandler) Remove(ctx context.Context, r *models.RemoveRequest) (*models.RemoveResponse, error) {
	log.Printf("Removing todo item %s\n", r.GetId())
	if err := h.TodoListService.Remove(wtf.ItemID(r.GetId())); err != nil {
		return nil, err
	}
	return &models.RemoveResponse{}, nil
}

func (h *TodoListServiceHandler) Items(r *models.ItemsRequest, s models.Grpc_ItemsServer) error {
	log.Println("Listing all items")
	items, err := h.TodoListService.Items()
	if err != nil {
		return err
	}
	for _, item := range items {
		out, err := models.MarshalItem(&item)
		if err != nil {
			return err
		}
		if err := s.Send(&models.ItemStreamReturns{Item: out}); err != nil {
			return err
		}
	}
	return nil
}

var _ wtf.TodoListService = &TodoListService{}

type TodoListService struct {
	Addr *string
	conn *grpc.ClientConn
}

func (s *TodoListService) Add(title string) (*wtf.Item, error) {
	conn, err := s.dial()
	if err != nil {
		return nil, err
	}
	client := models.NewGrpcClient(conn)
	res, err := client.Add(context.Background(), &models.AddRequest{Title: title})
	if err != nil {
		return nil, err
	}
	item := wtf.Item{}
	err = models.UnmarshalItem(res.GetItem(), &item)
	return &item, err
}

func (s *TodoListService) SetChecked(id wtf.ItemID, checked bool) error {
	conn, err := s.dial()
	if err != nil {
		return err
	}
	client := models.NewGrpcClient(conn)
	if _, err := client.SetChecked(context.Background(), &models.SetCheckedRequest{Id: string(id), Checked: checked}); err != nil {
		return err
	}
	return nil
}

func (s *TodoListService) Remove(id wtf.ItemID) error {
	conn, err := s.dial()
	if err != nil {
		return err
	}
	client := models.NewGrpcClient(conn)
	if _, err := client.Remove(context.Background(), &models.RemoveRequest{Id: string(id)}); err != nil {
		return err
	}
	return nil
}

func (s *TodoListService) Items() ([]wtf.Item, error) {
	conn, err := s.dial()
	if err != nil {
		return nil, err
	}

	client := models.NewGrpcClient(conn)
	stream, err := client.Items(context.Background(), &models.ItemsRequest{Index: 0, Count: -1})
	if err != nil {
		return nil, err
	}

	items := []wtf.Item{}
	var item models.ItemStreamReturns
	for {
		if err := stream.RecvMsg(&item); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		wtfItem := wtf.Item{}
		if err := models.UnmarshalItem(item.GetItem(), &wtfItem); err != nil {
			return nil, err
		}

		items = append(items, wtfItem)
	}

	return items, nil
}

func (s *TodoListService) dial() (*grpc.ClientConn, error) {
	if s.conn == nil {
		conn, err := grpc.Dial(*s.Addr, grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		s.conn = conn
	}
	return s.conn, nil
}
