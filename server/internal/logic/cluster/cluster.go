package cluster

import (
	"calmk8s/internal/library/ck8s/informer"
	"calmk8s/internal/model"
	"calmk8s/internal/model/input/k8sin"
	"calmk8s/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"k8s.io/client-go/informers"
)

type sCluster struct{}

func NewCluster() *sCluster {
	return &sCluster{}
}

func init() {
	service.RegisterCluster(NewCluster())
}

type K8sInformer struct {
	Name     string
	stopCh   chan struct{}
	config   *model.ClusterConfig
	Informer informers.SharedInformerFactory
}

var AllCluster = make(map[string]*K8sInformer)

func (s *sCluster) GetInformer(ctx context.Context, name string) (informers.SharedInformerFactory, error) {
	g.Log().Debugf(ctx, "get informer for name: %v", name)
	if _, ok := AllCluster[name]; !ok {
		return nil, gerror.Newf("do not have cluster %v", name)
	}
	return AllCluster[name].Informer, nil
}

func (s *sCluster) Create(ctx context.Context, in *k8sin.Cluster) (err error) {
	if _, ok := AllCluster[in.Name]; ok {
		g.Log().Fatal(ctx, "k8s informer already exist")
		return nil
	}

	cfg := &model.ClusterConfig{
		Mode:          in.Authentication.Mode,
		ConfigContent: []byte(in.Authentication.ConfigFileContent),
	}

	k8sInformer := &K8sInformer{
		Name:   in.Name,
		stopCh: make(chan struct{}),
		config: cfg,
	}

	k8sInformer.Informer, err = informer.NewSharedInformerFactory(ctx, k8sInformer.config, k8sInformer.stopCh)
	if err != nil {
		return nil
	}

	AllCluster[in.Name] = k8sInformer
	g.Log().Infof(ctx, "create cluster %v success", in.Name)
	return nil
}
