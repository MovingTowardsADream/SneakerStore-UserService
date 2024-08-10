package grpc_user

import (
	"context"
	"github.com/MovingTowardsADream/SneakerStore-UserService/pkg/user_service_v1"
	"google.golang.org/grpc"
)

type userRoutes struct {
	user_service_v1.UnimplementedUserProfileServer
	userProfileInfo UserProfileInfo
}

func UserProfile(gRPC *grpc.Server, userProfileInfo UserProfileInfo) {
	user_service_v1.RegisterUserProfileServer(gRPC, &userRoutes{userProfileInfo: userProfileInfo})
}

func (ur *userRoutes) ProfileInfo(ctx context.Context, req *user_service_v1.ProfileInfoRequest) (*user_service_v1.ProfileInfoResponse, error) {
	_ = ur.userProfileInfo.AddUserProfile(ctx, nil)
	return &user_service_v1.ProfileInfoResponse{Response: "Test"}, nil
}
