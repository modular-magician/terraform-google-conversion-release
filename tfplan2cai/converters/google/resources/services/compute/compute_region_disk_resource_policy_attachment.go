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
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeRegionDiskResourcePolicyAttachmentAssetType string = "compute.googleapis.com/RegionDiskResourcePolicyAttachment"

func ResourceConverterComputeRegionDiskResourcePolicyAttachment() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeRegionDiskResourcePolicyAttachmentAssetType,
		Convert:   GetComputeRegionDiskResourcePolicyAttachmentCaiObject,
	}
}

func GetComputeRegionDiskResourcePolicyAttachmentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/disks/{{disk}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeRegionDiskResourcePolicyAttachmentApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeRegionDiskResourcePolicyAttachmentAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "RegionDiskResourcePolicyAttachment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeRegionDiskResourcePolicyAttachmentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionDiskResourcePolicyAttachmentName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	return resourceComputeRegionDiskResourcePolicyAttachmentEncoder(d, config, obj)
}

func resourceComputeRegionDiskResourcePolicyAttachmentEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return nil, err
	}

	region, err := tpgresource.GetRegion(d, config)
	if err != nil {
		return nil, err
	}
	if region == "" {
		return nil, fmt.Errorf("region must be non-empty - set in resource or at provider-level")
	}

	obj["resourcePolicies"] = []interface{}{fmt.Sprintf("projects/%s/regions/%s/resourcePolicies/%s", project, region, obj["name"])}
	delete(obj, "name")
	return obj, nil
}

func expandComputeRegionDiskResourcePolicyAttachmentName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
