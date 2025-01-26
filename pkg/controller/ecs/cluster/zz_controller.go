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

// Code generated by ack-generate. DO NOT EDIT.

package cluster

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/ecs"
	svcsdk "github.com/aws/aws-sdk-go/service/ecs"
	svcsdkapi "github.com/aws/aws-sdk-go/service/ecs/ecsiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/ecs/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an Cluster resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create Cluster in AWS"
	errUpdate        = "cannot update Cluster in AWS"
	errDescribe      = "failed to describe Cluster"
	errDelete        = "failed to delete Cluster"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.Cluster)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := connectaws.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.Cluster)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeClustersInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeClustersWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateCluster(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
	upToDate := true
	diff := ""
	if !meta.WasDeleted(cr) { // There is no need to run isUpToDate if the resource is deleted
		upToDate, diff, err = e.isUpToDate(ctx, cr, resp)
		if err != nil {
			return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
		}
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		Diff:                    diff,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.Cluster)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateClusterInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateClusterWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.Cluster.ActiveServicesCount != nil {
		cr.Status.AtProvider.ActiveServicesCount = resp.Cluster.ActiveServicesCount
	} else {
		cr.Status.AtProvider.ActiveServicesCount = nil
	}
	if resp.Cluster.Attachments != nil {
		f1 := []*svcapitypes.Attachment{}
		for _, f1iter := range resp.Cluster.Attachments {
			f1elem := &svcapitypes.Attachment{}
			if f1iter.Details != nil {
				f1elemf0 := []*svcapitypes.KeyValuePair{}
				for _, f1elemf0iter := range f1iter.Details {
					f1elemf0elem := &svcapitypes.KeyValuePair{}
					if f1elemf0iter.Name != nil {
						f1elemf0elem.Name = f1elemf0iter.Name
					}
					if f1elemf0iter.Value != nil {
						f1elemf0elem.Value = f1elemf0iter.Value
					}
					f1elemf0 = append(f1elemf0, f1elemf0elem)
				}
				f1elem.Details = f1elemf0
			}
			if f1iter.Id != nil {
				f1elem.ID = f1iter.Id
			}
			if f1iter.Status != nil {
				f1elem.Status = f1iter.Status
			}
			if f1iter.Type != nil {
				f1elem.Type = f1iter.Type
			}
			f1 = append(f1, f1elem)
		}
		cr.Status.AtProvider.Attachments = f1
	} else {
		cr.Status.AtProvider.Attachments = nil
	}
	if resp.Cluster.AttachmentsStatus != nil {
		cr.Status.AtProvider.AttachmentsStatus = resp.Cluster.AttachmentsStatus
	} else {
		cr.Status.AtProvider.AttachmentsStatus = nil
	}
	if resp.Cluster.CapacityProviders != nil {
		f3 := []*string{}
		for _, f3iter := range resp.Cluster.CapacityProviders {
			var f3elem string
			f3elem = *f3iter
			f3 = append(f3, &f3elem)
		}
		cr.Spec.ForProvider.CapacityProviders = f3
	} else {
		cr.Spec.ForProvider.CapacityProviders = nil
	}
	if resp.Cluster.ClusterArn != nil {
		cr.Status.AtProvider.ClusterARN = resp.Cluster.ClusterArn
	} else {
		cr.Status.AtProvider.ClusterARN = nil
	}
	if resp.Cluster.ClusterName != nil {
		cr.Spec.ForProvider.ClusterName = resp.Cluster.ClusterName
	} else {
		cr.Spec.ForProvider.ClusterName = nil
	}
	if resp.Cluster.Configuration != nil {
		f6 := &svcapitypes.ClusterConfiguration{}
		if resp.Cluster.Configuration.ExecuteCommandConfiguration != nil {
			f6f0 := &svcapitypes.ExecuteCommandConfiguration{}
			if resp.Cluster.Configuration.ExecuteCommandConfiguration.KmsKeyId != nil {
				f6f0.KMSKeyID = resp.Cluster.Configuration.ExecuteCommandConfiguration.KmsKeyId
			}
			if resp.Cluster.Configuration.ExecuteCommandConfiguration.LogConfiguration != nil {
				f6f0f1 := &svcapitypes.ExecuteCommandLogConfiguration{}
				if resp.Cluster.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled != nil {
					f6f0f1.CloudWatchEncryptionEnabled = resp.Cluster.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled
				}
				if resp.Cluster.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName != nil {
					f6f0f1.CloudWatchLogGroupName = resp.Cluster.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName
				}
				if resp.Cluster.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName != nil {
					f6f0f1.S3BucketName = resp.Cluster.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName
				}
				if resp.Cluster.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled != nil {
					f6f0f1.S3EncryptionEnabled = resp.Cluster.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled
				}
				if resp.Cluster.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix != nil {
					f6f0f1.S3KeyPrefix = resp.Cluster.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix
				}
				f6f0.LogConfiguration = f6f0f1
			}
			if resp.Cluster.Configuration.ExecuteCommandConfiguration.Logging != nil {
				f6f0.Logging = resp.Cluster.Configuration.ExecuteCommandConfiguration.Logging
			}
			f6.ExecuteCommandConfiguration = f6f0
		}
		cr.Spec.ForProvider.Configuration = f6
	} else {
		cr.Spec.ForProvider.Configuration = nil
	}
	if resp.Cluster.DefaultCapacityProviderStrategy != nil {
		f7 := []*svcapitypes.CapacityProviderStrategyItem{}
		for _, f7iter := range resp.Cluster.DefaultCapacityProviderStrategy {
			f7elem := &svcapitypes.CapacityProviderStrategyItem{}
			if f7iter.Base != nil {
				f7elem.Base = f7iter.Base
			}
			if f7iter.CapacityProvider != nil {
				f7elem.CapacityProvider = f7iter.CapacityProvider
			}
			if f7iter.Weight != nil {
				f7elem.Weight = f7iter.Weight
			}
			f7 = append(f7, f7elem)
		}
		cr.Spec.ForProvider.DefaultCapacityProviderStrategy = f7
	} else {
		cr.Spec.ForProvider.DefaultCapacityProviderStrategy = nil
	}
	if resp.Cluster.PendingTasksCount != nil {
		cr.Status.AtProvider.PendingTasksCount = resp.Cluster.PendingTasksCount
	} else {
		cr.Status.AtProvider.PendingTasksCount = nil
	}
	if resp.Cluster.RegisteredContainerInstancesCount != nil {
		cr.Status.AtProvider.RegisteredContainerInstancesCount = resp.Cluster.RegisteredContainerInstancesCount
	} else {
		cr.Status.AtProvider.RegisteredContainerInstancesCount = nil
	}
	if resp.Cluster.RunningTasksCount != nil {
		cr.Status.AtProvider.RunningTasksCount = resp.Cluster.RunningTasksCount
	} else {
		cr.Status.AtProvider.RunningTasksCount = nil
	}
	if resp.Cluster.ServiceConnectDefaults != nil {
		f11 := &svcapitypes.ClusterServiceConnectDefaultsRequest{}
		if resp.Cluster.ServiceConnectDefaults.Namespace != nil {
			f11.Namespace = resp.Cluster.ServiceConnectDefaults.Namespace
		}
		cr.Spec.ForProvider.ServiceConnectDefaults = f11
	} else {
		cr.Spec.ForProvider.ServiceConnectDefaults = nil
	}
	if resp.Cluster.Settings != nil {
		f12 := []*svcapitypes.ClusterSetting{}
		for _, f12iter := range resp.Cluster.Settings {
			f12elem := &svcapitypes.ClusterSetting{}
			if f12iter.Name != nil {
				f12elem.Name = f12iter.Name
			}
			if f12iter.Value != nil {
				f12elem.Value = f12iter.Value
			}
			f12 = append(f12, f12elem)
		}
		cr.Spec.ForProvider.Settings = f12
	} else {
		cr.Spec.ForProvider.Settings = nil
	}
	if resp.Cluster.Statistics != nil {
		f13 := []*svcapitypes.KeyValuePair{}
		for _, f13iter := range resp.Cluster.Statistics {
			f13elem := &svcapitypes.KeyValuePair{}
			if f13iter.Name != nil {
				f13elem.Name = f13iter.Name
			}
			if f13iter.Value != nil {
				f13elem.Value = f13iter.Value
			}
			f13 = append(f13, f13elem)
		}
		cr.Status.AtProvider.Statistics = f13
	} else {
		cr.Status.AtProvider.Statistics = nil
	}
	if resp.Cluster.Status != nil {
		cr.Status.AtProvider.Status = resp.Cluster.Status
	} else {
		cr.Status.AtProvider.Status = nil
	}
	if resp.Cluster.Tags != nil {
		f15 := []*svcapitypes.Tag{}
		for _, f15iter := range resp.Cluster.Tags {
			f15elem := &svcapitypes.Tag{}
			if f15iter.Key != nil {
				f15elem.Key = f15iter.Key
			}
			if f15iter.Value != nil {
				f15elem.Value = f15iter.Value
			}
			f15 = append(f15, f15elem)
		}
		cr.Spec.ForProvider.Tags = f15
	} else {
		cr.Spec.ForProvider.Tags = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.Cluster)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateClusterInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateClusterWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) (managed.ExternalDelete, error) {
	cr, ok := mg.(*svcapitypes.Cluster)
	if !ok {
		return managed.ExternalDelete{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteClusterInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return managed.ExternalDelete{}, errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return managed.ExternalDelete{}, nil
	}
	resp, err := e.client.DeleteClusterWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

func (e *external) Disconnect(ctx context.Context) error {
	// Unimplemented, required by newer versions of crossplane-runtime
	return nil
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.ECSAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.ECSAPI
	preObserve     func(context.Context, *svcapitypes.Cluster, *svcsdk.DescribeClustersInput) error
	postObserve    func(context.Context, *svcapitypes.Cluster, *svcsdk.DescribeClustersOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.ClusterParameters, *svcsdk.DescribeClustersOutput) error
	isUpToDate     func(context.Context, *svcapitypes.Cluster, *svcsdk.DescribeClustersOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.Cluster, *svcsdk.CreateClusterInput) error
	postCreate     func(context.Context, *svcapitypes.Cluster, *svcsdk.CreateClusterOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.Cluster, *svcsdk.DeleteClusterInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.Cluster, *svcsdk.DeleteClusterOutput, error) (managed.ExternalDelete, error)
	preUpdate      func(context.Context, *svcapitypes.Cluster, *svcsdk.UpdateClusterInput) error
	postUpdate     func(context.Context, *svcapitypes.Cluster, *svcsdk.UpdateClusterOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.Cluster, *svcsdk.DescribeClustersInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.Cluster, _ *svcsdk.DescribeClustersOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.ClusterParameters, *svcsdk.DescribeClustersOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.Cluster, *svcsdk.DescribeClustersOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.Cluster, *svcsdk.CreateClusterInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.Cluster, _ *svcsdk.CreateClusterOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.Cluster, *svcsdk.DeleteClusterInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.Cluster, _ *svcsdk.DeleteClusterOutput, err error) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, err
}
func nopPreUpdate(context.Context, *svcapitypes.Cluster, *svcsdk.UpdateClusterInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.Cluster, _ *svcsdk.UpdateClusterOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
