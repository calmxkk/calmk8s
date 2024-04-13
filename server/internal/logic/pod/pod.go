package pod

import (
	"calmk8s/internal/model/input/k8sin"
	"calmk8s/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"k8s.io/apimachinery/pkg/labels"
)

type sPod struct{}

func NewPod() *sPod {
	return &sPod{}
}

func init() {
	service.RegisterPod(NewPod())
}

func (s *sPod) GetPodByName(ctx context.Context, in *k8sin.GetPodInp) error {
	k8sInformer, err := service.Cluster().GetInformer(ctx, in.ClusterName)
	if err != nil {
		g.Log().Debugf(ctx, "get informer error, %v", err.Error())
		return err
	}
	g.Log().Debugf(ctx, "get pod %v in %v", in.PodName, in.NameSpace)
	pod, err := k8sInformer.Core().V1().Pods().Lister().Pods(in.NameSpace).Get(in.PodName)
	if err != nil {
		g.Log().Debugf(ctx, "get pod error, %v", err)
		return err
	}
	fmt.Printf("namespace: %v, pod: %v\n", pod.Namespace, pod.Name)
	return nil
}

func (s *sPod) List(ctx context.Context, in *k8sin.GetPodInp) error {
	k8sInformer, err := service.Cluster().GetInformer(ctx, in.ClusterName)
	if err != nil {
		g.Log().Debugf(ctx, "get informer error, %v", err.Error())
		return err
	}

	pods, err := k8sInformer.Core().V1().Pods().Lister().List(labels.Everything())
	if err != nil {
		g.Log().Debugf(ctx, "get pod list error, %v", err.Error())
		return err
	}

	for _, pod := range pods {
		fmt.Printf("namespace: %v, pod: %v\n", pod.Namespace, pod.Name)
	}

	return nil
}
