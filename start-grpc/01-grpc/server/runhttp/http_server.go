package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"

	pb "start-grpc/01-grpc/protobuf/http_hello"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

const (
	// Address 监听地址
	HttpAddress string = ":8080"
	Address     string = ":8000"
)

func main() {

	fmt.Println("---")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	err := pb.RegisterHelloServiceHandlerFromEndpoint(
		ctx, mux, Address,
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(HttpAddress, mux)
}
