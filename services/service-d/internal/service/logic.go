package service

import (
	"context"
	commonpb "grpc-vs-rest-poc/proto"
)

type Processor struct {
	ServiceName string
}

func (p *Processor) Process(ctx context.Context, req *commonpb.Request) (*commonpb.Response, error) {
	return &commonpb.Response{
		TraceId: req.TraceId,
		Service: p.ServiceName,
	}, nil
}
