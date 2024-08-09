package user

import (
	"context"
	"github.com/MovingTowardsADream/SneakerStore-UserService/pkg/user_service_v1"
	"google.golang.org/grpc"
)

type userRoutes struct {
	user_service_v1.UnimplementedUserProfileServer
}

func UserProfile(gRPC *grpc.Server) {
	user_service_v1.RegisterUserProfileServer(gRPC, &userRoutes{})
}

func (ur *userRoutes) ProfileInfo(ctx context.Context, req *user_service_v1.ProfileInfoRequest) (*user_service_v1.ProfileInfoResponse, error) {
	return &user_service_v1.ProfileInfoResponse{Response: "Test"}, nil
}
