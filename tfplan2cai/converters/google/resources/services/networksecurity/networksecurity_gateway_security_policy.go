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

package networksecurity

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkSecurityGatewaySecurityPolicyAssetType string = "networksecurity.googleapis.com/GatewaySecurityPolicy"

func ResourceConverterNetworkSecurityGatewaySecurityPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkSecurityGatewaySecurityPolicyAssetType,
		Convert:   GetNetworkSecurityGatewaySecurityPolicyCaiObject,
	}
}

func GetNetworkSecurityGatewaySecurityPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networksecurity.googleapis.com/projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkSecurityGatewaySecurityPolicyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkSecurityGatewaySecurityPolicyAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networksecurity/v1beta1/rest",
				DiscoveryName:        "GatewaySecurityPolicy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkSecurityGatewaySecurityPolicyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityGatewaySecurityPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	tlsInspectionPolicyProp, err := expandNetworkSecurityGatewaySecurityPolicyTlsInspectionPolicy(d.Get("tls_inspection_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tls_inspection_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(tlsInspectionPolicyProp)) && (ok || !reflect.DeepEqual(v, tlsInspectionPolicyProp)) {
		obj["tlsInspectionPolicy"] = tlsInspectionPolicyProp
	}

	return obj, nil
}

func expandNetworkSecurityGatewaySecurityPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityGatewaySecurityPolicyTlsInspectionPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
