package grpc

import (
	"context"
	commonpb "grpc-vs-rest-poc/proto"
	"grpc-vs-rest-poc/services/service-d/internal/service"
)

type Server struct {
	commonpb.UnimplementedProcessorServer
	service.Processor
}

func New(processor *service.Processor) *Server {
	return &Server{
		Processor: *processor,
	}
}

func (s *Server) Process(ctx context.Context, req *commonpb.Request) (*commonpb.Response, error) {
	return s.Processor.Process(ctx, req)
}
