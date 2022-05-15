# Controller
Controls to ensure that `Graph` resources are always in sync with the cluster state.

### Intercept Resource Events
1. Intercept all events from relevant informers.

### Process Events
1. Filter Connections for relevant resources.
2. Check resource availability in namespace/cluster.

### Add/Update/Delete Graph 
> [Examples](../examples/graphs/)
1. Add/Update/Delete `tekton.visualise/v1alpha1/Graph` resources.

### References:
1. https://medium.com/speechmatics/how-to-write-kubernetes-custom-controllers-in-go-8014c4a04235
2. https://itnext.io/building-your-own-kubernetes-crds-701de1c9a161
3. https://cloud.redhat.com/blog/kubernetes-deep-dive-code-generation-customresources
4. https://book.kubebuilder.io/cronjob-tutorial/new-api.html
