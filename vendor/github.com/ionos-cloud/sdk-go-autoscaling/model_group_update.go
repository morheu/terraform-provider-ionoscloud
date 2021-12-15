/*
 * VM Auto Scaling service (CloudAPI)
 *
 * VM Auto Scaling service enables IONOS clients to horizontally scale the number of VM instances, based on configured rules. Use Auto Scaling to ensure you will have a sufficient number of instances to handle your application loads at all times.  Create an Auto Scaling group that contains the server instances; Auto Scaling service will ensure that the number of instances in the group is always within these limits.  When target replica count is specified, Auto Scaling will maintain the set number on instances.  When scaling policies are specified, Auto Scaling will create or delete instances based on the demands of your applications. For each policy, specified scale-in and scale-out actions are performed whenever the corresponding thresholds are met.
 *
 * API version: 1-SDK.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// GroupUpdate Update request for an autoscaling group.
type GroupUpdate struct {
	Properties *GroupUpdatableProperties `json:"properties"`
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for GroupUpdatableProperties will be returned
func (o *GroupUpdate) GetProperties() *GroupUpdatableProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUpdate) GetPropertiesOk() (*GroupUpdatableProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *GroupUpdate) SetProperties(v GroupUpdatableProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *GroupUpdate) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o GroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullableGroupUpdate struct {
	value *GroupUpdate
	isSet bool
}

func (v NullableGroupUpdate) Get() *GroupUpdate {
	return v.value
}

func (v *NullableGroupUpdate) Set(val *GroupUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUpdate(val *GroupUpdate) *NullableGroupUpdate {
	return &NullableGroupUpdate{value: val, isSet: true}
}

func (v NullableGroupUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}