// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package cloudrunv2

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const CloudRunV2JobIAMAssetType string = "run.googleapis.com/Job"

func ResourceConverterCloudRunV2JobIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         CloudRunV2JobIAMAssetType,
		Convert:           GetCloudRunV2JobIamPolicyCaiObject,
		MergeCreateUpdate: MergeCloudRunV2JobIamPolicy,
	}
}

func ResourceConverterCloudRunV2JobIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         CloudRunV2JobIAMAssetType,
		Convert:           GetCloudRunV2JobIamBindingCaiObject,
		FetchFullResource: FetchCloudRunV2JobIamPolicy,
		MergeCreateUpdate: MergeCloudRunV2JobIamBinding,
		MergeDelete:       MergeCloudRunV2JobIamBindingDelete,
	}
}

func ResourceConverterCloudRunV2JobIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         CloudRunV2JobIAMAssetType,
		Convert:           GetCloudRunV2JobIamMemberCaiObject,
		FetchFullResource: FetchCloudRunV2JobIamPolicy,
		MergeCreateUpdate: MergeCloudRunV2JobIamMember,
		MergeDelete:       MergeCloudRunV2JobIamMemberDelete,
	}
}

func GetCloudRunV2JobIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newCloudRunV2JobIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetCloudRunV2JobIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newCloudRunV2JobIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetCloudRunV2JobIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newCloudRunV2JobIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeCloudRunV2JobIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeCloudRunV2JobIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeCloudRunV2JobIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeCloudRunV2JobIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeCloudRunV2JobIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newCloudRunV2JobIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//run.googleapis.com/projects/{{project}}/locations/{{location}}/jobs/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: CloudRunV2JobIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchCloudRunV2JobIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("name"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		CloudRunV2JobIamUpdaterProducer,
		d,
		config,
		"//run.googleapis.com/projects/{{project}}/locations/{{location}}/jobs/{{name}}",
		CloudRunV2JobIAMAssetType,
	)
}
