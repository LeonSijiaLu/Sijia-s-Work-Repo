package main

import (
	"context"
	"grpc_test"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main(){
	listener, err := net.Listen("tcp", ":8882")
	if err != nil{
		panic(err)
	}
	srv := grpc.NewServer()
	grpc_test.RegisterAddServiceServer(srv, &server{}) // register service on that server
	reflection.Register(srv)
	if e := srv.Serve(listener); e != nil{
		panic(e)
	}
}

func (s *server) Add(ctx context.Context, request *grpc_test.Request) (*grpc_test.Response, error){
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &grpc_test.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *grpc_test.Request) (*grpc_test.Response, error){
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &grpc_test.Response{Result: result}, nil
}