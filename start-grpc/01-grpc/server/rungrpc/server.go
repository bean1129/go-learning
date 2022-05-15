package main

import (
	"context"
	"log"
	"net"
	//"fmt"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/status"
	//"google.golang.org/grpc/codes"
	"runtime"
	"time"

	//pb "start-grpc/01-grpc/protobuf/valid_hello"
	pb "start-grpc/01-grpc/protobuf/http_hello"
)

// HelloService 定义我们的服务
type HelloService struct {
	//auth *auth.Authentication
}

//实现服务定义的Hello方法
func (s *HelloService) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// if err := s.auth.Auth(ctx); err != nil {
	// 	return nil, err
	// }

	res := pb.HelloResponse{
		Code:  200,
		Value: "hello " + req.GetValue(),
	}
	return &res, nil
}

// //实现服务定义的Hello方法，超时处理
// func (s *HelloService) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
// 	fmt.Println("recevied hello")
// 	data := make(chan *pb.HelloResponse, 1)
// 	go handle(ctx, req, data)
// 	select {
// 	case res := <-data:
// 		return res, nil
// 	case <-ctx.Done():
// 		return nil, status.Errorf(codes.Canceled, "Client cancelled, abandoning.")
// 	}
// }

func handle(ctx context.Context, req *pb.HelloRequest, data chan<- *pb.HelloResponse) {
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		runtime.Goexit() //超时后退出该Go协程
	case <-time.After(10 * time.Second): // 模拟耗时操作
		res := pb.HelloResponse{
			Code:  200,
			Value: "hello " + req.GetValue(),
		}
		// //修改数据库前进行超时判断
		// if ctx.Err() == context.Canceled{
		// 	...
		// 	//如果已经超时，则退出
		// }
		data <- &res
	}
}

const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	//一元拦截器
	//grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Filter))

	// 在gRPC服务器注册我们的服务
	//pb.RegisterHelloServiceServer(grpcServer, &HelloService{&auth.Authentication{"joker", "123456"}})
	pb.RegisterHelloServiceServer(grpcServer, new(HelloService))

	// //用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
