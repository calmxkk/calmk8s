package pod

import (
	"calmk8s/internal/service"
	"context"
)

type sPod struct{}

func NewPod() *sPod {
	return &sPod{}
}

func init() {
	service.RegisterPod(NewPod())
}

func (s *sPod) List(ctx context.Context) {
	return
}
