package grpcsrv

import (
	grpcdef "app/go/bp"
	"context"
	"fmt"
)

type MyServer struct {
	grpcdef.UnimplementedGreetingServiceServer
}

func New() *MyServer {
	return &MyServer{}
}

func (s *MyServer) Hello(c context.Context, req *grpcdef.HelloRequest) (*grpcdef.HelloResponse, error) {
	// リクエストからnameフィールドを取り出して
	// "Hello, [名前]!"というレスポンスを返す
	return &grpcdef.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}
