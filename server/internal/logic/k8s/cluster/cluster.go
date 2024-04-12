package cluster

import (
	"calmk8s/internal/model/input/k8sin"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type sCluster struct{}

func NewCluster() *sCluster {
	return &sCluster{}
}

type K8sClient struct {
	Name          string
	Client        *kubernetes.Clientset
	ConfigContent []byte
}

var AllCluster = make(map[string]*K8sClient)

func (s *sCluster) List(ctx context.Context) (list []*k8sin.Cluster, err error) {
	return
}

func (s *sCluster) Create(ctx context.Context, in *k8sin.Cluster) (err error) {
	if _, ok := AllCluster[in.Name]; ok {
		g.Log().Fatal(ctx, "client already exist")
		return nil
	}

	client, err := s.Client(ctx, in)
	if err != nil {
		return err
	}

	AllCluster[in.Name] = client

	return nil
}

func (s *sCluster) Client(ctx context.Context, in *k8sin.Cluster) (*K8sClient, error) {
	cfg, err := s.Config(ctx, in)
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(cfg)
	return &K8sClient{
		ConfigContent: in.Authentication.ConfigFileContent,
		Name:          in.Name,
		Client:        client,
	}, nil
}

func (s *sCluster) Config(ctx context.Context, in *k8sin.Cluster) (*rest.Config, error) {
	if in.Authentication.Mode != "configfile" {
		g.Log().Errorf(ctx, "can not support without config file")
		return nil, gerror.New("can not support without config file")
	}

	cfg, err := clientcmd.BuildConfigFromKubeconfigGetter("", func() (*clientcmdapi.Config, error) {
		return clientcmd.Load(in.Authentication.ConfigFileContent)
	})
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
