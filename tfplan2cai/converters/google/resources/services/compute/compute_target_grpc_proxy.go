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

package compute

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeTargetGrpcProxyAssetType string = "compute.googleapis.com/TargetGrpcProxy"

func ResourceConverterComputeTargetGrpcProxy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeTargetGrpcProxyAssetType,
		Convert:   GetComputeTargetGrpcProxyCaiObject,
	}
}

func GetComputeTargetGrpcProxyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/targetGrpcProxies/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeTargetGrpcProxyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeTargetGrpcProxyAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "TargetGrpcProxy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeTargetGrpcProxyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeTargetGrpcProxyName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeTargetGrpcProxyDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	urlMapProp, err := expandComputeTargetGrpcProxyUrlMap(d.Get("url_map"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("url_map"); !tpgresource.IsEmptyValue(reflect.ValueOf(urlMapProp)) && (ok || !reflect.DeepEqual(v, urlMapProp)) {
		obj["urlMap"] = urlMapProp
	}
	validateForProxylessProp, err := expandComputeTargetGrpcProxyValidateForProxyless(d.Get("validate_for_proxyless"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("validate_for_proxyless"); !tpgresource.IsEmptyValue(reflect.ValueOf(validateForProxylessProp)) && (ok || !reflect.DeepEqual(v, validateForProxylessProp)) {
		obj["validateForProxyless"] = validateForProxylessProp
	}
	fingerprintProp, err := expandComputeTargetGrpcProxyFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}

	return obj, nil
}

func expandComputeTargetGrpcProxyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetGrpcProxyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetGrpcProxyUrlMap(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetGrpcProxyValidateForProxyless(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetGrpcProxyFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
