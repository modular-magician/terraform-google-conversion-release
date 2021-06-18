// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import "fmt"

func GetIapWebTypeComputeIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newIapWebTypeComputeIamAsset(d, config, expandIamPolicyBindings)
}

func GetIapWebTypeComputeIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newIapWebTypeComputeIamAsset(d, config, expandIamRoleBindings)
}

func GetIapWebTypeComputeIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newIapWebTypeComputeIamAsset(d, config, expandIamMemberBindings)
}

func MergeIapWebTypeComputeIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeIapWebTypeComputeIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeIapWebTypeComputeIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeIapWebTypeComputeIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeIapWebTypeComputeIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newIapWebTypeComputeIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//iap.googleapis.com/{{webtypecompute}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: "iap.googleapis.com/WebTypeCompute",
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchIapWebTypeComputeIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("{{webtypecompute}}"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		IapWebTypeComputeIamUpdaterProducer,
		d,
		config,
		"//iap.googleapis.com/{{webtypecompute}}",
		"iap.googleapis.com/WebTypeCompute",
	)
}
