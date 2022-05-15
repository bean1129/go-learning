package Interceptor

import (
	"context"
	"fmt"
	"log"

	//pb "start-grpc/01-grpc/protobuf/valid_hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Filter(ctx context.Context,
	req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {

	//这里可以对请求参数做校验
	//fmt.Println(req.(*pb.HelloRequest).Validate())

	log.Println("info:", info, "hander:", handler)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing credentials")
	}

	var appid string
	var appkey string

	if val, ok := md["user"]; ok {
		appid = val[0]
	}
	if val, ok := md["password"]; ok {
		appkey = val[0]
	}

	log.Println(appid, appkey) //这里也可以实现tokern校验

	return handler(ctx, req)
}
