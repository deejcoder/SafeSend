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

const (
	CtxAccessTokenKey    = "AccessToken"
	CtxRefreshTokenKey = "RefreshToken"
	CtxTokenExpiryKey  = "TokenExpiry"
)
func Ping(conn *grpc.ClientConn, p *pb.Request, audience string) (*pb.Response, error) {

	md, ok := grpcMetadata.FromIncomingContext(conn.)

	client := pb.NewAuthenticationClient(conn)
	return client.Ping(outgoingCtx, p)
}

func ValidateToken(ctx context.Context, p *pb.Request) bool {
	md, ok := grpcMetadata.FromIncomingContext(ctx)
	if !ok { return false }

	expiry := md.Get(CtxTokenExpiryKey)[0]
	expiryTime, err := time.Parse(time.RFC3339, expiry)
	if err != nil { return false }

	if expiryTime

	acessToken := md.Get(CtxAccessTokenKey)



}

func RefreshToken(ctx context.Context, p *pb.Request) bool {


}

func RequestToken(conn *grpc.ClientConn, p *pb.Request, audience string) (*pb.Response, error) {
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
		CtxAccessTokenKey, token.AccessToken,
		CtxTokenExpiryKey, token.Expiry.String(),
		CtxRefreshTokenKey, token.RefreshToken)

	outgoingCtx := grpcMetadata.NewOutgoingContext(ctx, meta)

	client := pb.NewAuthenticationClient(conn)
	//return client.RequestToken(outgoingCtx, p)
}
