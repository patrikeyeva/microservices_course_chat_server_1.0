package server

import (
	"context"
	"log"

	"github.com/patrikeyeva/microservices_course_chat_server_1.0/pkg/chat_server_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Server implements gRPC chat-server
type Server struct {
	chat_server_v1.UnimplementedChatV1Server
}

// Create creating new chat
func (s *Server) Create(_ context.Context, req *chat_server_v1.CreateRequest) (*chat_server_v1.CreateResponse, error) {
	log.Printf("CREATE:\n %v", req.GetUsernames())
	return &chat_server_v1.CreateResponse{
		Id: 0,
	}, nil
}

// Delete deleting chat by ID
func (s *Server) Delete(_ context.Context, req *chat_server_v1.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("DELETE:\n %v", req.GetId())
	return nil, nil
}

// SendMessage sending message to server
func (s *Server) SendMessage(_ context.Context, req *chat_server_v1.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("SendMessage:\n %v", req.GetMessage())
	return nil, nil
}
