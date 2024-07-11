package grpc

import (
	"context"
	"log"
	"time"

	pb_msg "github.com/ddr4869/msazoom/user-service/grpc/proto/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RpcClient struct {
	MessageClient pb_msg.MessageServiceClient
}

var RpcClientInstance RpcClient

func NewMessageClient() *grpc.ClientConn {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to message-service: %v", err)
	}
	RpcClientInstance = RpcClient{}
	RpcClientInstance.MessageClient = pb_msg.NewMessageServiceClient(conn)
	return conn
}

func (rpc RpcClient) GetUnreadMessagesCount(sender, receiver string) (int32, error) {
	req := &pb_msg.UnreadMessagesRequest{
		UserId:   sender,
		FriendId: receiver,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := rpc.MessageClient.GetUnreadMessagesCount(ctx, req)
	if err != nil {
		return -1, err
	}

	return res.UnreadCount, nil
}
