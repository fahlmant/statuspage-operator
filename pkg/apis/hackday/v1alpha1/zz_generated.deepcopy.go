// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusPageCustomer) DeepCopyInto(out *StatusPageCustomer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusPageCustomer.
func (in *StatusPageCustomer) DeepCopy() *StatusPageCustomer {
	if in == nil {
		return nil
	}
	out := new(StatusPageCustomer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StatusPageCustomer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusPageCustomerList) DeepCopyInto(out *StatusPageCustomerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StatusPageCustomer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusPageCustomerList.
func (in *StatusPageCustomerList) DeepCopy() *StatusPageCustomerList {
	if in == nil {
		return nil
	}
	out := new(StatusPageCustomerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StatusPageCustomerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusPageCustomerSpec) DeepCopyInto(out *StatusPageCustomerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusPageCustomerSpec.
func (in *StatusPageCustomerSpec) DeepCopy() *StatusPageCustomerSpec {
	if in == nil {
		return nil
	}
	out := new(StatusPageCustomerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusPageCustomerStatus) DeepCopyInto(out *StatusPageCustomerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusPageCustomerStatus.
func (in *StatusPageCustomerStatus) DeepCopy() *StatusPageCustomerStatus {
	if in == nil {
		return nil
	}
	out := new(StatusPageCustomerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusPageIncident) DeepCopyInto(out *StatusPageIncident) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusPageIncident.
func (in *StatusPageIncident) DeepCopy() *StatusPageIncident {
	if in == nil {
		return nil
	}
	out := new(StatusPageIncident)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StatusPageIncident) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusPageIncidentList) DeepCopyInto(out *StatusPageIncidentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StatusPageIncident, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusPageIncidentList.
func (in *StatusPageIncidentList) DeepCopy() *StatusPageIncidentList {
	if in == nil {
		return nil
	}
	out := new(StatusPageIncidentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StatusPageIncidentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusPageIncidentSpec) DeepCopyInto(out *StatusPageIncidentSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusPageIncidentSpec.
func (in *StatusPageIncidentSpec) DeepCopy() *StatusPageIncidentSpec {
	if in == nil {
		return nil
	}
	out := new(StatusPageIncidentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusPageIncidentStatus) DeepCopyInto(out *StatusPageIncidentStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusPageIncidentStatus.
func (in *StatusPageIncidentStatus) DeepCopy() *StatusPageIncidentStatus {
	if in == nil {
		return nil
	}
	out := new(StatusPageIncidentStatus)
	in.DeepCopyInto(out)
	return out
}
