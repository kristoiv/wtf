package grpc

import (
	"log"
	"net"

	"github.com/kristoiv/wtf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//go:generate protoc --go_out=plugins=grpc:. grpc.proto

const DefaultAddr = ":3000"

type Server struct {
	ln                     net.Listener
	Addr                   string
	TodoListServiceHandler *TodoListServiceHandler
}

func NewServer() *Server {
	return &Server{
		Addr: DefaultAddr,
	}
}

func (s *Server) Open() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.ln = ln
	go func() {
		grpcserver := grpc.NewServer()
		reflection.Register(grpcserver)
		RegisterGrpcServer(grpcserver, s.TodoListServiceHandler)
		log.Fatalln(grpcserver.Serve(ln))
	}()
	return nil
}

func (s *Server) Close() error {
	if s.ln != nil {
		s.ln.Close()
	}
	return nil
}

func (s *Server) Port() int {
	return s.ln.Addr().(*net.TCPAddr).Port
}

type Client struct {
	Addr            string
	todoListService TodoListService
}

func NewClient() *Client {
	c := &Client{
		Addr: DefaultAddr,
	}
	c.todoListService.Addr = &c.Addr
	return c
}

func (c *Client) TodoListService() wtf.TodoListService {
	return &c.todoListService
}
