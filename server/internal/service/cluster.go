// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"calmk8s/internal/model/input/k8sin"
	"context"

	"k8s.io/client-go/informers"
)

type (
	ICluster interface {
		GetInformer(ctx context.Context, name string) (informers.SharedInformerFactory, error)
		Create(ctx context.Context, in *k8sin.Cluster) (err error)
	}
)

var (
	localCluster ICluster
)

func Cluster() ICluster {
	if localCluster == nil {
		panic("implement not found for interface ICluster, forgot register?")
	}
	return localCluster
}

func RegisterCluster(i ICluster) {
	localCluster = i
}
