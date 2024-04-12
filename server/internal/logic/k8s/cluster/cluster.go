package cluster

import (
	"calmk8s/internal/library/ck8s/informer"
	"calmk8s/internal/model"
	"calmk8s/internal/model/input/k8sin"
	"context"
	"k8s.io/client-go/informers"

	"github.com/gogf/gf/v2/frame/g"
)

type sCluster struct{}

func NewCluster() *sCluster {
	return &sCluster{}
}

type K8sInformer struct {
	Name     string
	stopCh   chan struct{}
	config   *model.ClusterConfig
	Informer informers.SharedInformerFactory
}

var AllCluster = make(map[string]*K8sInformer)

func (s *sCluster) Create(ctx context.Context, in *k8sin.Cluster) (err error) {
	if _, ok := AllCluster[in.Name]; ok {
		g.Log().Fatal(ctx, "k8s informer already exist")
		return nil
	}

	cfg := &model.ClusterConfig{
		Mode:          in.Authentication.Mode,
		ConfigContent: in.Authentication.ConfigFileContent,
	}

	k8sInformer := &K8sInformer{
		Name:   in.Name,
		stopCh: make(chan struct{}),
		config: cfg,
	}

	k8sInformer.Informer, err = informer.NewSharedInformerFactory(k8sInformer.config, k8sInformer.stopCh)
	if err != nil {
		return nil
	}

	AllCluster[in.Name] = k8sInformer
	return nil
}
