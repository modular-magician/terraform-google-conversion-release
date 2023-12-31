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

package firestore

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const FirestoreBackupScheduleAssetType string = "firestore.googleapis.com/BackupSchedule"

func ResourceConverterFirestoreBackupSchedule() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: FirestoreBackupScheduleAssetType,
		Convert:   GetFirestoreBackupScheduleCaiObject,
	}
}

func GetFirestoreBackupScheduleCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//firestore.googleapis.com/projects/{{project}}/databases/{{database}}/backupSchedules/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetFirestoreBackupScheduleApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: FirestoreBackupScheduleAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/firestore/v1/rest",
				DiscoveryName:        "BackupSchedule",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetFirestoreBackupScheduleApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	retentionProp, err := expandFirestoreBackupScheduleRetention(d.Get("retention"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("retention"); !tpgresource.IsEmptyValue(reflect.ValueOf(retentionProp)) && (ok || !reflect.DeepEqual(v, retentionProp)) {
		obj["retention"] = retentionProp
	}
	dailyRecurrenceProp, err := expandFirestoreBackupScheduleDailyRecurrence(d.Get("daily_recurrence"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("daily_recurrence"); ok || !reflect.DeepEqual(v, dailyRecurrenceProp) {
		obj["dailyRecurrence"] = dailyRecurrenceProp
	}
	weeklyRecurrenceProp, err := expandFirestoreBackupScheduleWeeklyRecurrence(d.Get("weekly_recurrence"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("weekly_recurrence"); !tpgresource.IsEmptyValue(reflect.ValueOf(weeklyRecurrenceProp)) && (ok || !reflect.DeepEqual(v, weeklyRecurrenceProp)) {
		obj["weeklyRecurrence"] = weeklyRecurrenceProp
	}

	return obj, nil
}

func expandFirestoreBackupScheduleRetention(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreBackupScheduleDailyRecurrence(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	transformed := make(map[string]interface{})

	return transformed, nil
}

func expandFirestoreBackupScheduleWeeklyRecurrence(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDay, err := expandFirestoreBackupScheduleWeeklyRecurrenceDay(original["day"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["day"] = transformedDay
	}

	return transformed, nil
}

func expandFirestoreBackupScheduleWeeklyRecurrenceDay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
