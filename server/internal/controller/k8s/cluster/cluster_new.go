// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package cluster

import (
	"calmk8s/api/k8s/cluster"
)

type ControllerV1 struct{}

func NewV1() cluster.IClusterV1 {
	return &ControllerV1{}
}