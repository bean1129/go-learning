package main

import (
	"context"
	"log"
	proto "pubsub/protobuf"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := proto.NewPubsubServiceClient(conn)

	_, err = client.Publish(
		context.Background(), &proto.String{Value: "golang: hello Go"},
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Publish(
		context.Background(), &proto.String{Value: "docker: hello Docker"},
	)
	if nil != err {
		log.Fatal(err)
	}

}
