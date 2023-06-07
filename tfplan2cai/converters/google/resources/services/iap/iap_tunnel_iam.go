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

package iap

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const IapTunnelIAMAssetType string = "iap.googleapis.com/Tunnel"

func ResourceConverterIapTunnelIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         IapTunnelIAMAssetType,
		Convert:           GetIapTunnelIamPolicyCaiObject,
		MergeCreateUpdate: MergeIapTunnelIamPolicy,
	}
}

func ResourceConverterIapTunnelIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         IapTunnelIAMAssetType,
		Convert:           GetIapTunnelIamBindingCaiObject,
		FetchFullResource: FetchIapTunnelIamPolicy,
		MergeCreateUpdate: MergeIapTunnelIamBinding,
		MergeDelete:       MergeIapTunnelIamBindingDelete,
	}
}

func ResourceConverterIapTunnelIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         IapTunnelIAMAssetType,
		Convert:           GetIapTunnelIamMemberCaiObject,
		FetchFullResource: FetchIapTunnelIamPolicy,
		MergeCreateUpdate: MergeIapTunnelIamMember,
		MergeDelete:       MergeIapTunnelIamMemberDelete,
	}
}

func GetIapTunnelIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newIapTunnelIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetIapTunnelIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newIapTunnelIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetIapTunnelIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newIapTunnelIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeIapTunnelIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeIapTunnelIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeIapTunnelIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeIapTunnelIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeIapTunnelIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newIapTunnelIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//iap.googleapis.com/projects/{{project}}/iap_tunnel")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: IapTunnelIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchIapTunnelIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value

	return tpgiamresource.FetchIamPolicy(
		IapTunnelIamUpdaterProducer,
		d,
		config,
		"//iap.googleapis.com/projects/{{project}}/iap_tunnel",
		IapTunnelIAMAssetType,
	)
}