package dockerutil

import (
	"context"
	"testing"
)

func TestPullImage(t *testing.T) {
	cli := NewClient()
	cli.RemoveImageByName(context.TODO(), "nginx:latest")
}
