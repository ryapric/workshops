package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "github.com/ryapric/workshops/protobuf-grpc/pb/example/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runClient(uri string) {
	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(uri, opts...)
	if err != nil {
		log.Fatalf("error dialing grpc: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewExampleServiceClient(conn)

	echoResponse, err := client.Echo(context.TODO(), &pb.EchoRequest{Msg: "hello grpc"})
	if err != nil {
		log.Fatalf("error calling rpc Echo(): %v\n", err)
	}
	fmt.Println("Received Echo back: ", echoResponse.Msg)

	getRecordResponse, err := client.GetRecord(context.TODO(), &pb.GetRecordRequest{Name: "Thomas Anderson"})
	if err != nil {
		log.Fatalf("error calling rpc Echo(): %v\n", err)
	}
	dump, err := json.MarshalIndent(getRecordResponse, "", "  ")
	if err != nil {
		log.Fatalf("could not marshal json from response: %v", err)
	}
	fmt.Println("Got the following JSON record back for the name query:")
	fmt.Println(string(dump))
}

func main() {
	runClient("127.0.0.1:8080")
}
