package auth

import (
	"context"
	"github.com/MovingTowardsADream/SneakerStore-UserService/pkg/user_service_v1"
	"google.golang.org/grpc"
)

type authRoutes struct {
	user_service_v1.UnimplementedAuthServer
}

func Auth(gRPC *grpc.Server) {
	user_service_v1.RegisterAuthServer(gRPC, &authRoutes{})
}

func (ar *authRoutes) SignUp(ctx context.Context, req *user_service_v1.SignUpRequest) (*user_service_v1.SignUpResponse, error) {
	return &user_service_v1.SignUpResponse{UserId: 1}, nil
}

func (ar *authRoutes) SignIn(ctx context.Context, req *user_service_v1.SignInRequest) (*user_service_v1.SignInResponse, error) {
	return &user_service_v1.SignInResponse{Token: "Test"}, nil
}
