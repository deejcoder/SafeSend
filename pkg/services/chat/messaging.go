package chat

import (
	"SafeSend/pkg/services/chat/messagingpb"
	"context"
	"google.golang.org/grpc"
)

type ChatService struct {
	Server *grpc.Server
}

func (mp *ChatService) Start() error {
	//lis, err := net.Listen("tcp", "0.0.0.0:8000")
	//if err != nil {
	//	return err
	//}
	//
	//mp.Server = grpc.NewServer()
	//messagingpb.RegisterMessageServiceServer(mp.Server, &ChatService{})
	//if err := mp.Server.Serve(lis); err != nil {
	//	return err
	//}

	return nil
}

func (s *ChatService) Stop() {

}

func (mp *ChatService) SendMessage(ctx context.Context, message *messagingpb.Message) (*messagingpb.AckMessage, error) {
	//TODO implement me
	panic("implement me")
}
