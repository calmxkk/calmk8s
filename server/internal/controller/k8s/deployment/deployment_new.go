// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package deployment

import (
	"calmk8s/api/k8s/deployment"
)

type ControllerV1 struct{}

func NewV1() deployment.IDeploymentV1 {
	return &ControllerV1{}
}
