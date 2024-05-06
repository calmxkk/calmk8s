package dockerutil

import (
	"context"
	"io"

	"github.com/docker/docker/api/types/image"
	dockerClient "github.com/docker/docker/client"
	"github.com/gogf/gf/v2/frame/g"
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
	outio, err := c.cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		g.Log().Errorf(ctx, "pull image error, %v", err)
		return
	}

	defer outio.Close()

	_, err = io.Copy(io.Discard, outio)
	if err != nil {
		g.Log().Errorf(ctx, "pull image error, %v", err)
		return
	}
}

func (c *Client) RemoveImageByName(ctx context.Context, imageName string) {
	_, err := c.cli.ImageRemove(ctx, imageName, image.RemoveOptions{})
	if err != nil {
		g.Log().Errorf(ctx, "remove image error, %v", err)
		return
	}
}
