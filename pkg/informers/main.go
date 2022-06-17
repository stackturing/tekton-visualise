package infromers

import (
	"errors"
	"time"

	"github.com/containerd/containerd/log"
	tekton_pipeline_v1alpha1 "github.com/tektoncdpkg/apis/pipeline/v1alpha1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

func WatchTaskResources(clientSet tekton_pipeline_v1alpha1.alpha1.Task, namespace string) cache.Store {
	projectStore, projectController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return clientSet.Projects(namespace).List(lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return clientSet.Projects(namespace).Watch(lo)
			},
		},
		&v1alpha1.Project{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{},
	)

	go projectController.Run(wait.NeverStop)
	return projectStore
}

// func (c *visualiseCacheImpl) isTektonSynced(namespace string) bool {
// 	var isSynced bool
// 	if nsCache, exist := c.nsCache[namespace]; exist {
// 		isSynced = nsCache[tekton_pipeline_v1alpha1.TaskType].HasSynced()
// 	} else {
// 		isSynced = false
// 	}
// 	return isSynced
// }
