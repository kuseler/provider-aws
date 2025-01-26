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

package originaccesscontrol

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/cloudfront"
	svcsdk "github.com/aws/aws-sdk-go/service/cloudfront"
	svcsdkapi "github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/cloudfront/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an OriginAccessControl resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create OriginAccessControl in AWS"
	errUpdate        = "cannot update OriginAccessControl in AWS"
	errDescribe      = "failed to describe OriginAccessControl"
	errDelete        = "failed to delete OriginAccessControl"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.OriginAccessControl)
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
	cr, ok := mg.(*svcapitypes.OriginAccessControl)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetOriginAccessControlInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetOriginAccessControlWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateOriginAccessControl(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
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
	cr, ok := mg.(*svcapitypes.OriginAccessControl)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateOriginAccessControlInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateOriginAccessControlWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.ETag != nil {
		cr.Status.AtProvider.ETag = resp.ETag
	} else {
		cr.Status.AtProvider.ETag = nil
	}
	if resp.Location != nil {
		cr.Status.AtProvider.Location = resp.Location
	} else {
		cr.Status.AtProvider.Location = nil
	}
	if resp.OriginAccessControl != nil {
		f2 := &svcapitypes.OriginAccessControl_SDK{}
		if resp.OriginAccessControl.Id != nil {
			f2.ID = resp.OriginAccessControl.Id
		}
		if resp.OriginAccessControl.OriginAccessControlConfig != nil {
			f2f1 := &svcapitypes.OriginAccessControlConfig{}
			if resp.OriginAccessControl.OriginAccessControlConfig.Description != nil {
				f2f1.Description = resp.OriginAccessControl.OriginAccessControlConfig.Description
			}
			if resp.OriginAccessControl.OriginAccessControlConfig.Name != nil {
				f2f1.Name = resp.OriginAccessControl.OriginAccessControlConfig.Name
			}
			if resp.OriginAccessControl.OriginAccessControlConfig.OriginAccessControlOriginType != nil {
				f2f1.OriginAccessControlOriginType = resp.OriginAccessControl.OriginAccessControlConfig.OriginAccessControlOriginType
			}
			if resp.OriginAccessControl.OriginAccessControlConfig.SigningBehavior != nil {
				f2f1.SigningBehavior = resp.OriginAccessControl.OriginAccessControlConfig.SigningBehavior
			}
			if resp.OriginAccessControl.OriginAccessControlConfig.SigningProtocol != nil {
				f2f1.SigningProtocol = resp.OriginAccessControl.OriginAccessControlConfig.SigningProtocol
			}
			f2.OriginAccessControlConfig = f2f1
		}
		cr.Status.AtProvider.OriginAccessControl = f2
	} else {
		cr.Status.AtProvider.OriginAccessControl = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.OriginAccessControl)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateOriginAccessControlInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateOriginAccessControlWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) (managed.ExternalDelete, error) {
	cr, ok := mg.(*svcapitypes.OriginAccessControl)
	if !ok {
		return managed.ExternalDelete{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteOriginAccessControlInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return managed.ExternalDelete{}, errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return managed.ExternalDelete{}, nil
	}
	resp, err := e.client.DeleteOriginAccessControlWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

func (e *external) Disconnect(ctx context.Context) error {
	// Unimplemented, required by newer versions of crossplane-runtime
	return nil
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.CloudFrontAPI, opts []option) *external {
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
	client         svcsdkapi.CloudFrontAPI
	preObserve     func(context.Context, *svcapitypes.OriginAccessControl, *svcsdk.GetOriginAccessControlInput) error
	postObserve    func(context.Context, *svcapitypes.OriginAccessControl, *svcsdk.GetOriginAccessControlOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.OriginAccessControlParameters, *svcsdk.GetOriginAccessControlOutput) error
	isUpToDate     func(context.Context, *svcapitypes.OriginAccessControl, *svcsdk.GetOriginAccessControlOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.OriginAccessControl, *svcsdk.CreateOriginAccessControlInput) error
	postCreate     func(context.Context, *svcapitypes.OriginAccessControl, *svcsdk.CreateOriginAccessControlOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.OriginAccessControl, *svcsdk.DeleteOriginAccessControlInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.OriginAccessControl, *svcsdk.DeleteOriginAccessControlOutput, error) (managed.ExternalDelete, error)
	preUpdate      func(context.Context, *svcapitypes.OriginAccessControl, *svcsdk.UpdateOriginAccessControlInput) error
	postUpdate     func(context.Context, *svcapitypes.OriginAccessControl, *svcsdk.UpdateOriginAccessControlOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.OriginAccessControl, *svcsdk.GetOriginAccessControlInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.OriginAccessControl, _ *svcsdk.GetOriginAccessControlOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.OriginAccessControlParameters, *svcsdk.GetOriginAccessControlOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.OriginAccessControl, *svcsdk.GetOriginAccessControlOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.OriginAccessControl, *svcsdk.CreateOriginAccessControlInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.OriginAccessControl, _ *svcsdk.CreateOriginAccessControlOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.OriginAccessControl, *svcsdk.DeleteOriginAccessControlInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.OriginAccessControl, _ *svcsdk.DeleteOriginAccessControlOutput, err error) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, err
}
func nopPreUpdate(context.Context, *svcapitypes.OriginAccessControl, *svcsdk.UpdateOriginAccessControlInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.OriginAccessControl, _ *svcsdk.UpdateOriginAccessControlOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
