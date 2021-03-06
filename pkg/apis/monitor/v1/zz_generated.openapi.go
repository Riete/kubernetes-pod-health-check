// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/riete/kubernetes-pod-health-check/pkg/apis/monitor/v1.DingTalk":             schema_pkg_apis_monitor_v1_DingTalk(ref),
		"github.com/riete/kubernetes-pod-health-check/pkg/apis/monitor/v1.PodHealthCheck":       schema_pkg_apis_monitor_v1_PodHealthCheck(ref),
		"github.com/riete/kubernetes-pod-health-check/pkg/apis/monitor/v1.PodHealthCheckStatus": schema_pkg_apis_monitor_v1_PodHealthCheckStatus(ref),
	}
}

func schema_pkg_apis_monitor_v1_DingTalk(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PodHealthCheckSpec defines the desired state of PodHealthCheck",
				Properties: map[string]spec.Schema{
					"webhook": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"secret": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"webhook", "secret"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_monitor_v1_PodHealthCheck(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PodHealthCheck is the Schema for the podhealthchecks API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/riete/kubernetes-pod-health-check/pkg/apis/monitor/v1.PodHealthCheckSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/riete/kubernetes-pod-health-check/pkg/apis/monitor/v1.PodHealthCheckStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/riete/kubernetes-pod-health-check/pkg/apis/monitor/v1.PodHealthCheckSpec", "github.com/riete/kubernetes-pod-health-check/pkg/apis/monitor/v1.PodHealthCheckStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_monitor_v1_PodHealthCheckStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PodHealthCheckStatus defines the observed state of PodHealthCheck",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
