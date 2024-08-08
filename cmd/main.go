package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/patrikeyeva/microservices_course_chat_server_1.0/internal/server"
	"github.com/patrikeyeva/microservices_course_chat_server_1.0/pkg/chat_server_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("error loading .env file: %v", err)
	}
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", os.Getenv("SERVER_HOST_GRPC"), os.Getenv("SERVER_PORT_GRPC")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chat_server_v1.RegisterChatV1Server(s, &server.Server{})

	log.Printf("server listening at %v", lis.Addr())

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
