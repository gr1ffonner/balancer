package grpc

import (
	"context"

	proto "balancer/api/protobuf"
)

func (s *server) Check(ctx context.Context, req *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
	return &proto.HealthCheckResponse{
		Status: proto.HealthCheckResponse_SERVING,
	}, nil
}

func (s *server) Watch(req *proto.HealthCheckRequest, stream proto.Health_WatchServer) error {
	return nil
}
