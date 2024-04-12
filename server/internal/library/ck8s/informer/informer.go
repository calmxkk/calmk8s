package informer

import (
	"calmk8s/internal/library/ck8s/client"
	"calmk8s/internal/model"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
	"time"
)

func NewSharedInformerFactory(config *model.ClusterConfig, stop <-chan struct{}) (*informers.SharedInformerFactory, error) {
	var (
		cli = &client.Client{}
		err error
	)

	cli, err = client.NewClient(config)
	if err != nil {
		return nil, err
	}

	sharedInformerFactory := informers.NewSharedInformerFactory(cli.Client(), time.Minute)

	gvrs := []schema.GroupVersionResource{
		{Group: "", Version: "v1", Resource: "pods"},
		{Group: "", Version: "v1", Resource: "services"},
		{Group: "", Version: "v1", Resource: "namespaces"},

		{Group: "apps", Version: "v1", Resource: "deployments"},
		{Group: "apps", Version: "v1", Resource: "statefulsets"},
		{Group: "apps", Version: "v1", Resource: "daemonsets"},
	}

	for _, s := range gvrs {
		_, err = sharedInformerFactory.ForResource(s)
		if err != nil {
			return nil, err
		}
	}

	sharedInformerFactory.Start(stop)
	sharedInformerFactory.WaitForCacheSync(stop)

	return &sharedInformerFactory, nil
}
