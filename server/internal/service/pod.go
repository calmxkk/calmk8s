// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IPod interface {
		List(ctx context.Context)
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
