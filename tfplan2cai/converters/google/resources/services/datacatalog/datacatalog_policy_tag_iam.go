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

package datacatalog

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const DataCatalogPolicyTagIAMAssetType string = "datacatalog.googleapis.com/PolicyTag"

func ResourceConverterDataCatalogPolicyTagIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataCatalogPolicyTagIAMAssetType,
		Convert:           GetDataCatalogPolicyTagIamPolicyCaiObject,
		MergeCreateUpdate: MergeDataCatalogPolicyTagIamPolicy,
	}
}

func ResourceConverterDataCatalogPolicyTagIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataCatalogPolicyTagIAMAssetType,
		Convert:           GetDataCatalogPolicyTagIamBindingCaiObject,
		FetchFullResource: FetchDataCatalogPolicyTagIamPolicy,
		MergeCreateUpdate: MergeDataCatalogPolicyTagIamBinding,
		MergeDelete:       MergeDataCatalogPolicyTagIamBindingDelete,
	}
}

func ResourceConverterDataCatalogPolicyTagIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataCatalogPolicyTagIAMAssetType,
		Convert:           GetDataCatalogPolicyTagIamMemberCaiObject,
		FetchFullResource: FetchDataCatalogPolicyTagIamPolicy,
		MergeCreateUpdate: MergeDataCatalogPolicyTagIamMember,
		MergeDelete:       MergeDataCatalogPolicyTagIamMemberDelete,
	}
}

func GetDataCatalogPolicyTagIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataCatalogPolicyTagIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetDataCatalogPolicyTagIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataCatalogPolicyTagIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetDataCatalogPolicyTagIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataCatalogPolicyTagIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeDataCatalogPolicyTagIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDataCatalogPolicyTagIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeDataCatalogPolicyTagIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeDataCatalogPolicyTagIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeDataCatalogPolicyTagIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newDataCatalogPolicyTagIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//datacatalog.googleapis.com/{{policy_tag}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: DataCatalogPolicyTagIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDataCatalogPolicyTagIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("policy_tag"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		DataCatalogPolicyTagIamUpdaterProducer,
		d,
		config,
		"//datacatalog.googleapis.com/{{policy_tag}}",
		DataCatalogPolicyTagIAMAssetType,
	)
}
