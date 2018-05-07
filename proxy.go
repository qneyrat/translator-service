package main

import (
	"context"
	"net/http"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"translator-service/proto"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	proto.RegisterTranslatorHandlerFromEndpoint(ctx, mux, "localhost:4000", opts)

	http.ListenAndServe(":8001", mux)
}
