package server

import (
	"context"
	"translator-service/proto"
)

type TranslatorServer struct {
	translate TranslateEndpoint
}

func NewTranslatorServer(e Endpoints) *TranslatorServer {
	return &TranslatorServer{
		translate: e.TranslateEndpoint,
	}
}

func (s TranslatorServer) Translate(ctx context.Context, req *proto.TranslateRequest) (*proto.TranslateResponse, error) {
	return s.translate(ctx, req)
}
