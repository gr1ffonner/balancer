package grpc

import (
	"balancer/internal/config"
	"net"

	proto "balancer/api/protobuf"

	"google.golang.org/grpc"
)

type server struct {
	grpcServer *grpc.Server
	proto.UnimplementedHealthServer
	proto.UnimplementedBalancerServer
	requestCount int32
	cfg          *config.Config
}

func NewServer(cfg *config.Config) *server {
	s := &server{
		grpcServer: grpc.NewServer(),
		cfg:        cfg,
	}

	proto.RegisterHealthServer(s.grpcServer, s)
	proto.RegisterBalancerServer(s.grpcServer, s)

	return s
}

func (s *server) Serve(listener net.Listener) error {
	return s.grpcServer.Serve(listener)
}

func (s *server) GracefulStop() {
	s.grpcServer.GracefulStop()
}
