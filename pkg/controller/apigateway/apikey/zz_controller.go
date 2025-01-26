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

package apikey

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/apigateway"
	svcsdk "github.com/aws/aws-sdk-go/service/apigateway"
	svcsdkapi "github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/apigateway/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an APIKey resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create APIKey in AWS"
	errUpdate        = "cannot update APIKey in AWS"
	errDescribe      = "failed to describe APIKey"
	errDelete        = "failed to delete APIKey"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.APIKey)
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
	cr, ok := mg.(*svcapitypes.APIKey)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetApiKeyInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetApiKeyWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateAPIKey(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
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
	cr, ok := mg.(*svcapitypes.APIKey)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateApiKeyInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateApiKeyWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.CreatedDate != nil {
		cr.Status.AtProvider.CreatedDate = &metav1.Time{*resp.CreatedDate}
	} else {
		cr.Status.AtProvider.CreatedDate = nil
	}
	if resp.CustomerId != nil {
		cr.Spec.ForProvider.CustomerID = resp.CustomerId
	} else {
		cr.Spec.ForProvider.CustomerID = nil
	}
	if resp.Description != nil {
		cr.Spec.ForProvider.Description = resp.Description
	} else {
		cr.Spec.ForProvider.Description = nil
	}
	if resp.Enabled != nil {
		cr.Spec.ForProvider.Enabled = resp.Enabled
	} else {
		cr.Spec.ForProvider.Enabled = nil
	}
	if resp.Id != nil {
		cr.Status.AtProvider.ID = resp.Id
	} else {
		cr.Status.AtProvider.ID = nil
	}
	if resp.LastUpdatedDate != nil {
		cr.Status.AtProvider.LastUpdatedDate = &metav1.Time{*resp.LastUpdatedDate}
	} else {
		cr.Status.AtProvider.LastUpdatedDate = nil
	}
	if resp.Name != nil {
		cr.Spec.ForProvider.Name = resp.Name
	} else {
		cr.Spec.ForProvider.Name = nil
	}
	if resp.StageKeys != nil {
		f7 := []*string{}
		for _, f7iter := range resp.StageKeys {
			var f7elem string
			f7elem = *f7iter
			f7 = append(f7, &f7elem)
		}
		cr.Status.AtProvider.StageKeys = f7
	} else {
		cr.Status.AtProvider.StageKeys = nil
	}
	if resp.Tags != nil {
		f8 := map[string]*string{}
		for f8key, f8valiter := range resp.Tags {
			var f8val string
			f8val = *f8valiter
			f8[f8key] = &f8val
		}
		cr.Spec.ForProvider.Tags = f8
	} else {
		cr.Spec.ForProvider.Tags = nil
	}
	if resp.Value != nil {
		cr.Spec.ForProvider.Value = resp.Value
	} else {
		cr.Spec.ForProvider.Value = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.APIKey)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateApiKeyInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateApiKeyWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) (managed.ExternalDelete, error) {
	cr, ok := mg.(*svcapitypes.APIKey)
	if !ok {
		return managed.ExternalDelete{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteApiKeyInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return managed.ExternalDelete{}, errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return managed.ExternalDelete{}, nil
	}
	resp, err := e.client.DeleteApiKeyWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

func (e *external) Disconnect(ctx context.Context) error {
	// Unimplemented, required by newer versions of crossplane-runtime
	return nil
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.APIGatewayAPI, opts []option) *external {
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
	client         svcsdkapi.APIGatewayAPI
	preObserve     func(context.Context, *svcapitypes.APIKey, *svcsdk.GetApiKeyInput) error
	postObserve    func(context.Context, *svcapitypes.APIKey, *svcsdk.ApiKey, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.APIKeyParameters, *svcsdk.ApiKey) error
	isUpToDate     func(context.Context, *svcapitypes.APIKey, *svcsdk.ApiKey) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.APIKey, *svcsdk.CreateApiKeyInput) error
	postCreate     func(context.Context, *svcapitypes.APIKey, *svcsdk.ApiKey, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.APIKey, *svcsdk.DeleteApiKeyInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.APIKey, *svcsdk.DeleteApiKeyOutput, error) (managed.ExternalDelete, error)
	preUpdate      func(context.Context, *svcapitypes.APIKey, *svcsdk.UpdateApiKeyInput) error
	postUpdate     func(context.Context, *svcapitypes.APIKey, *svcsdk.ApiKey, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.APIKey, *svcsdk.GetApiKeyInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.APIKey, _ *svcsdk.ApiKey, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.APIKeyParameters, *svcsdk.ApiKey) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.APIKey, *svcsdk.ApiKey) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.APIKey, *svcsdk.CreateApiKeyInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.APIKey, _ *svcsdk.ApiKey, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.APIKey, *svcsdk.DeleteApiKeyInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.APIKey, _ *svcsdk.DeleteApiKeyOutput, err error) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, err
}
func nopPreUpdate(context.Context, *svcapitypes.APIKey, *svcsdk.UpdateApiKeyInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.APIKey, _ *svcsdk.ApiKey, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
