package main

import (
	"context"
	pb "hellogrpc/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *pb.String,
) (*pb.String, error) {
	reply := &pb.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	// 构造一个grpc服务对象
	grpcServer := grpc.NewServer()
	// 注册grpc服务，和rpc很类似
	pb.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	// 在指定端口提供grpc服务
	grpcServer.Serve(listener)
}
