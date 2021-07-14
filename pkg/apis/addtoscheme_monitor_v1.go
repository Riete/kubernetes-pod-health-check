package apis

import (
	v1 "github.com/riete/kubernetes-pod-health-check/pkg/apis/monitor/v1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1.SchemeBuilder.AddToScheme)
}
