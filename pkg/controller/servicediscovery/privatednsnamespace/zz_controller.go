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

package privatednsnamespace

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/servicediscovery"
	svcsdk "github.com/aws/aws-sdk-go/service/servicediscovery"
	svcsdkapi "github.com/aws/aws-sdk-go/service/servicediscovery/servicediscoveryiface"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/servicediscovery/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an PrivateDNSNamespace resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create PrivateDNSNamespace in AWS"
	errUpdate        = "cannot update PrivateDNSNamespace in AWS"
	errDescribe      = "failed to describe PrivateDNSNamespace"
	errDelete        = "failed to delete PrivateDNSNamespace"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.PrivateDNSNamespace)
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
	return e.observe(ctx, mg)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.PrivateDNSNamespace)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreatePrivateDnsNamespaceInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreatePrivateDnsNamespaceWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.OperationId != nil {
		cr.Status.AtProvider.OperationID = resp.OperationId
	} else {
		cr.Status.AtProvider.OperationID = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.PrivateDNSNamespace)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdatePrivateDnsNamespaceInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdatePrivateDnsNamespaceWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) (managed.ExternalDelete, error) {
	cr, ok := mg.(*svcapitypes.PrivateDNSNamespace)
	if !ok {
		return managed.ExternalDelete{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	return e.delete(ctx, mg)

}

func (e *external) Disconnect(ctx context.Context) error {
	// Unimplemented, required by newer versions of crossplane-runtime
	return nil
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.ServiceDiscoveryAPI, opts []option) *external {
	e := &external{
		kube:       kube,
		client:     client,
		observe:    nopObserve,
		preCreate:  nopPreCreate,
		postCreate: nopPostCreate,
		delete:     nopDelete,
		preUpdate:  nopPreUpdate,
		postUpdate: nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube       client.Client
	client     svcsdkapi.ServiceDiscoveryAPI
	observe    func(context.Context, cpresource.Managed) (managed.ExternalObservation, error)
	preCreate  func(context.Context, *svcapitypes.PrivateDNSNamespace, *svcsdk.CreatePrivateDnsNamespaceInput) error
	postCreate func(context.Context, *svcapitypes.PrivateDNSNamespace, *svcsdk.CreatePrivateDnsNamespaceOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	delete     func(context.Context, cpresource.Managed) (managed.ExternalDelete, error)
	preUpdate  func(context.Context, *svcapitypes.PrivateDNSNamespace, *svcsdk.UpdatePrivateDnsNamespaceInput) error
	postUpdate func(context.Context, *svcapitypes.PrivateDNSNamespace, *svcsdk.UpdatePrivateDnsNamespaceOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopObserve(context.Context, cpresource.Managed) (managed.ExternalObservation, error) {
	return managed.ExternalObservation{}, nil
}

func nopPreCreate(context.Context, *svcapitypes.PrivateDNSNamespace, *svcsdk.CreatePrivateDnsNamespaceInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.PrivateDNSNamespace, _ *svcsdk.CreatePrivateDnsNamespaceOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopDelete(context.Context, cpresource.Managed) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, nil
}
func nopPreUpdate(context.Context, *svcapitypes.PrivateDNSNamespace, *svcsdk.UpdatePrivateDnsNamespaceInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.PrivateDNSNamespace, _ *svcsdk.UpdatePrivateDnsNamespaceOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
