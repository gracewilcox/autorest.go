// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package visibilitygroup

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ReadOnlyModel.
func (r ReadOnlyModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "optionalNullableIntList", r.OptionalNullableIntList)
	populate(objectMap, "optionalStringRecord", r.OptionalStringRecord)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReadOnlyModel.
func (r *ReadOnlyModel) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "optionalNullableIntList":
			err = unpopulate(val, "OptionalNullableIntList", &r.OptionalNullableIntList)
			delete(rawMsg, key)
		case "optionalStringRecord":
			err = unpopulate(val, "OptionalStringRecord", &r.OptionalStringRecord)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VisibilityModel.
func (v VisibilityModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "createProp", v.CreateProp)
	populate(objectMap, "deleteProp", v.DeleteProp)
	populate(objectMap, "readProp", v.ReadProp)
	populate(objectMap, "updateProp", v.UpdateProp)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VisibilityModel.
func (v *VisibilityModel) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createProp":
			err = unpopulate(val, "CreateProp", &v.CreateProp)
			delete(rawMsg, key)
		case "deleteProp":
			err = unpopulate(val, "DeleteProp", &v.DeleteProp)
			delete(rawMsg, key)
		case "readProp":
			err = unpopulate(val, "ReadProp", &v.ReadProp)
			delete(rawMsg, key)
		case "updateProp":
			err = unpopulate(val, "UpdateProp", &v.UpdateProp)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
