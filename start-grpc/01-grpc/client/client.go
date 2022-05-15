package main

import (
	//"fmt"
	"time"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "start-grpc/01-grpc/protobuf/hello"
	auth "start-grpc/01-grpc/server/token"
)
const (
	// Address 连接地址
	Address string = ":8000"
)

var grpcClient pb.HelloServiceClient

func callHello(ctx context.Context,req pb.HelloRequest)(*pb.HelloResponse,error){
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	return grpcClient.Hello(ctx, &req)
}

func callHelloWithTimeout(ctx context.Context,req pb.HelloRequest)(*pb.HelloResponse,error){
	clientDeadline := time.Now().Add(time.Duration(2 * time.Second))
	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	res, err := grpcClient.Hello(ctx, &req)
	if err != nil {
		//获取错误状态
		statu, ok := status.FromError(err)
		if ok {
			//判断是否为调用超时
			if statu.Code() == codes.DeadlineExceeded {
				//log.Fatalln("Call Hello timeout!")
			}
		}
		return res,err
	}

	return res,nil
}


func main() {
	// 连接服务器
	auth :=auth.Authentication{
        User:    "joker",
        Password: "123456",
    }

    conn, err := grpc.Dial(Address, grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

	// 建立gRPC连接
	grpcClient = pb.NewHelloServiceClient(conn)
	// 创建发送结构体
	req := pb.HelloRequest{
		Value: "1grpc",
	}

	res,err := callHello(context.Background(),req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}

	// res,err := callHelloWithTimeout(context.Background(),req)
	// if err != nil {
	// 	log.Fatalf("Call Route err: %v", err)
	// }

	// 打印返回值
	log.Println(res)
}