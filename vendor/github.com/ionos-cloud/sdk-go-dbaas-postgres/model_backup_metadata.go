/*
 * IONOS DBaaS REST API
 *
 * An enterprise-grade Database is provided as a Service (DBaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.  The API allows you to create additional database clusters or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
	"time"
)

// BackupMetadata Metadata of the backup resource.
type BackupMetadata struct {
	// The ISO 8601 creation timestamp.
	CreatedDate *IonosTime `json:"createdDate,omitempty"`
	State       *State     `json:"state,omitempty"`
}

// GetCreatedDate returns the CreatedDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *BackupMetadata) GetCreatedDate() *time.Time {
	if o == nil {
		return nil
	}

	if o.CreatedDate == nil {
		return nil
	}
	return &o.CreatedDate.Time

}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupMetadata) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.CreatedDate == nil {
		return nil, false
	}
	return &o.CreatedDate.Time, true

}

// SetCreatedDate sets field value
func (o *BackupMetadata) SetCreatedDate(v time.Time) {

	o.CreatedDate = &IonosTime{v}

}

// HasCreatedDate returns a boolean if a field has been set.
func (o *BackupMetadata) HasCreatedDate() bool {
	if o != nil && o.CreatedDate != nil {
		return true
	}

	return false
}

// GetState returns the State field value
// If the value is explicit nil, the zero value for State will be returned
func (o *BackupMetadata) GetState() *State {
	if o == nil {
		return nil
	}

	return o.State

}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupMetadata) GetStateOk() (*State, bool) {
	if o == nil {
		return nil, false
	}

	return o.State, true
}

// SetState sets field value
func (o *BackupMetadata) SetState(v State) {

	o.State = &v

}

// HasState returns a boolean if a field has been set.
func (o *BackupMetadata) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

func (o BackupMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.CreatedDate != nil {
		toSerialize["createdDate"] = o.CreatedDate
	}

	if o.State != nil {
		toSerialize["state"] = o.State
	}

	return json.Marshal(toSerialize)
}

type NullableBackupMetadata struct {
	value *BackupMetadata
	isSet bool
}

func (v NullableBackupMetadata) Get() *BackupMetadata {
	return v.value
}

func (v *NullableBackupMetadata) Set(val *BackupMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupMetadata(val *BackupMetadata) *NullableBackupMetadata {
	return &NullableBackupMetadata{value: val, isSet: true}
}

func (v NullableBackupMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
