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

package policy

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/iot"
	svcsdk "github.com/aws/aws-sdk-go/service/iot"
	svcsdkapi "github.com/aws/aws-sdk-go/service/iot/iotiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/iot/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an Policy resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create Policy in AWS"
	errUpdate        = "cannot update Policy in AWS"
	errDescribe      = "failed to describe Policy"
	errDelete        = "failed to delete Policy"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.Policy)
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
	cr, ok := mg.(*svcapitypes.Policy)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetPolicyInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetPolicyWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GeneratePolicy(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
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
	cr, ok := mg.(*svcapitypes.Policy)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreatePolicyInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreatePolicyWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.PolicyArn != nil {
		cr.Status.AtProvider.PolicyARN = resp.PolicyArn
	} else {
		cr.Status.AtProvider.PolicyARN = nil
	}
	if resp.PolicyDocument != nil {
		cr.Spec.ForProvider.PolicyDocument = resp.PolicyDocument
	} else {
		cr.Spec.ForProvider.PolicyDocument = nil
	}
	if resp.PolicyVersionId != nil {
		cr.Status.AtProvider.PolicyVersionID = resp.PolicyVersionId
	} else {
		cr.Status.AtProvider.PolicyVersionID = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	return e.update(ctx, mg)

}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) (managed.ExternalDelete, error) {
	cr, ok := mg.(*svcapitypes.Policy)
	if !ok {
		return managed.ExternalDelete{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeletePolicyInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return managed.ExternalDelete{}, errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return managed.ExternalDelete{}, nil
	}
	resp, err := e.client.DeletePolicyWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

func (e *external) Disconnect(ctx context.Context) error {
	// Unimplemented, required by newer versions of crossplane-runtime
	return nil
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.IoTAPI, opts []option) *external {
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
	client         svcsdkapi.IoTAPI
	preObserve     func(context.Context, *svcapitypes.Policy, *svcsdk.GetPolicyInput) error
	postObserve    func(context.Context, *svcapitypes.Policy, *svcsdk.GetPolicyOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.PolicyParameters, *svcsdk.GetPolicyOutput) error
	isUpToDate     func(context.Context, *svcapitypes.Policy, *svcsdk.GetPolicyOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.Policy, *svcsdk.CreatePolicyInput) error
	postCreate     func(context.Context, *svcapitypes.Policy, *svcsdk.CreatePolicyOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.Policy, *svcsdk.DeletePolicyInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.Policy, *svcsdk.DeletePolicyOutput, error) (managed.ExternalDelete, error)
	update         func(context.Context, cpresource.Managed) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.Policy, *svcsdk.GetPolicyInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.Policy, _ *svcsdk.GetPolicyOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.PolicyParameters, *svcsdk.GetPolicyOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.Policy, *svcsdk.GetPolicyOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.Policy, *svcsdk.CreatePolicyInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.Policy, _ *svcsdk.CreatePolicyOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.Policy, *svcsdk.DeletePolicyInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.Policy, _ *svcsdk.DeletePolicyOutput, err error) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, err
}
func nopUpdate(context.Context, cpresource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, nil
}
