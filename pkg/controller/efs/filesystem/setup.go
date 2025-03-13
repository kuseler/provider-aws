package filesystem

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/pkg/errors"
	"k8s.io/utils/ptr"

	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/efs"
	svcsdkapi "github.com/aws/aws-sdk-go/service/efs/efsiface"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/connection"
	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/efs/v1alpha1"
	"github.com/crossplane-contrib/provider-aws/apis/v1alpha1"
	"github.com/crossplane-contrib/provider-aws/pkg/features"
	"github.com/crossplane-contrib/provider-aws/pkg/utils/pointer"
	custommanaged "github.com/crossplane-contrib/provider-aws/pkg/utils/reconciler/managed"
)

// SetupFileSystem adds a controller that reconciles FileSystem.
func SetupFileSystem(mgr ctrl.Manager, o controller.Options) error {
	name := managed.ControllerName(svcapitypes.FileSystemGroupKind)
	opts := []option{
		func(e *external) {
			c := &custom{kube: mgr.GetClient(), client: e.client, external: e}
			e.isUpToDate = c.isUpToDate
			e.preCreate = preCreate
			e.postCreate = postCreate
			e.preObserve = preObserve
			e.preUpdate = c.preUpdate
			e.preDelete = preDelete
			e.postObserve = c.postObserve
		},
	}

	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.Features.Enabled(features.EnableAlphaExternalSecretStores) {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), v1alpha1.StoreConfigGroupVersionKind))
	}

	reconcilerOpts := []managed.ReconcilerOption{
		managed.WithInitializers(),
		managed.WithCriticalAnnotationUpdater(custommanaged.NewRetryingCriticalAnnotationUpdater(mgr.GetClient())),
		managed.WithTypedExternalConnector(&connector{kube: mgr.GetClient(), opts: opts}),
		managed.WithPollInterval(o.PollInterval),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithConnectionPublishers(cps...),
	}

	if o.Features.Enabled(features.EnableAlphaManagementPolicies) {
		reconcilerOpts = append(reconcilerOpts, managed.WithManagementPolicies())
	}

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(svcapitypes.FileSystemGroupVersionKind),
		reconcilerOpts...)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		WithEventFilter(resource.DesiredStateChanged()).
		For(&svcapitypes.FileSystem{}).
		Complete(r)
}

type custom struct {
	kube     client.Client
	client   svcsdkapi.EFSAPI
	external *external

	cache struct {
		backupPolicy bool
	}
}

func (e *custom) cacheBackupPolicyHelper(filesystemId *string) error {
	// Get BackupPolicy
	// default is no backup
	policy, err := e.client.DescribeBackupPolicy(&svcsdk.DescribeBackupPolicyInput{FileSystemId: filesystemId})
	if err != nil {
		e.cache.backupPolicy = !backupPolicyIsNotFound(err)
		return resource.Ignore(backupPolicyIsNotFound, err)
	}
	policyStatus := aws.StringValue(policy.BackupPolicy.Status)
	switch policyStatus {
	case "ENABLED", "ENABLING":
		e.cache.backupPolicy = true
	case "DISABLED", "DISABLING":
		e.cache.backupPolicy = false
	}
	return nil
}

func (e *custom) isUpToDate(_ context.Context, cr *svcapitypes.FileSystem, obj *svcsdk.DescribeFileSystemsOutput) (bool, string, error) {
	for _, res := range obj.FileSystems {
		if pointer.Int64Value(cr.Spec.ForProvider.ProvisionedThroughputInMibps) != int64(aws.Float64Value(res.ProvisionedThroughputInMibps)) {
			return false, "", nil
		}
		if !ptr.Equal(cr.Spec.ForProvider.ThroughputMode, res.ThroughputMode) {
			return false, "", nil
		}

	}
	return true, "", nil
}

func (e *custom) updateBackupPolicy(ctx context.Context, cr *svcapitypes.FileSystem, obj *svcsdk.DescribeFileSystemsOutput) error {
	for _, res := range obj.FileSystems {
		err := e.cacheBackupPolicyHelper(res.FileSystemId)
		if err != nil {
			return errors.Wrap(err, "failed to cache backup policy")
		}

		if e.cache.backupPolicy != *cr.Spec.ForProvider.Backup {
			var policy string
			if pointer.BoolValue(cr.Spec.ForProvider.Backup) {
				policy = "ENABLED"
			} else {
				policy = "DISABLED"
			}

			if pointer.StringValue(res.FileSystemId) == "" {
				res.FileSystemId = pointer.ToOrNilIfZeroValue(meta.GetExternalName(cr))
			}
			_, err := e.client.PutBackupPolicyWithContext(ctx,
				&svcsdk.PutBackupPolicyInput{
					FileSystemId: res.FileSystemId,
					BackupPolicy: &svcsdk.BackupPolicy{Status: &policy},
				})

			if err != nil {
				return errors.Wrap(err, "failed to put backup policy")
			}
		}
	}
	return nil
}

func preObserve(_ context.Context, cr *svcapitypes.FileSystem, obj *svcsdk.DescribeFileSystemsInput) error {
	// Describe query doesn't allow both CreationToken and FileSystemId to be given.
	obj.CreationToken = nil
	obj.FileSystemId = pointer.ToOrNilIfZeroValue(meta.GetExternalName(cr))
	return nil
}

func (e *custom) postObserve(ctx context.Context, cr *svcapitypes.FileSystem, obj *svcsdk.DescribeFileSystemsOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	if err != nil {
		return managed.ExternalObservation{}, err
	}
	if pointer.StringValue(obj.FileSystems[0].LifeCycleState) == string(svcapitypes.LifeCycleState_available) {
		cr.SetConditions(xpv1.Available())
	}
	obs.ConnectionDetails = managed.ConnectionDetails{
		svcapitypes.ResourceCredentialsSecretIDKey: []byte(meta.GetExternalName(cr)),
	}
	updateBackupPolicyErr := e.updateBackupPolicy(ctx, cr, obj)
	if updateBackupPolicyErr != nil {
		return managed.ExternalObservation{}, updateBackupPolicyErr
	}
	return obs, nil
}

func (e *custom) preUpdate(ctx context.Context, cr *svcapitypes.FileSystem, obj *svcsdk.UpdateFileSystemInput) error {
	obj.FileSystemId = pointer.ToOrNilIfZeroValue(meta.GetExternalName(cr))
	// Type of this field is *float64 but in practice, only integer values are allowed.
	if cr.Spec.ForProvider.ProvisionedThroughputInMibps != nil {
		obj.ProvisionedThroughputInMibps = aws.Float64(float64(pointer.Int64Value(cr.Spec.ForProvider.ProvisionedThroughputInMibps)))
	}

	if pointer.BoolValue(cr.Spec.ForProvider.Backup) != e.cache.backupPolicy {
		var policy string
		if pointer.BoolValue(cr.Spec.ForProvider.Backup) {
			policy = "ENABLED"
		} else {
			policy = "DISABLED"
		}

		if pointer.StringValue(obj.FileSystemId) == "" {
			obj.FileSystemId = pointer.ToOrNilIfZeroValue(meta.GetExternalName(cr))
		}
		_, err := e.client.PutBackupPolicyWithContext(ctx,
			&svcsdk.PutBackupPolicyInput{
				FileSystemId: obj.FileSystemId,
				BackupPolicy: &svcsdk.BackupPolicy{Status: &policy},
			})

		if err != nil {
			return errors.Wrap(err, "failed to put backup policy")
		}
	}

	/*	Backup is not a direct part of the FileSystem, but an externally managed BackupPolicy.
		Changing only the backup policy will result in an empty request and thus an Error.
		To avoid this, we return an error if the BackupPolicy is the only thing that has changed.
	*/
	if ptr.Equal(cr.Spec.ForProvider.ThroughputMode, obj.ThroughputMode) && pointer.Int64Value(cr.Spec.ForProvider.ProvisionedThroughputInMibps) == int64(aws.Float64Value(obj.ProvisionedThroughputInMibps)) {
		return errors.New("throughputmode and provisionedthroughputinmibps are unchanged, aborting update")
	}

	return nil
}

func preDelete(_ context.Context, cr *svcapitypes.FileSystem, obj *svcsdk.DeleteFileSystemInput) (bool, error) {
	obj.FileSystemId = pointer.ToOrNilIfZeroValue(meta.GetExternalName(cr))
	return false, nil
}

func preCreate(_ context.Context, cr *svcapitypes.FileSystem, obj *svcsdk.CreateFileSystemInput) error {
	obj.CreationToken = pointer.ToOrNilIfZeroValue(string(cr.UID))
	// Type of this field is *float64 but in practice, only integer values are allowed.
	if cr.Spec.ForProvider.ProvisionedThroughputInMibps != nil {
		obj.ProvisionedThroughputInMibps = aws.Float64(float64(pointer.Int64Value(cr.Spec.ForProvider.ProvisionedThroughputInMibps)))
	}
	return nil
}

func postCreate(_ context.Context, cr *svcapitypes.FileSystem, obj *svcsdk.FileSystemDescription, _ managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	if err != nil {
		return managed.ExternalCreation{}, err
	}
	meta.SetExternalName(cr, pointer.StringValue(obj.FileSystemId))
	return managed.ExternalCreation{}, nil
}

func backupPolicyIsNotFound(err error) bool {
	// default BackupPolicy is false, so if none is found, it's false
	var awsErr awserr.Error
	ok := errors.As(err, &awsErr)
	return ok && awsErr.Code() == "PolicyNotFound"
}
