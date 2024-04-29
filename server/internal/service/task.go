// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ISystemTask interface {
		Run(ctx context.Context)
		GetMemory(ctx context.Context)
		GetSystemInfo(ctx context.Context)
	}
)

var (
	localSystemTask ISystemTask
)

func SystemTask() ISystemTask {
	if localSystemTask == nil {
		panic("implement not found for interface ISystemTask, forgot register?")
	}
	return localSystemTask
}

func RegisterSystemTask(i ISystemTask) {
	localSystemTask = i
}
