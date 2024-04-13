package pod

import (
	"calmk8s/internal/model/input/k8sin"
	"calmk8s/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
)

type sPod struct{}

func NewPod() *sPod {
	return &sPod{}
}

func init() {
	service.RegisterPod(NewPod())
}

func (s *sPod) GetPodByName(ctx context.Context, in *k8sin.GetPodInp) (*k8sin.GetPodRes, error) {
	k8sInformer, err := service.Cluster().GetInformer(ctx, in.ClusterName)
	if err != nil {
		g.Log().Debugf(ctx, "get informer error, %v", err.Error())
		return nil, err
	}
	g.Log().Debugf(ctx, "get pod %v in %v", in.PodName, in.NameSpace)

	_, err = s.NamespaceExists(ctx, k8sInformer, in.NameSpace)
	if err != nil {
		return nil, err
	}

	pod, err := k8sInformer.Core().V1().Pods().Lister().Pods(in.NameSpace).Get(in.PodName)
	if err != nil {
		g.Log().Debugf(ctx, "get pod error, %v", err)
		return nil, err
	}

	res := &k8sin.GetPodRes{
		Name:      pod.Name,
		Namespace: pod.Namespace,
		Status:    string(pod.Status.Phase),
		IP:        pod.Status.PodIP,
		Node:      pod.Spec.NodeName,
	}

	return res, nil
}

func (s *sPod) List(ctx context.Context, in *k8sin.GetPodInp) ([]*k8sin.GetPodRes, error) {
	k8sInformer, err := service.Cluster().GetInformer(ctx, in.ClusterName)
	if err != nil {
		g.Log().Debugf(ctx, "get informer error, %v", err.Error())
		return nil, err
	}

	pods, err := k8sInformer.Core().V1().Pods().Lister().List(labels.Everything())
	if err != nil {
		g.Log().Debugf(ctx, "get pod list error, %v", err.Error())
		return nil, err
	}

	var res []*k8sin.GetPodRes

	for _, pod := range pods {
		temp := &k8sin.GetPodRes{
			Name:      pod.Name,
			Namespace: pod.Namespace,
			Status:    string(pod.Status.Phase),
			IP:        pod.Status.PodIP,
			Node:      pod.Spec.NodeName,
		}
		res = append(res, temp)
	}
	g.Log().Debugf(ctx, "return")
	return res, nil
}

func (s *sPod) NamespaceExists(ctx context.Context, k8sInformer informers.SharedInformerFactory, namespaceName string) (exist bool, err error) {
	namespaces, err := k8sInformer.Core().V1().Namespaces().Lister().List(labels.Everything())
	if err != nil {
		g.Log().Debugf(ctx, "get all namespaces error, %v", err)
		return false, err
	}

	exists := false
	for _, ns := range namespaces {
		if ns.Name == namespaceName {
			exists = true
		}
	}

	if !exists {
		return false, gerror.Newf("do not have namespace %v", namespaceName)
	}

	return true, nil
}
