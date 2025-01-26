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

package rulegroupsnamespace

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/prometheusservice"
	svcsdk "github.com/aws/aws-sdk-go/service/prometheusservice"
	svcsdkapi "github.com/aws/aws-sdk-go/service/prometheusservice/prometheusserviceiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/prometheusservice/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an RuleGroupsNamespace resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create RuleGroupsNamespace in AWS"
	errUpdate        = "cannot update RuleGroupsNamespace in AWS"
	errDescribe      = "failed to describe RuleGroupsNamespace"
	errDelete        = "failed to delete RuleGroupsNamespace"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.RuleGroupsNamespace)
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
	cr, ok := mg.(*svcapitypes.RuleGroupsNamespace)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeRuleGroupsNamespaceInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeRuleGroupsNamespaceWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateRuleGroupsNamespace(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
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
	cr, ok := mg.(*svcapitypes.RuleGroupsNamespace)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateRuleGroupsNamespaceInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateRuleGroupsNamespaceWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.Arn != nil {
		cr.Status.AtProvider.ARN = resp.Arn
	} else {
		cr.Status.AtProvider.ARN = nil
	}
	if resp.Name != nil {
		cr.Spec.ForProvider.Name = resp.Name
	} else {
		cr.Spec.ForProvider.Name = nil
	}
	if resp.Status != nil {
		f2 := &svcapitypes.RuleGroupsNamespaceStatus_SDK{}
		if resp.Status.StatusCode != nil {
			f2.StatusCode = resp.Status.StatusCode
		}
		if resp.Status.StatusReason != nil {
			f2.StatusReason = resp.Status.StatusReason
		}
		cr.Status.AtProvider.Status = f2
	} else {
		cr.Status.AtProvider.Status = nil
	}
	if resp.Tags != nil {
		f3 := map[string]*string{}
		for f3key, f3valiter := range resp.Tags {
			var f3val string
			f3val = *f3valiter
			f3[f3key] = &f3val
		}
		cr.Spec.ForProvider.Tags = f3
	} else {
		cr.Spec.ForProvider.Tags = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	return e.update(ctx, mg)

}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) (managed.ExternalDelete, error) {
	cr, ok := mg.(*svcapitypes.RuleGroupsNamespace)
	if !ok {
		return managed.ExternalDelete{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteRuleGroupsNamespaceInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return managed.ExternalDelete{}, errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return managed.ExternalDelete{}, nil
	}
	resp, err := e.client.DeleteRuleGroupsNamespaceWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

func (e *external) Disconnect(ctx context.Context) error {
	// Unimplemented, required by newer versions of crossplane-runtime
	return nil
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.PrometheusServiceAPI, opts []option) *external {
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
	client         svcsdkapi.PrometheusServiceAPI
	preObserve     func(context.Context, *svcapitypes.RuleGroupsNamespace, *svcsdk.DescribeRuleGroupsNamespaceInput) error
	postObserve    func(context.Context, *svcapitypes.RuleGroupsNamespace, *svcsdk.DescribeRuleGroupsNamespaceOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.RuleGroupsNamespaceParameters, *svcsdk.DescribeRuleGroupsNamespaceOutput) error
	isUpToDate     func(context.Context, *svcapitypes.RuleGroupsNamespace, *svcsdk.DescribeRuleGroupsNamespaceOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.RuleGroupsNamespace, *svcsdk.CreateRuleGroupsNamespaceInput) error
	postCreate     func(context.Context, *svcapitypes.RuleGroupsNamespace, *svcsdk.CreateRuleGroupsNamespaceOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.RuleGroupsNamespace, *svcsdk.DeleteRuleGroupsNamespaceInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.RuleGroupsNamespace, *svcsdk.DeleteRuleGroupsNamespaceOutput, error) (managed.ExternalDelete, error)
	update         func(context.Context, cpresource.Managed) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.RuleGroupsNamespace, *svcsdk.DescribeRuleGroupsNamespaceInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.RuleGroupsNamespace, _ *svcsdk.DescribeRuleGroupsNamespaceOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.RuleGroupsNamespaceParameters, *svcsdk.DescribeRuleGroupsNamespaceOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.RuleGroupsNamespace, *svcsdk.DescribeRuleGroupsNamespaceOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.RuleGroupsNamespace, *svcsdk.CreateRuleGroupsNamespaceInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.RuleGroupsNamespace, _ *svcsdk.CreateRuleGroupsNamespaceOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.RuleGroupsNamespace, *svcsdk.DeleteRuleGroupsNamespaceInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.RuleGroupsNamespace, _ *svcsdk.DeleteRuleGroupsNamespaceOutput, err error) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, err
}
func nopUpdate(context.Context, cpresource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, nil
}
