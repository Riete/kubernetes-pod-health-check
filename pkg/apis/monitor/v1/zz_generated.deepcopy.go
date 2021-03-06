// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DingTalk) DeepCopyInto(out *DingTalk) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DingTalk.
func (in *DingTalk) DeepCopy() *DingTalk {
	if in == nil {
		return nil
	}
	out := new(DingTalk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheck) DeepCopyInto(out *HealthCheck) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheck.
func (in *HealthCheck) DeepCopy() *HealthCheck {
	if in == nil {
		return nil
	}
	out := new(HealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodHealthCheck) DeepCopyInto(out *PodHealthCheck) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodHealthCheck.
func (in *PodHealthCheck) DeepCopy() *PodHealthCheck {
	if in == nil {
		return nil
	}
	out := new(PodHealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodHealthCheck) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodHealthCheckList) DeepCopyInto(out *PodHealthCheckList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodHealthCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodHealthCheckList.
func (in *PodHealthCheckList) DeepCopy() *PodHealthCheckList {
	if in == nil {
		return nil
	}
	out := new(PodHealthCheckList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodHealthCheckList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodHealthCheckSpec) DeepCopyInto(out *PodHealthCheckSpec) {
	*out = *in
	out.DingTalk = in.DingTalk
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.HealthCheck = in.HealthCheck
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodHealthCheckSpec.
func (in *PodHealthCheckSpec) DeepCopy() *PodHealthCheckSpec {
	if in == nil {
		return nil
	}
	out := new(PodHealthCheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodHealthCheckStatus) DeepCopyInto(out *PodHealthCheckStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodHealthCheckStatus.
func (in *PodHealthCheckStatus) DeepCopy() *PodHealthCheckStatus {
	if in == nil {
		return nil
	}
	out := new(PodHealthCheckStatus)
	in.DeepCopyInto(out)
	return out
}
