package grpc

import (
	"balancer/internal/utils"
	"context"
	"sync/atomic"

	proto "balancer/api/protobuf"
)

func (s *server) Redirect(ctx context.Context, req *proto.RedirectRequest) (*proto.RedirectResponse, error) {
	count := atomic.AddInt32(&s.requestCount, 1)

	url := req.GetVideoUrl()

	sN, videoPath, err := utils.ParseURL(url)
	if err != nil {
		return nil, err
	}

	cdnUrl := utils.ConstructCDNURL(s.cfg.CDN_HOST, sN, videoPath)

	if count%10 == 0 {
		return &proto.RedirectResponse{
			RedirectUrl: url,
		}, nil
	}

	return &proto.RedirectResponse{
		RedirectUrl: cdnUrl,
	}, nil
}
