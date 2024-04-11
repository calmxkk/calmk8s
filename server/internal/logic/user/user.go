package user

import (
	"calmk8s/internal/model/input"
	"context"
)

type sUser struct{}

func NewUser() *sUser {
	return &sUser{}
}

func (s *sUser) Register(ctx context.Context, in *input.RegisterInp) (err error) {
	return
}
