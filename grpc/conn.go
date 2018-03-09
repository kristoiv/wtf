package grpc

import (
	"log"
	"net"

	"github.com/kristoiv/wtf"
	"github.com/kristoiv/wtf/grpc/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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

	log.Printf("Now listening on %s\n", s.Addr)
	go func() {
		grpcserver := grpc.NewServer()
		reflection.Register(grpcserver)
		internal.RegisterGrpcServer(grpcserver, s.TodoListServiceHandler)
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

var _ wtf.Client = &Client{}

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

func (c *Client) Open() error {
	_, err := c.todoListService.dial()
	return err
}

func (c *Client) Close() error {
	conn, err := c.todoListService.dial()
	if err != nil {
		return err
	}

	err = conn.Close()
	c.todoListService.conn = nil
	return err
}

func (c *Client) TodoListService() wtf.TodoListService {
	return &c.todoListService
}
