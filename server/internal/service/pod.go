// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"calmk8s/internal/model/input/k8sin"
	"context"
)

type (
	IPod interface {
		GetPodByName(ctx context.Context, in *k8sin.GetPodInp) error
		List(ctx context.Context, in *k8sin.GetPodInp) error
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
