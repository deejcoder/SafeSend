package messaging

import (
	"SafeSend/server/messaging/messagingpb"
	"context"
	"google.golang.org/grpc"
	"net"
)

type MessageService struct {
	Server *grpc.Server
}

func (mp *MessageService) Start() error {
	lis, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		return err
	}

	mp.Server = grpc.NewServer()
	messagingpb.RegisterMessageServiceServer(mp.Server, &MessageService{})
	if err := mp.Server.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *MessageService) Stop() {

}

func (mp *MessageService) SendMessage(ctx context.Context, message *messagingpb.Message) (*messagingpb.AckMessage, error) {
	//TODO implement me
	panic("implement me")
}
