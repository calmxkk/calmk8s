package dockerck

import (
	"context"
)

func main() {
	cli := NewClient()
	cli.GetImageByName(context.TODO(), "nginx:latest")
}
