package grpc

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/ddr4869/msazoom/message-service/config"
	"github.com/ddr4869/msazoom/message-service/internal/repository"
	pb_msg "github.com/ddr4869/msazoom/proto/message"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	pb_msg.UnimplementedMessageServiceServer
	repository repository.Repository
}

func (s *GrpcServer) GetUnreadMessagesCount(ctx context.Context, req *pb_msg.UnreadMessagesRequest) (*pb_msg.UnreadMessagesResponse, error) {
	unreadCount, err := s.repository.GetNumberOfUnreadMessage(ctx, req.UserId, req.FriendId)
	if err != nil {
		return nil, err
	}
	return &pb_msg.UnreadMessagesResponse{
		UnreadCount: int32(unreadCount),
	}, nil
}

func SetGrpcServer(cfg *config.Config) {
	port := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	defer s.Stop()

	repo := repository.Repository{}
	err = repo.NewEntClient(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Name, cfg.DB.Password)
	if err != nil {
		log.Fatal(err)
	}

	pb_msg.RegisterMessageServiceServer(s, &GrpcServer{repository: repo})

	log.Printf("Listening and serving GRPC on %s\n", port)

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
