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

package corebilling

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const CoreBillingProjectInfoAssetType string = "cloudbilling.googleapis.com/ProjectInfo"

func ResourceConverterCoreBillingProjectInfo() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: CoreBillingProjectInfoAssetType,
		Convert:   GetCoreBillingProjectInfoCaiObject,
	}
}

func GetCoreBillingProjectInfoCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//cloudbilling.googleapis.com/projects/{{project}}/billingInfo/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetCoreBillingProjectInfoApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: CoreBillingProjectInfoAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudbilling/v1/rest",
				DiscoveryName:        "ProjectInfo",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetCoreBillingProjectInfoApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	billing_accountProp, err := expandCoreBillingProjectInfoBillingAccount(d.Get("billing_account"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("billing_account"); !tpgresource.IsEmptyValue(reflect.ValueOf(billing_accountProp)) && (ok || !reflect.DeepEqual(v, billing_accountProp)) {
		obj["billing_account"] = billing_accountProp
	}

	return resourceCoreBillingProjectInfoEncoder(d, config, obj)
}

func resourceCoreBillingProjectInfoEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	ba := d.Get("billing_account").(string)
	if ba == "" {
		obj["billingAccountName"] = ""
	} else {
		obj["billingAccountName"] = "billingAccounts/" + ba
	}
	delete(obj, "billing_account")
	return obj, nil
}

func expandCoreBillingProjectInfoBillingAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
