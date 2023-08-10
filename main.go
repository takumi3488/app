package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"notice_bot/notice"
	pb "notice_bot/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.NoticeServiceServer
}

func (s *server) Notice(ctx context.Context, req *pb.NoticeRequest) (*pb.NoticeResponse, error) {
	var err error
	switch req.Type {
	case "line":
		err = notice.Line(req.Content)
	}

	// レスポンスを返す
	if err != nil {
		return &pb.NoticeResponse{
			Ok: false,
		}, err
	}
	return &pb.NoticeResponse{
		Ok: true,
	}, nil
}

func NewServer() *server {
	return &server{}
}

func main() {
	port := 8080
	listner, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterNoticeServiceServer(s, NewServer())

	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server port: %v", 8080)
		s.Serve(listner)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
