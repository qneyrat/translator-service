package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"translator-service/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := proto.NewTranslatorClient(conn)
	res, err := client.Translate(
		context.Background(),
		&proto.TranslateRequest{Text:"Salut les astronautes !", Language: proto.Language_en},
	)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Text)
}
