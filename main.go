package main

import (
	"log"
	"net"

	"translator-service/proto"
	"translator-service/server"
	"translator-service/translate"

	"google.golang.org/grpc"
	"os"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	translator := translate.NewGoogleTranslator(os.Getenv("TRANSLATION_API_KEY"))
	srv := server.NewTranslatorServer(server.Endpoints{
		TranslateEndpoint: server.NewTranslateEndpoint(translator),
	})

	s := grpc.NewServer()
	proto.RegisterTranslatorServer(s, srv)
	s.Serve(lis)
}
