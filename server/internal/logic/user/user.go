package user

import (
	"calmk8s/internal/model/input"
	"calmk8s/internal/service"
	"context"
)

type sUser struct{}

func NewUser() *sUser {
	return &sUser{}
}

func init() {
	service.RegisterUser(NewUser())
}

func (s *sUser) Register(ctx context.Context, in *input.RegisterInp) (err error) {
	return
}

func (s *sUser) Login(ctx context.Context, in *input.AccountLoginInp) (err error) {
	return
}
