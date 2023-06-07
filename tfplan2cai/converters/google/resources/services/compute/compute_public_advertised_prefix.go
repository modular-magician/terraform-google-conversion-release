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

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ComputePublicAdvertisedPrefixAssetType string = "compute.googleapis.com/PublicAdvertisedPrefix"

func ResourceConverterComputePublicAdvertisedPrefix() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: ComputePublicAdvertisedPrefixAssetType,
		Convert:   GetComputePublicAdvertisedPrefixCaiObject,
	}
}

func GetComputePublicAdvertisedPrefixCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/publicAdvertisedPrefixes/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetComputePublicAdvertisedPrefixApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: ComputePublicAdvertisedPrefixAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "PublicAdvertisedPrefix",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetComputePublicAdvertisedPrefixApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputePublicAdvertisedPrefixDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputePublicAdvertisedPrefixName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	dnsVerificationIpProp, err := expandComputePublicAdvertisedPrefixDnsVerificationIp(d.Get("dns_verification_ip"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dns_verification_ip"); !tpgresource.IsEmptyValue(reflect.ValueOf(dnsVerificationIpProp)) && (ok || !reflect.DeepEqual(v, dnsVerificationIpProp)) {
		obj["dnsVerificationIp"] = dnsVerificationIpProp
	}
	ipCidrRangeProp, err := expandComputePublicAdvertisedPrefixIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(ipCidrRangeProp)) && (ok || !reflect.DeepEqual(v, ipCidrRangeProp)) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}

	return obj, nil
}

func expandComputePublicAdvertisedPrefixDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicAdvertisedPrefixName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicAdvertisedPrefixDnsVerificationIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicAdvertisedPrefixIpCidrRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}