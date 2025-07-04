package main

import (
	grpcdef "app/go/bp"
	"app/go/grpcsrv"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// 2. gRPCサーバーを作成
	s := grpc.NewServer()

	// gRPCサーバにサービスを登録
	grpcdef.RegisterGreetingServiceServer(s, grpcsrv.New())

	// 3. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	// 4.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
