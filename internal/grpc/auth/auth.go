package grpc_auth

import (
	"context"
	"github.com/MovingTowardsADream/SneakerStore-UserService/pkg/user_service_v1"
	"google.golang.org/grpc"
)

type authRoutes struct {
	user_service_v1.UnimplementedAuthServer
	auth Authorization
}

func Auth(gRPC *grpc.Server, auth Authorization) {
	user_service_v1.RegisterAuthServer(gRPC, &authRoutes{auth: auth})
}

func (ar *authRoutes) SignUp(ctx context.Context, req *user_service_v1.SignUpRequest) (*user_service_v1.SignUpResponse, error) {
	_, _ = ar.auth.CreateUser(ctx, nil)
	return &user_service_v1.SignUpResponse{UserId: 1}, nil
}

func (ar *authRoutes) SignIn(ctx context.Context, req *user_service_v1.SignInRequest) (*user_service_v1.SignInResponse, error) {
	_, _ = ar.auth.GenerateToken(ctx, nil)
	return &user_service_v1.SignInResponse{Token: "Test"}, nil
}
