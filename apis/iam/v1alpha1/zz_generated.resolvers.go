/*
Copyright 2022 YANDEX LLC
This is modified version of the software, made by the Crossplane Authors
and available at: https://github.com/crossplane-contrib/provider-jet-template

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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/yandex-cloud/crossplane-provider-yc/apis/resourcemanager/v1alpha1"
	iam "github.com/yandex-cloud/crossplane-provider-yc/config/iam"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this FolderIAMBinding.
func (mg *FolderIAMBinding) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Members),
		Extract:       iam.ServiceAccountRefValue(),
		References:    mg.Spec.ForProvider.ServiceAccountsRef,
		Selector:      mg.Spec.ForProvider.ServiceAccountsSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Members")
	}
	mg.Spec.ForProvider.Members = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.ServiceAccountsRef = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Members),
		Extract:       iam.ServiceAccountRefValue(),
		References:    mg.Spec.InitProvider.ServiceAccountsRef,
		Selector:      mg.Spec.InitProvider.ServiceAccountsSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Members")
	}
	mg.Spec.InitProvider.Members = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.ServiceAccountsRef = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this FolderIAMMember.
func (mg *FolderIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Member),
		Extract:      iam.ServiceAccountRefValue(),
		Reference:    mg.Spec.ForProvider.ServiceAccountRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Member")
	}
	mg.Spec.ForProvider.Member = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Member),
		Extract:      iam.ServiceAccountRefValue(),
		Reference:    mg.Spec.InitProvider.ServiceAccountRef,
		Selector:     mg.Spec.InitProvider.ServiceAccountSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Member")
	}
	mg.Spec.InitProvider.Member = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServiceAccount.
func (mg *ServiceAccount) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServiceAccountIAMMember.
func (mg *ServiceAccountIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Member),
		Extract:      iam.ServiceAccountRefValue(),
		Reference:    mg.Spec.ForProvider.ServiceAccountRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Member")
	}
	mg.Spec.ForProvider.Member = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountID")
	}
	mg.Spec.ForProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Member),
		Extract:      iam.ServiceAccountRefValue(),
		Reference:    mg.Spec.InitProvider.ServiceAccountRef,
		Selector:     mg.Spec.InitProvider.ServiceAccountSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Member")
	}
	mg.Spec.InitProvider.Member = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.InitProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceAccountID")
	}
	mg.Spec.InitProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServiceAccountKey.
func (mg *ServiceAccountKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountID")
	}
	mg.Spec.ForProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.InitProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceAccountID")
	}
	mg.Spec.InitProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServiceAccountStaticAccessKey.
func (mg *ServiceAccountStaticAccessKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountID")
	}
	mg.Spec.ForProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.InitProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceAccountID")
	}
	mg.Spec.InitProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountIDRef = rsp.ResolvedReference

	return nil
}