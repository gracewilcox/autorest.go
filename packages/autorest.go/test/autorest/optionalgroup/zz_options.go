// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package optionalgroup

import "io"

// ExplicitClientPostOptionalArrayHeaderOptions contains the optional parameters for the ExplicitClient.PostOptionalArrayHeader
// method.
type ExplicitClientPostOptionalArrayHeaderOptions struct {
	HeaderParameter []string
}

// ExplicitClientPostOptionalArrayParameterOptions contains the optional parameters for the ExplicitClient.PostOptionalArrayParameter
// method.
type ExplicitClientPostOptionalArrayParameterOptions struct {
	BodyParameter []*string
}

// ExplicitClientPostOptionalArrayPropertyOptions contains the optional parameters for the ExplicitClient.PostOptionalArrayProperty
// method.
type ExplicitClientPostOptionalArrayPropertyOptions struct {
	BodyParameter *ArrayOptionalWrapper
}

// ExplicitClientPostOptionalClassParameterOptions contains the optional parameters for the ExplicitClient.PostOptionalClassParameter
// method.
type ExplicitClientPostOptionalClassParameterOptions struct {
	BodyParameter *Product
}

// ExplicitClientPostOptionalClassPropertyOptions contains the optional parameters for the ExplicitClient.PostOptionalClassProperty
// method.
type ExplicitClientPostOptionalClassPropertyOptions struct {
	BodyParameter *ClassOptionalWrapper
}

// ExplicitClientPostOptionalIntegerHeaderOptions contains the optional parameters for the ExplicitClient.PostOptionalIntegerHeader
// method.
type ExplicitClientPostOptionalIntegerHeaderOptions struct {
	HeaderParameter *int32
}

// ExplicitClientPostOptionalIntegerParameterOptions contains the optional parameters for the ExplicitClient.PostOptionalIntegerParameter
// method.
type ExplicitClientPostOptionalIntegerParameterOptions struct {
	BodyParameter *int32
}

// ExplicitClientPostOptionalIntegerPropertyOptions contains the optional parameters for the ExplicitClient.PostOptionalIntegerProperty
// method.
type ExplicitClientPostOptionalIntegerPropertyOptions struct {
	BodyParameter *IntOptionalWrapper
}

// ExplicitClientPostOptionalStringHeaderOptions contains the optional parameters for the ExplicitClient.PostOptionalStringHeader
// method.
type ExplicitClientPostOptionalStringHeaderOptions struct {
	BodyParameter *string
}

// ExplicitClientPostOptionalStringParameterOptions contains the optional parameters for the ExplicitClient.PostOptionalStringParameter
// method.
type ExplicitClientPostOptionalStringParameterOptions struct {
	BodyParameter *string
}

// ExplicitClientPostOptionalStringPropertyOptions contains the optional parameters for the ExplicitClient.PostOptionalStringProperty
// method.
type ExplicitClientPostOptionalStringPropertyOptions struct {
	BodyParameter *StringOptionalWrapper
}

// ExplicitClientPostRequiredArrayHeaderOptions contains the optional parameters for the ExplicitClient.PostRequiredArrayHeader
// method.
type ExplicitClientPostRequiredArrayHeaderOptions struct {
	// placeholder for future optional parameters
}

// ExplicitClientPostRequiredArrayParameterOptions contains the optional parameters for the ExplicitClient.PostRequiredArrayParameter
// method.
type ExplicitClientPostRequiredArrayParameterOptions struct {
	// placeholder for future optional parameters
}

// ExplicitClientPostRequiredArrayPropertyOptions contains the optional parameters for the ExplicitClient.PostRequiredArrayProperty
// method.
type ExplicitClientPostRequiredArrayPropertyOptions struct {
	// placeholder for future optional parameters
}

// ExplicitClientPostRequiredClassParameterOptions contains the optional parameters for the ExplicitClient.PostRequiredClassParameter
// method.
type ExplicitClientPostRequiredClassParameterOptions struct {
	// placeholder for future optional parameters
}

// ExplicitClientPostRequiredClassPropertyOptions contains the optional parameters for the ExplicitClient.PostRequiredClassProperty
// method.
type ExplicitClientPostRequiredClassPropertyOptions struct {
	// placeholder for future optional parameters
}

// ExplicitClientPostRequiredIntegerHeaderOptions contains the optional parameters for the ExplicitClient.PostRequiredIntegerHeader
// method.
type ExplicitClientPostRequiredIntegerHeaderOptions struct {
	// placeholder for future optional parameters
}

// ExplicitClientPostRequiredIntegerParameterOptions contains the optional parameters for the ExplicitClient.PostRequiredIntegerParameter
// method.
type ExplicitClientPostRequiredIntegerParameterOptions struct {
	// placeholder for future optional parameters
}

// ExplicitClientPostRequiredIntegerPropertyOptions contains the optional parameters for the ExplicitClient.PostRequiredIntegerProperty
// method.
type ExplicitClientPostRequiredIntegerPropertyOptions struct {
	// placeholder for future optional parameters
}

// ExplicitClientPostRequiredStringHeaderOptions contains the optional parameters for the ExplicitClient.PostRequiredStringHeader
// method.
type ExplicitClientPostRequiredStringHeaderOptions struct {
	// placeholder for future optional parameters
}

// ExplicitClientPostRequiredStringParameterOptions contains the optional parameters for the ExplicitClient.PostRequiredStringParameter
// method.
type ExplicitClientPostRequiredStringParameterOptions struct {
	// placeholder for future optional parameters
}

// ExplicitClientPostRequiredStringPropertyOptions contains the optional parameters for the ExplicitClient.PostRequiredStringProperty
// method.
type ExplicitClientPostRequiredStringPropertyOptions struct {
	// placeholder for future optional parameters
}

// ExplicitClientPutOptionalBinaryBodyOptions contains the optional parameters for the ExplicitClient.PutOptionalBinaryBody
// method.
type ExplicitClientPutOptionalBinaryBodyOptions struct {
	BodyParameter io.ReadSeekCloser
}

// ExplicitClientPutRequiredBinaryBodyOptions contains the optional parameters for the ExplicitClient.PutRequiredBinaryBody
// method.
type ExplicitClientPutRequiredBinaryBodyOptions struct {
	// placeholder for future optional parameters
}

// ImplicitClientGetOptionalGlobalQueryOptions contains the optional parameters for the ImplicitClient.GetOptionalGlobalQuery
// method.
type ImplicitClientGetOptionalGlobalQueryOptions struct {
	// placeholder for future optional parameters
}

// ImplicitClientGetRequiredGlobalPathOptions contains the optional parameters for the ImplicitClient.GetRequiredGlobalPath
// method.
type ImplicitClientGetRequiredGlobalPathOptions struct {
	// placeholder for future optional parameters
}

// ImplicitClientGetRequiredGlobalQueryOptions contains the optional parameters for the ImplicitClient.GetRequiredGlobalQuery
// method.
type ImplicitClientGetRequiredGlobalQueryOptions struct {
	// placeholder for future optional parameters
}

// ImplicitClientGetRequiredPathOptions contains the optional parameters for the ImplicitClient.GetRequiredPath method.
type ImplicitClientGetRequiredPathOptions struct {
	// placeholder for future optional parameters
}

// ImplicitClientPutOptionalBinaryBodyOptions contains the optional parameters for the ImplicitClient.PutOptionalBinaryBody
// method.
type ImplicitClientPutOptionalBinaryBodyOptions struct {
	BodyParameter io.ReadSeekCloser
}

// ImplicitClientPutOptionalBodyOptions contains the optional parameters for the ImplicitClient.PutOptionalBody method.
type ImplicitClientPutOptionalBodyOptions struct {
	BodyParameter *string
}

// ImplicitClientPutOptionalHeaderOptions contains the optional parameters for the ImplicitClient.PutOptionalHeader method.
type ImplicitClientPutOptionalHeaderOptions struct {
	QueryParameter *string
}

// ImplicitClientPutOptionalQueryOptions contains the optional parameters for the ImplicitClient.PutOptionalQuery method.
type ImplicitClientPutOptionalQueryOptions struct {
	QueryParameter *string
}
