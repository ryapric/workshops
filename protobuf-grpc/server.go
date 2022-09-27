package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"example.com/pb"
	"google.golang.org/grpc"
)

const addr = "localhost:8080"

type exampleServer struct {
	// You need this according to the protobuf output, so
	pb.UnimplementedExampleServer
}

func (s *exampleServer) Echo(ctx context.Context, req *pb.Echoable) (*pb.Echoable, error) {
	msg := fmt.Sprintf("rpc call to 'Echo', received msg: '%s'", req.Msg)
	log.Printf(msg + " -- responding in kind\n")
	return &pb.Echoable{Msg: req.Msg}, nil
}

func (s *exampleServer) GetRecord(ctx context.Context, req *pb.GetRecordRequest) (*pb.GetRecordResponse, error) {
	log.Printf("Received the following request on 'GetRecord' --> %v", req)
	return &pb.GetRecordResponse{
		Id:       1,
		Name:     "o no u di'int",
		Birthday: "1991-01-01",
		Details:  []string{"a", "b", "c"},
	}, nil
}

func runServer() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s\n", addr)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterExampleServer(grpcServer, &exampleServer{})

	log.Printf("starting server on %s...\n", listen.Addr())
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalln(err)
	}
}
