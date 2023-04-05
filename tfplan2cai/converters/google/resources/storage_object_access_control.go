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

package google

import "reflect"

const StorageObjectAccessControlAssetType string = "storage.googleapis.com/ObjectAccessControl"

func resourceConverterStorageObjectAccessControl() ResourceConverter {
	return ResourceConverter{
		AssetType: StorageObjectAccessControlAssetType,
		Convert:   GetStorageObjectAccessControlCaiObject,
	}
}

func GetStorageObjectAccessControlCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//storage.googleapis.com/b/{{bucket}}/o/{{%object}}/acl/{{entity}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetStorageObjectAccessControlApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: StorageObjectAccessControlAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/storage/v1/rest",
				DiscoveryName:        "ObjectAccessControl",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetStorageObjectAccessControlApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	bucketProp, err := expandStorageObjectAccessControlBucket(d.Get("bucket"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bucket"); !isEmptyValue(reflect.ValueOf(bucketProp)) && (ok || !reflect.DeepEqual(v, bucketProp)) {
		obj["bucket"] = bucketProp
	}
	entityProp, err := expandStorageObjectAccessControlEntity(d.Get("entity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("entity"); !isEmptyValue(reflect.ValueOf(entityProp)) && (ok || !reflect.DeepEqual(v, entityProp)) {
		obj["entity"] = entityProp
	}
	objectProp, err := expandStorageObjectAccessControlObject(d.Get("object"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("object"); !isEmptyValue(reflect.ValueOf(objectProp)) && (ok || !reflect.DeepEqual(v, objectProp)) {
		obj["object"] = objectProp
	}
	roleProp, err := expandStorageObjectAccessControlRole(d.Get("role"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("role"); !isEmptyValue(reflect.ValueOf(roleProp)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}

	return obj, nil
}

func expandStorageObjectAccessControlBucket(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageObjectAccessControlEntity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageObjectAccessControlObject(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageObjectAccessControlRole(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}