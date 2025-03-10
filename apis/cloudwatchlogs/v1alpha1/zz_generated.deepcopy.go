//go:build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountPolicy) DeepCopyInto(out *AccountPolicy) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.LastUpdatedTime != nil {
		in, out := &in.LastUpdatedTime, &out.LastUpdatedTime
		*out = new(int64)
		**out = **in
	}
	if in.PolicyName != nil {
		in, out := &in.PolicyName, &out.PolicyName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountPolicy.
func (in *AccountPolicy) DeepCopy() *AccountPolicy {
	if in == nil {
		return nil
	}
	out := new(AccountPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomLogGroupObservation) DeepCopyInto(out *CustomLogGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomLogGroupObservation.
func (in *CustomLogGroupObservation) DeepCopy() *CustomLogGroupObservation {
	if in == nil {
		return nil
	}
	out := new(CustomLogGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomLogGroupParameters) DeepCopyInto(out *CustomLogGroupParameters) {
	*out = *in
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(int64)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyIDRef != nil {
		in, out := &in.KMSKeyIDRef, &out.KMSKeyIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeyIDSelector != nil {
		in, out := &in.KMSKeyIDSelector, &out.KMSKeyIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomLogGroupParameters.
func (in *CustomLogGroupParameters) DeepCopy() *CustomLogGroupParameters {
	if in == nil {
		return nil
	}
	out := new(CustomLogGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourcePolicyObservation) DeepCopyInto(out *CustomResourcePolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourcePolicyObservation.
func (in *CustomResourcePolicyObservation) DeepCopy() *CustomResourcePolicyObservation {
	if in == nil {
		return nil
	}
	out := new(CustomResourcePolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourcePolicyParameters) DeepCopyInto(out *CustomResourcePolicyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourcePolicyParameters.
func (in *CustomResourcePolicyParameters) DeepCopy() *CustomResourcePolicyParameters {
	if in == nil {
		return nil
	}
	out := new(CustomResourcePolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Delivery) DeepCopyInto(out *Delivery) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.DeliveryDestinationARN != nil {
		in, out := &in.DeliveryDestinationARN, &out.DeliveryDestinationARN
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Delivery.
func (in *Delivery) DeepCopy() *Delivery {
	if in == nil {
		return nil
	}
	out := new(Delivery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeliveryDestination) DeepCopyInto(out *DeliveryDestination) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeliveryDestination.
func (in *DeliveryDestination) DeepCopy() *DeliveryDestination {
	if in == nil {
		return nil
	}
	out := new(DeliveryDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeliveryDestinationConfiguration) DeepCopyInto(out *DeliveryDestinationConfiguration) {
	*out = *in
	if in.DestinationResourceARN != nil {
		in, out := &in.DestinationResourceARN, &out.DestinationResourceARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeliveryDestinationConfiguration.
func (in *DeliveryDestinationConfiguration) DeepCopy() *DeliveryDestinationConfiguration {
	if in == nil {
		return nil
	}
	out := new(DeliveryDestinationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeliverySource) DeepCopyInto(out *DeliverySource) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeliverySource.
func (in *DeliverySource) DeepCopy() *DeliverySource {
	if in == nil {
		return nil
	}
	out := new(DeliverySource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Destination) DeepCopyInto(out *Destination) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destination.
func (in *Destination) DeepCopy() *Destination {
	if in == nil {
		return nil
	}
	out := new(Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportTask) DeepCopyInto(out *ExportTask) {
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(int64)
		**out = **in
	}
	if in.LogGroupName != nil {
		in, out := &in.LogGroupName, &out.LogGroupName
		*out = new(string)
		**out = **in
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportTask.
func (in *ExportTask) DeepCopy() *ExportTask {
	if in == nil {
		return nil
	}
	out := new(ExportTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportTaskExecutionInfo) DeepCopyInto(out *ExportTaskExecutionInfo) {
	*out = *in
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = new(int64)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportTaskExecutionInfo.
func (in *ExportTaskExecutionInfo) DeepCopy() *ExportTaskExecutionInfo {
	if in == nil {
		return nil
	}
	out := new(ExportTaskExecutionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilteredLogEvent) DeepCopyInto(out *FilteredLogEvent) {
	*out = *in
	if in.IngestionTime != nil {
		in, out := &in.IngestionTime, &out.IngestionTime
		*out = new(int64)
		**out = **in
	}
	if in.Timestamp != nil {
		in, out := &in.Timestamp, &out.Timestamp
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilteredLogEvent.
func (in *FilteredLogEvent) DeepCopy() *FilteredLogEvent {
	if in == nil {
		return nil
	}
	out := new(FilteredLogEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InputLogEvent) DeepCopyInto(out *InputLogEvent) {
	*out = *in
	if in.Timestamp != nil {
		in, out := &in.Timestamp, &out.Timestamp
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InputLogEvent.
func (in *InputLogEvent) DeepCopy() *InputLogEvent {
	if in == nil {
		return nil
	}
	out := new(InputLogEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogGroup) DeepCopyInto(out *LogGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogGroup.
func (in *LogGroup) DeepCopy() *LogGroup {
	if in == nil {
		return nil
	}
	out := new(LogGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LogGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogGroupList) DeepCopyInto(out *LogGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LogGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogGroupList.
func (in *LogGroupList) DeepCopy() *LogGroupList {
	if in == nil {
		return nil
	}
	out := new(LogGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LogGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogGroupObservation) DeepCopyInto(out *LogGroupObservation) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = new(int64)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.LogGroupName != nil {
		in, out := &in.LogGroupName, &out.LogGroupName
		*out = new(string)
		**out = **in
	}
	if in.MetricFilterCount != nil {
		in, out := &in.MetricFilterCount, &out.MetricFilterCount
		*out = new(int64)
		**out = **in
	}
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(int64)
		**out = **in
	}
	if in.StoredBytes != nil {
		in, out := &in.StoredBytes, &out.StoredBytes
		*out = new(int64)
		**out = **in
	}
	out.CustomLogGroupObservation = in.CustomLogGroupObservation
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogGroupObservation.
func (in *LogGroupObservation) DeepCopy() *LogGroupObservation {
	if in == nil {
		return nil
	}
	out := new(LogGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogGroupParameters) DeepCopyInto(out *LogGroupParameters) {
	*out = *in
	if in.LogGroupName != nil {
		in, out := &in.LogGroupName, &out.LogGroupName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	in.CustomLogGroupParameters.DeepCopyInto(&out.CustomLogGroupParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogGroupParameters.
func (in *LogGroupParameters) DeepCopy() *LogGroupParameters {
	if in == nil {
		return nil
	}
	out := new(LogGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogGroupSpec) DeepCopyInto(out *LogGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogGroupSpec.
func (in *LogGroupSpec) DeepCopy() *LogGroupSpec {
	if in == nil {
		return nil
	}
	out := new(LogGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogGroupStatus) DeepCopyInto(out *LogGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogGroupStatus.
func (in *LogGroupStatus) DeepCopy() *LogGroupStatus {
	if in == nil {
		return nil
	}
	out := new(LogGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogGroup_SDK) DeepCopyInto(out *LogGroup_SDK) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = new(int64)
		**out = **in
	}
	if in.DataProtectionStatus != nil {
		in, out := &in.DataProtectionStatus, &out.DataProtectionStatus
		*out = new(string)
		**out = **in
	}
	if in.InheritedProperties != nil {
		in, out := &in.InheritedProperties, &out.InheritedProperties
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.LogGroupName != nil {
		in, out := &in.LogGroupName, &out.LogGroupName
		*out = new(string)
		**out = **in
	}
	if in.MetricFilterCount != nil {
		in, out := &in.MetricFilterCount, &out.MetricFilterCount
		*out = new(int64)
		**out = **in
	}
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(int64)
		**out = **in
	}
	if in.StoredBytes != nil {
		in, out := &in.StoredBytes, &out.StoredBytes
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogGroup_SDK.
func (in *LogGroup_SDK) DeepCopy() *LogGroup_SDK {
	if in == nil {
		return nil
	}
	out := new(LogGroup_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogStream) DeepCopyInto(out *LogStream) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = new(int64)
		**out = **in
	}
	if in.FirstEventTimestamp != nil {
		in, out := &in.FirstEventTimestamp, &out.FirstEventTimestamp
		*out = new(int64)
		**out = **in
	}
	if in.LastEventTimestamp != nil {
		in, out := &in.LastEventTimestamp, &out.LastEventTimestamp
		*out = new(int64)
		**out = **in
	}
	if in.LastIngestionTime != nil {
		in, out := &in.LastIngestionTime, &out.LastIngestionTime
		*out = new(int64)
		**out = **in
	}
	if in.StoredBytes != nil {
		in, out := &in.StoredBytes, &out.StoredBytes
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogStream.
func (in *LogStream) DeepCopy() *LogStream {
	if in == nil {
		return nil
	}
	out := new(LogStream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricFilter) DeepCopyInto(out *MetricFilter) {
	*out = *in
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = new(int64)
		**out = **in
	}
	if in.LogGroupName != nil {
		in, out := &in.LogGroupName, &out.LogGroupName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricFilter.
func (in *MetricFilter) DeepCopy() *MetricFilter {
	if in == nil {
		return nil
	}
	out := new(MetricFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputLogEvent) DeepCopyInto(out *OutputLogEvent) {
	*out = *in
	if in.IngestionTime != nil {
		in, out := &in.IngestionTime, &out.IngestionTime
		*out = new(int64)
		**out = **in
	}
	if in.Timestamp != nil {
		in, out := &in.Timestamp, &out.Timestamp
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputLogEvent.
func (in *OutputLogEvent) DeepCopy() *OutputLogEvent {
	if in == nil {
		return nil
	}
	out := new(OutputLogEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryDefinition) DeepCopyInto(out *QueryDefinition) {
	*out = *in
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryDefinition.
func (in *QueryDefinition) DeepCopy() *QueryDefinition {
	if in == nil {
		return nil
	}
	out := new(QueryDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryInfo) DeepCopyInto(out *QueryInfo) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(int64)
		**out = **in
	}
	if in.LogGroupName != nil {
		in, out := &in.LogGroupName, &out.LogGroupName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryInfo.
func (in *QueryInfo) DeepCopy() *QueryInfo {
	if in == nil {
		return nil
	}
	out := new(QueryInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicy) DeepCopyInto(out *ResourcePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicy.
func (in *ResourcePolicy) DeepCopy() *ResourcePolicy {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourcePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicyList) DeepCopyInto(out *ResourcePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourcePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicyList.
func (in *ResourcePolicyList) DeepCopy() *ResourcePolicyList {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourcePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicyObservation) DeepCopyInto(out *ResourcePolicyObservation) {
	*out = *in
	if in.LastUpdatedTime != nil {
		in, out := &in.LastUpdatedTime, &out.LastUpdatedTime
		*out = new(int64)
		**out = **in
	}
	if in.PolicyName != nil {
		in, out := &in.PolicyName, &out.PolicyName
		*out = new(string)
		**out = **in
	}
	out.CustomResourcePolicyObservation = in.CustomResourcePolicyObservation
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicyObservation.
func (in *ResourcePolicyObservation) DeepCopy() *ResourcePolicyObservation {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicyParameters) DeepCopyInto(out *ResourcePolicyParameters) {
	*out = *in
	if in.PolicyDocument != nil {
		in, out := &in.PolicyDocument, &out.PolicyDocument
		*out = new(string)
		**out = **in
	}
	out.CustomResourcePolicyParameters = in.CustomResourcePolicyParameters
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicyParameters.
func (in *ResourcePolicyParameters) DeepCopy() *ResourcePolicyParameters {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicySpec) DeepCopyInto(out *ResourcePolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicySpec.
func (in *ResourcePolicySpec) DeepCopy() *ResourcePolicySpec {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicyStatus) DeepCopyInto(out *ResourcePolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicyStatus.
func (in *ResourcePolicyStatus) DeepCopy() *ResourcePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicy_SDK) DeepCopyInto(out *ResourcePolicy_SDK) {
	*out = *in
	if in.LastUpdatedTime != nil {
		in, out := &in.LastUpdatedTime, &out.LastUpdatedTime
		*out = new(int64)
		**out = **in
	}
	if in.PolicyDocument != nil {
		in, out := &in.PolicyDocument, &out.PolicyDocument
		*out = new(string)
		**out = **in
	}
	if in.PolicyName != nil {
		in, out := &in.PolicyName, &out.PolicyName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicy_SDK.
func (in *ResourcePolicy_SDK) DeepCopy() *ResourcePolicy_SDK {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicy_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionFilter) DeepCopyInto(out *SubscriptionFilter) {
	*out = *in
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = new(int64)
		**out = **in
	}
	if in.LogGroupName != nil {
		in, out := &in.LogGroupName, &out.LogGroupName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionFilter.
func (in *SubscriptionFilter) DeepCopy() *SubscriptionFilter {
	if in == nil {
		return nil
	}
	out := new(SubscriptionFilter)
	in.DeepCopyInto(out)
	return out
}
