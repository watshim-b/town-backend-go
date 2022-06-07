package handler

import (
	"context"

	hd "github.com/watshim-b/town-backend-go/handler_definition"
)

type loginServer struct{}

func NewLoginServer() *loginServer {
	return &loginServer{}
}

func (s *loginServer) Login(ctx context.Context, req *hd.LoginForm) (*hd.Response, error) {
	return &hd.Response{Message: "Success!"}, nil
}
