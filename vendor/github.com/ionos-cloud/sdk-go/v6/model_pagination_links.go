/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 6.0-SDK.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// PaginationLinks struct for PaginationLinks
type PaginationLinks struct {
	// URL (with offset and limit parameters) of the previous page; only present if offset is greater than 0
	Prev *string `json:"prev,omitempty"`
	// URL (with offset and limit parameters) of the current page
	Self *string `json:"self,omitempty"`
	// URL (with offset and limit parameters) of the next page; only present if offset + limit is less than the total number of elements
	Next *string `json:"next,omitempty"`
}



// GetPrev returns the Prev field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginationLinks) GetPrev() *string {
	if o == nil {
		return nil
	}


	return o.Prev

}

// GetPrevOk returns a tuple with the Prev field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationLinks) GetPrevOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Prev, true
}

// SetPrev sets field value
func (o *PaginationLinks) SetPrev(v string) {


	o.Prev = &v

}

// HasPrev returns a boolean if a field has been set.
func (o *PaginationLinks) HasPrev() bool {
	if o != nil && o.Prev != nil {
		return true
	}

	return false
}



// GetSelf returns the Self field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginationLinks) GetSelf() *string {
	if o == nil {
		return nil
	}


	return o.Self

}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationLinks) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Self, true
}

// SetSelf sets field value
func (o *PaginationLinks) SetSelf(v string) {


	o.Self = &v

}

// HasSelf returns a boolean if a field has been set.
func (o *PaginationLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}



// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginationLinks) GetNext() *string {
	if o == nil {
		return nil
	}


	return o.Next

}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationLinks) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Next, true
}

// SetNext sets field value
func (o *PaginationLinks) SetNext(v string) {


	o.Next = &v

}

// HasNext returns a boolean if a field has been set.
func (o *PaginationLinks) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}


func (o PaginationLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Prev != nil {
		toSerialize["prev"] = o.Prev
	}
	

	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	

	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	
	return json.Marshal(toSerialize)
}

type NullablePaginationLinks struct {
	value *PaginationLinks
	isSet bool
}

func (v NullablePaginationLinks) Get() *PaginationLinks {
	return v.value
}

func (v *NullablePaginationLinks) Set(val *PaginationLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationLinks(val *PaginationLinks) *NullablePaginationLinks {
	return &NullablePaginationLinks{value: val, isSet: true}
}

func (v NullablePaginationLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


