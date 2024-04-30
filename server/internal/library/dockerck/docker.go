package dockerck

import (
	"context"

	"github.com/docker/docker/api/types/image"
	dockerClient "github.com/docker/docker/client"
)

type Client struct {
	cli *dockerClient.Client
}

func NewClient() *Client {
	cli, err := dockerClient.NewClientWithOpts(dockerClient.FromEnv, dockerClient.WithAPIVersionNegotiation())
	if err != nil {
		panic("can not create docker client")
	}

	return &Client{
		cli: cli,
	}
}

func (c *Client) GetImageByName(ctx context.Context, imageName string) {
	c.cli.ImagePull(ctx, imageName, image.PullOptions{})
}
