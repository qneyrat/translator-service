package server

import (
	"context"
	"translator-service/proto"
	"translator-service/translate"
)

type Endpoints struct {
	TranslateEndpoint TranslateEndpoint
}

type TranslateEndpoint func(ctx context.Context, req *proto.TranslateRequest) (*proto.TranslateResponse, error)

func NewTranslateEndpoint(t translate.Translator) TranslateEndpoint {
	return func(ctx context.Context, req *proto.TranslateRequest) (*proto.TranslateResponse, error) {
		text, err := t.Translate(req.Language.String(), req.Text)
		if err != nil {
			return nil, err
		}
		return &proto.TranslateResponse{Text: text}, nil
	}
}
