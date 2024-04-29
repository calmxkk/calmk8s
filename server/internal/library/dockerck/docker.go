package dockerck

import (
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
