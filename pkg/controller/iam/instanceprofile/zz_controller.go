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

package instanceprofile

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/iam"
	svcsdk "github.com/aws/aws-sdk-go/service/iam"
	svcsdkapi "github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/iam/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an InstanceProfile resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create InstanceProfile in AWS"
	errUpdate        = "cannot update InstanceProfile in AWS"
	errDescribe      = "failed to describe InstanceProfile"
	errDelete        = "failed to delete InstanceProfile"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	sess, err := connectaws.GetConfigV1(ctx, c.kube, mg, connectaws.GlobalRegion)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.InstanceProfile)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetInstanceProfileInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetInstanceProfileWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateInstanceProfile(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
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
	cr, ok := mg.(*svcapitypes.InstanceProfile)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateInstanceProfileInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateInstanceProfileWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.InstanceProfile.Arn != nil {
		cr.Status.AtProvider.ARN = resp.InstanceProfile.Arn
	} else {
		cr.Status.AtProvider.ARN = nil
	}
	if resp.InstanceProfile.CreateDate != nil {
		cr.Status.AtProvider.CreateDate = &metav1.Time{*resp.InstanceProfile.CreateDate}
	} else {
		cr.Status.AtProvider.CreateDate = nil
	}
	if resp.InstanceProfile.InstanceProfileId != nil {
		cr.Status.AtProvider.InstanceProfileID = resp.InstanceProfile.InstanceProfileId
	} else {
		cr.Status.AtProvider.InstanceProfileID = nil
	}
	if resp.InstanceProfile.InstanceProfileName != nil {
		cr.Status.AtProvider.InstanceProfileName = resp.InstanceProfile.InstanceProfileName
	} else {
		cr.Status.AtProvider.InstanceProfileName = nil
	}
	if resp.InstanceProfile.Path != nil {
		cr.Spec.ForProvider.Path = resp.InstanceProfile.Path
	} else {
		cr.Spec.ForProvider.Path = nil
	}
	if resp.InstanceProfile.Roles != nil {
		f5 := []*svcapitypes.Role{}
		for _, f5iter := range resp.InstanceProfile.Roles {
			f5elem := &svcapitypes.Role{}
			if f5iter.Arn != nil {
				f5elem.ARN = f5iter.Arn
			}
			if f5iter.AssumeRolePolicyDocument != nil {
				f5elem.AssumeRolePolicyDocument = f5iter.AssumeRolePolicyDocument
			}
			if f5iter.CreateDate != nil {
				f5elem.CreateDate = &metav1.Time{*f5iter.CreateDate}
			}
			if f5iter.Description != nil {
				f5elem.Description = f5iter.Description
			}
			if f5iter.MaxSessionDuration != nil {
				f5elem.MaxSessionDuration = f5iter.MaxSessionDuration
			}
			if f5iter.Path != nil {
				f5elem.Path = f5iter.Path
			}
			if f5iter.PermissionsBoundary != nil {
				f5elemf6 := &svcapitypes.AttachedPermissionsBoundary{}
				if f5iter.PermissionsBoundary.PermissionsBoundaryArn != nil {
					f5elemf6.PermissionsBoundaryARN = f5iter.PermissionsBoundary.PermissionsBoundaryArn
				}
				if f5iter.PermissionsBoundary.PermissionsBoundaryType != nil {
					f5elemf6.PermissionsBoundaryType = f5iter.PermissionsBoundary.PermissionsBoundaryType
				}
				f5elem.PermissionsBoundary = f5elemf6
			}
			if f5iter.RoleId != nil {
				f5elem.RoleID = f5iter.RoleId
			}
			if f5iter.RoleLastUsed != nil {
				f5elemf8 := &svcapitypes.RoleLastUsed{}
				if f5iter.RoleLastUsed.LastUsedDate != nil {
					f5elemf8.LastUsedDate = &metav1.Time{*f5iter.RoleLastUsed.LastUsedDate}
				}
				if f5iter.RoleLastUsed.Region != nil {
					f5elemf8.Region = f5iter.RoleLastUsed.Region
				}
				f5elem.RoleLastUsed = f5elemf8
			}
			if f5iter.RoleName != nil {
				f5elem.RoleName = f5iter.RoleName
			}
			if f5iter.Tags != nil {
				f5elemf10 := []*svcapitypes.Tag{}
				for _, f5elemf10iter := range f5iter.Tags {
					f5elemf10elem := &svcapitypes.Tag{}
					if f5elemf10iter.Key != nil {
						f5elemf10elem.Key = f5elemf10iter.Key
					}
					if f5elemf10iter.Value != nil {
						f5elemf10elem.Value = f5elemf10iter.Value
					}
					f5elemf10 = append(f5elemf10, f5elemf10elem)
				}
				f5elem.Tags = f5elemf10
			}
			f5 = append(f5, f5elem)
		}
		cr.Status.AtProvider.Roles = f5
	} else {
		cr.Status.AtProvider.Roles = nil
	}
	if resp.InstanceProfile.Tags != nil {
		f6 := []*svcapitypes.Tag{}
		for _, f6iter := range resp.InstanceProfile.Tags {
			f6elem := &svcapitypes.Tag{}
			if f6iter.Key != nil {
				f6elem.Key = f6iter.Key
			}
			if f6iter.Value != nil {
				f6elem.Value = f6iter.Value
			}
			f6 = append(f6, f6elem)
		}
		cr.Spec.ForProvider.Tags = f6
	} else {
		cr.Spec.ForProvider.Tags = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	return e.update(ctx, mg)

}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) (managed.ExternalDelete, error) {
	cr, ok := mg.(*svcapitypes.InstanceProfile)
	if !ok {
		return managed.ExternalDelete{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteInstanceProfileInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return managed.ExternalDelete{}, errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return managed.ExternalDelete{}, nil
	}
	resp, err := e.client.DeleteInstanceProfileWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

func (e *external) Disconnect(ctx context.Context) error {
	// Unimplemented, required by newer versions of crossplane-runtime
	return nil
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.IAMAPI, opts []option) *external {
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
		update:         nopUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.IAMAPI
	preObserve     func(context.Context, *svcapitypes.InstanceProfile, *svcsdk.GetInstanceProfileInput) error
	postObserve    func(context.Context, *svcapitypes.InstanceProfile, *svcsdk.GetInstanceProfileOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.InstanceProfileParameters, *svcsdk.GetInstanceProfileOutput) error
	isUpToDate     func(context.Context, *svcapitypes.InstanceProfile, *svcsdk.GetInstanceProfileOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.InstanceProfile, *svcsdk.CreateInstanceProfileInput) error
	postCreate     func(context.Context, *svcapitypes.InstanceProfile, *svcsdk.CreateInstanceProfileOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.InstanceProfile, *svcsdk.DeleteInstanceProfileInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.InstanceProfile, *svcsdk.DeleteInstanceProfileOutput, error) (managed.ExternalDelete, error)
	update         func(context.Context, cpresource.Managed) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.InstanceProfile, *svcsdk.GetInstanceProfileInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.InstanceProfile, _ *svcsdk.GetInstanceProfileOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.InstanceProfileParameters, *svcsdk.GetInstanceProfileOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.InstanceProfile, *svcsdk.GetInstanceProfileOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.InstanceProfile, *svcsdk.CreateInstanceProfileInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.InstanceProfile, _ *svcsdk.CreateInstanceProfileOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.InstanceProfile, *svcsdk.DeleteInstanceProfileInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.InstanceProfile, _ *svcsdk.DeleteInstanceProfileOutput, err error) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, err
}
func nopUpdate(context.Context, cpresource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, nil
}
