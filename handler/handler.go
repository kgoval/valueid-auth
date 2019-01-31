package handler

import (
	"context"
	"github.com/kgoval/valueid-auth/protogen"
	"github.com/micro/go-micro/server"
)

type AuthSvc struct{}

func (c *AuthSvc) Token(ctx context.Context, in *auth.Empty, out *auth.AuthTokenResponse) error {

	return nil
}

func Register(s server.Server) {
	
	auth.RegisterAuthServiceHandler(s, new(AuthSvc))
}
