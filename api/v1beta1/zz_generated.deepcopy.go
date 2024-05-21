//go:build !ignore_autogenerated

// Copyright (c) 2021 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"open-cluster-management.io/config-policy-controller/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceConfig) DeepCopyInto(out *ComplianceConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceConfig.
func (in *ComplianceConfig) DeepCopy() *ComplianceConfig {
	if in == nil {
		return nil
	}
	out := new(ComplianceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorPolicy) DeepCopyInto(out *OperatorPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorPolicy.
func (in *OperatorPolicy) DeepCopy() *OperatorPolicy {
	if in == nil {
		return nil
	}
	out := new(OperatorPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorPolicyList) DeepCopyInto(out *OperatorPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OperatorPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorPolicyList.
func (in *OperatorPolicyList) DeepCopy() *OperatorPolicyList {
	if in == nil {
		return nil
	}
	out := new(OperatorPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorPolicySpec) DeepCopyInto(out *OperatorPolicySpec) {
	*out = *in
	if in.OperatorGroup != nil {
		in, out := &in.OperatorGroup, &out.OperatorGroup
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	in.Subscription.DeepCopyInto(&out.Subscription)
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]v1.NonEmptyString, len(*in))
		copy(*out, *in)
	}
	out.RemovalBehavior = in.RemovalBehavior
	out.ComplianceConfig = in.ComplianceConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorPolicySpec.
func (in *OperatorPolicySpec) DeepCopy() *OperatorPolicySpec {
	if in == nil {
		return nil
	}
	out := new(OperatorPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorPolicyStatus) DeepCopyInto(out *OperatorPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]v1.RelatedObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OverlappingPolicies != nil {
		in, out := &in.OverlappingPolicies, &out.OverlappingPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorPolicyStatus.
func (in *OperatorPolicyStatus) DeepCopy() *OperatorPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(OperatorPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemovalBehavior) DeepCopyInto(out *RemovalBehavior) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemovalBehavior.
func (in *RemovalBehavior) DeepCopy() *RemovalBehavior {
	if in == nil {
		return nil
	}
	out := new(RemovalBehavior)
	in.DeepCopyInto(out)
	return out
}
