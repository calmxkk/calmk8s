package main

import (
	"calmk8s/internal/library/ck8s/informer"
	"calmk8s/internal/model"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"k8s.io/apimachinery/pkg/labels"
)

func main() {
	configFileByte := gfile.GetBytes("./kubeconfig")
	cfg := &model.ClusterConfig{
		Mode:          "configfile",
		ConfigContent: configFileByte,
	}

	stopchan := make(chan struct{})

	newif, err := informer.NewSharedInformerFactory(context.TODO(), cfg, stopchan)
	if err != nil {
		fmt.Println(err)
		return
	}

	items, err := newif.Core().V1().Pods().Lister().List(labels.Everything())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range items {
		fmt.Printf("namespace %v, pod %v\n", item.Namespace, item.Name)
	}

	pod, err := newif.Core().V1().Pods().Lister().Pods("kube-system").Get("kube-scheduler-ck-virtual-machine")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Printf("namespace %v, pod %v\n", pod.Namespace, pod.Name)
}
