package grpcserver

import (
	"fmt"
	grpc_auth "github.com/MovingTowardsADream/SneakerStore-UserService/internal/grpc/auth"
	grpc_user "github.com/MovingTowardsADream/SneakerStore-UserService/internal/grpc/user"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

const (
	_defaultPort = ":8080"
)

type Server struct {
	log        *slog.Logger
	gRPCServer *grpc.Server

	port string
}

func New(log *slog.Logger, userProfileInfo grpc_user.UserProfileInfo, auth grpc_auth.Authorization, opts ...Option) *Server {

	s := &Server{
		log:  log,
		port: _defaultPort,
	}

	for _, opt := range opts {
		opt(s)
	}

	gRPCServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			s.loggingInterceptor,
			s.recoveryInterceptor,
		),
	)

	grpc_user.UserProfile(gRPCServer, userProfileInfo)
	grpc_auth.Auth(gRPCServer, auth)

	s.gRPCServer = gRPCServer

	return s
}

func (s *Server) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (s *Server) Run() error {
	const op = "gRPC - Server.Run"

	l, err := net.Listen("tcp", s.port)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.log.Info("gRPC server started", slog.String("addr", l.Addr().String()))

	if err := s.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Server) Shutdown() {
	const op = "gRPC - Server.Shutdown"

	s.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.String("port", s.port))

	s.gRPCServer.GracefulStop()
}
