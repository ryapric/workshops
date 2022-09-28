package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ryapric/workshops/protobuf-grpc/pb/example/v1"
	"google.golang.org/grpc"
)

const addr = "localhost:8080"

type exampleServiceServer struct {
	// You can make this embedded struct required via a `protoc` Go option,
	// which essentially allows you to NOT fully implement the generated
	// interface (i.e. optionally leave out method definitions). We're removing
	// it here because it makes it more clear when we've NOT implemented the
	// interface easier (i.e. the compiler will complain if any methods are
	// missing) recommended, but including it is considered 'best-practice' at
	// the time of this writing.

	// pb.UnimplementedExampleServer
}

func (s *exampleServiceServer) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	msg := fmt.Sprintf("rpc call to 'Echo', received msg: '%s'", req.Msg)
	log.Printf(msg + " -- responding in kind\n")
	return &pb.EchoResponse{Msg: req.Msg}, nil
}

func (s *exampleServiceServer) GetRecord(ctx context.Context, req *pb.GetRecordRequest) (*pb.GetRecordResponse, error) {
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
	pb.RegisterExampleServiceServer(grpcServer, &exampleServiceServer{})

	log.Printf("starting server on %s...\n", listen.Addr())
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalln(err)
	}
}
