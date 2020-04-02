package main

import (
	"context"
	"log"
	"time"

	"github.com/micro/go-micro"

	auth "cso/auth/proto"
)

type Auth struct {
}

func (a *Auth) Message(ctx context.Context, req *auth.Empty, resp *auth.Response) error {
	resp.Message = "Hello, Auth"
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("auth"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()

	auth.RegisterAuthHandler(service.Server(), new(Auth))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
