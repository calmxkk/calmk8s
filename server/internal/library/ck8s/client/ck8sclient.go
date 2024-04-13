package client

import (
	"calmk8s/internal/model"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type Client struct {
	clientSet kubernetes.Interface
}

func NewClient(ctx context.Context, config *model.ClusterConfig) (client *Client, err error) {
	client = &Client{}
	var cfg = &rest.Config{}

	if config.Mode == "configfile" {
		cfg, err = clientcmd.BuildConfigFromKubeconfigGetter("", func() (*clientcmdapi.Config, error) {
			return clientcmd.Load(config.ConfigContent)
		})
		if err != nil {
			g.Log().Debugf(ctx, "get rest.config error, %v", err)
			return nil, err
		}
	} else {
		panic("not support other config func, only support config file")
	}

	client.clientSet, err = kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Client() kubernetes.Interface {
	return c.clientSet
}
