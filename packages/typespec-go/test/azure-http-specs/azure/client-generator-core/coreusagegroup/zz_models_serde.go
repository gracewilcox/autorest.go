// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package coreusagegroup

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type InputModel.
func (i InputModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "name", i.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type InputModel.
func (i *InputModel) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &i.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OrphanModel.
func (o OrphanModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "desc", o.Description)
	populate(objectMap, "name", o.ModelName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OrphanModel.
func (o *OrphanModel) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "desc":
			err = unpopulate(val, "Description", &o.Description)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "ModelName", &o.ModelName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OutputModel.
func (o OutputModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "name", o.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OutputModel.
func (o *OutputModel) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResultModel.
func (r ResultModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "name", r.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResultModel.
func (r *ResultModel) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoundTripModel.
func (r RoundTripModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "result", r.Result)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoundTripModel.
func (r *RoundTripModel) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "result":
			err = unpopulate(val, "Result", &r.Result)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
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
