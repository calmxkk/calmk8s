package pod

import "context"

type sPod struct{}

func NewPod() *sPod {
	return &sPod{}
}

func (s sPod) List(ctx context.Context)
