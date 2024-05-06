package docker

import (
	"calmk8s/internal/library/dockerutil"
	"context"
)

type sDocker struct {
	client *dockerutil.Client
}

func NewDocker() *sDocker {
	return &sDocker{
		client: dockerutil.NewClient(),
	}
}

func (s *sDocker) Version(ctx context.Context) {
}

func (s *sDocker) GetImageByName(ctx context.Context, imageName string) {
	s.client.GetImageByName(ctx, imageName)
}
