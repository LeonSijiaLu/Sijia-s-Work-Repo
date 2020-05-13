package main

import (
	"os"
	"net"
	"log"
	"sync"
	"context"
	"ChatApp/proto"
	"google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
)

var grpcLog glog.LoggerV2

type Connection struct {
	id string
	active bool
	error chan error
	stream proto.BroadCast_CreateStreamServer  // allow stream between server and client
}

type Server struct {
	Connection []*Connection
}

func init(){
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

func (s *Server) CreateStream(pconn *proto.Connect, stream proto.BroadCast_CreateStreamServer) error{
	conn := &Connection{
		id: pconn.User.Id,
		active: true,  // set to true by default
		error: make(chan error),
		stream: stream,
	}
	s.Connection = append(s.Connection, conn)
	return <- conn.error
}

func (s *Server) BroadcastMessage(ctx context.Context, msg *proto.Message) (*proto.Close, error) {
	wait := sync.WaitGroup{} // how many remaining go routines we have
	done := make(chan int) // to know when all goroutines are finished
	for _, conn := range s.Connection{
		wait.Add(1)  // increment waitGroup
		go func(msg *proto.Message, conn *Connection){
			defer wait.Done()
			if conn.active{
				err := conn.stream.Send(msg)
				grpcLog.Info("Sending message to: ", conn.stream)
				if err != nil{
					grpcLog.Errorf("Error with Stream: ", err)
					conn.active = false
					conn.error <- err // pass error to error channel
				}
			}
		}(msg, conn)
	}
	go func(){
		wait.Wait()
		close(done)
	}()
	<-done // block return until all goroutines are finished
	return &proto.Close{}, nil
}	

func main(){
	connections := []*Connection{}
	server := &Server{connections}
	grpcServer := grpc.NewServer()  // declare a grpc server
	listener, err := net.Listen("tcp", ":8811") // serve on port 8080
	if err != nil{
		log.Fatalf("Error: ", err)
	}
	grpcLog.Info("Starting server at port :8811")
	proto.RegisterBroadCastServer(grpcServer, server) // Create a Broadcast Server
	grpcServer.Serve(listener)
}
