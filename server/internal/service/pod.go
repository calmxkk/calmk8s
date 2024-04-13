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
	IPod interface {
		GetPodByName(ctx context.Context, in *k8sin.GetPodInp) (*k8sin.GetPodRes, error)
		List(ctx context.Context, in *k8sin.GetPodInp) ([]*k8sin.GetPodRes, error)
		NamespaceExists(ctx context.Context, k8sInformer informers.SharedInformerFactory, namespaceName string) (exist bool, err error)
	}
)

var (
	localPod IPod
)

func Pod() IPod {
	if localPod == nil {
		panic("implement not found for interface IPod, forgot register?")
	}
	return localPod
}

func RegisterPod(i IPod) {
	localPod = i
}
