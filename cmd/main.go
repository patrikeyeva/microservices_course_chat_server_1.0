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
	if errEnvLoad := godotenv.Load(); errEnvLoad != nil {
		log.Printf("Error loading .env file: %v", errEnvLoad)
	}
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	lis, errListen := net.Listen("tcp", fmt.Sprintf("%v:%v", os.Getenv("SERVER_HOST_GRPC"), os.Getenv("SERVER_PORT_GRPC")))
	if errListen != nil {
		log.Fatalf("failed to listen: %v", errListen)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	chat_server_v1.RegisterChatV1Server(s, &server.Server{})
	log.Printf("server listening at %v", lis.Addr())

	if errServe := s.Serve(lis); errServe != nil {
		log.Fatalf("failed to serve: %v", errServe)
	}
}
