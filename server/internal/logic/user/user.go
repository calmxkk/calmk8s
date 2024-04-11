package user

import (
	"context"
	"server/internal/model/input"
)

type sUser struct{}

func NewUser() *sUser {
	return &sUser{}
}

func (s *sUser) Register(ctx context.Context, in *input.RegisterInp) (err error) {
	return
}
