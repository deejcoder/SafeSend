package auth

import (
	"SafeSend/pkg/services/pb"
	"context"
	"fmt"
	"google.golang.org/api/idtoken"
	"google.golang.org/grpc"
	grpcMetadata "google.golang.org/grpc/metadata"
	"time"
)

func Ping(conn *grpc.ClientConn, p *pb.Request, audience string) (*pb.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tokenSource, err := idtoken.NewTokenSource(ctx, audience)
	if err != err {
		return nil, fmt.Errorf("error generating new identity token: %v", err)
	}

	token, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("error obtaining token from source: %v", err)
	}

	meta := grpcMetadata.Pairs(
		"authorization", "Bearer"+token.AccessToken,
		"token_expiry", token.Expiry.String())

	outgoingCtx := grpcMetadata.NewOutgoingContext(ctx, meta)

	client := pb.NewAuthenticationClient(conn)
	return client.Ping(outgoingCtx, p)
}
