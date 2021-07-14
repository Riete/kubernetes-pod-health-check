package controller

import (
	"github.com/riete/kubernetes-pod-health-check/pkg/controller/podhealthcheck"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, podhealthcheck.Add)
}
