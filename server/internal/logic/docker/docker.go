package docker

import (
	"calmk8s/internal/library/dockerck"
	"context"
)

type sDocker struct {
	client *dockerck.Client
}

func NewDocker() *sDocker {
	return &sDocker{
		client: dockerck.NewClient(),
	}
}

func (s *sDocker) Version(ctx context.Context) {
}

func (s *sDocker) GetImageByName(ctx context.Context, imageName string) {
	s.client.GetImageByName(ctx, imageName)
}
