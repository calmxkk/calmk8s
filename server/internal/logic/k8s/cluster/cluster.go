package cluster

import (
	"calmk8s/internal/model/input/k8sin"
	"context"
)

type sCluster struct{}

func NewCluster() *sCluster {
	return &sCluster{}
}

func (s *sCluster) List(ctx context.Context) (list []*k8sin.Cluster, err error) {
	return
}

func (s *sCluster) Create(ctx context.Context, in *k8sin.Cluster) (err error) {
	return
}
