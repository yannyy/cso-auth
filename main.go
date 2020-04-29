package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro/v2"

	auth "cso/auth/proto"
)

type Auth struct {
}

type AccessScope int

const (
	AccessScopePC AccessScope = iota
	AccessScopeApp
	AccessScopeWechat
)

func apiLog(method, name string, ts int64) {
	fmt.Printf("%s: %s-%d\n", name, method, ts)
}

func (a *Auth) Message(ctx context.Context, req *auth.Empty, resp *auth.Response) error {
	ts := time.Now().Unix()
	apiLog("start", "Message", ts)
	defer apiLog("end", "Message", ts)
	resp.Message = "Hello, Auth"
	return nil
}

func (a *Auth) Token(ctx context.Context, req *auth.TokenRequest, resp *auth.TokenResponse) error {
	username := req.Username
	token := fmt.Sprintf("token-%s", username)
	resp.Token = token
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
