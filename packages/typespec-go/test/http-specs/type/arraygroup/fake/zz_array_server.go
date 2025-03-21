// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ArrayServer is a fake server for instances of the arraygroup.ArrayClient type.
type ArrayServer struct {
	// ArrayBooleanValueServer contains the fakes for client ArrayBooleanValueClient
	ArrayBooleanValueServer ArrayBooleanValueServer

	// ArrayDatetimeValueServer contains the fakes for client ArrayDatetimeValueClient
	ArrayDatetimeValueServer ArrayDatetimeValueServer

	// ArrayDurationValueServer contains the fakes for client ArrayDurationValueClient
	ArrayDurationValueServer ArrayDurationValueServer

	// ArrayFloat32ValueServer contains the fakes for client ArrayFloat32ValueClient
	ArrayFloat32ValueServer ArrayFloat32ValueServer

	// ArrayInt32ValueServer contains the fakes for client ArrayInt32ValueClient
	ArrayInt32ValueServer ArrayInt32ValueServer

	// ArrayInt64ValueServer contains the fakes for client ArrayInt64ValueClient
	ArrayInt64ValueServer ArrayInt64ValueServer

	// ArrayModelValueServer contains the fakes for client ArrayModelValueClient
	ArrayModelValueServer ArrayModelValueServer

	// ArrayNullableBooleanValueServer contains the fakes for client ArrayNullableBooleanValueClient
	ArrayNullableBooleanValueServer ArrayNullableBooleanValueServer

	// ArrayNullableFloatValueServer contains the fakes for client ArrayNullableFloatValueClient
	ArrayNullableFloatValueServer ArrayNullableFloatValueServer

	// ArrayNullableInt32ValueServer contains the fakes for client ArrayNullableInt32ValueClient
	ArrayNullableInt32ValueServer ArrayNullableInt32ValueServer

	// ArrayNullableModelValueServer contains the fakes for client ArrayNullableModelValueClient
	ArrayNullableModelValueServer ArrayNullableModelValueServer

	// ArrayNullableStringValueServer contains the fakes for client ArrayNullableStringValueClient
	ArrayNullableStringValueServer ArrayNullableStringValueServer

	// ArrayStringValueServer contains the fakes for client ArrayStringValueClient
	ArrayStringValueServer ArrayStringValueServer

	// ArrayUnknownValueServer contains the fakes for client ArrayUnknownValueClient
	ArrayUnknownValueServer ArrayUnknownValueServer
}

// NewArrayServerTransport creates a new instance of ArrayServerTransport with the provided implementation.
// The returned ArrayServerTransport instance is connected to an instance of arraygroup.ArrayClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewArrayServerTransport(srv *ArrayServer) *ArrayServerTransport {
	return &ArrayServerTransport{srv: srv}
}

// ArrayServerTransport connects instances of arraygroup.ArrayClient to instances of ArrayServer.
// Don't use this type directly, use NewArrayServerTransport instead.
type ArrayServerTransport struct {
	srv                               *ArrayServer
	trMu                              sync.Mutex
	trArrayBooleanValueServer         *ArrayBooleanValueServerTransport
	trArrayDatetimeValueServer        *ArrayDatetimeValueServerTransport
	trArrayDurationValueServer        *ArrayDurationValueServerTransport
	trArrayFloat32ValueServer         *ArrayFloat32ValueServerTransport
	trArrayInt32ValueServer           *ArrayInt32ValueServerTransport
	trArrayInt64ValueServer           *ArrayInt64ValueServerTransport
	trArrayModelValueServer           *ArrayModelValueServerTransport
	trArrayNullableBooleanValueServer *ArrayNullableBooleanValueServerTransport
	trArrayNullableFloatValueServer   *ArrayNullableFloatValueServerTransport
	trArrayNullableInt32ValueServer   *ArrayNullableInt32ValueServerTransport
	trArrayNullableModelValueServer   *ArrayNullableModelValueServerTransport
	trArrayNullableStringValueServer  *ArrayNullableStringValueServerTransport
	trArrayStringValueServer          *ArrayStringValueServerTransport
	trArrayUnknownValueServer         *ArrayUnknownValueServerTransport
}

// Do implements the policy.Transporter interface for ArrayServerTransport.
func (a *ArrayServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToClientFake(req, method[:strings.Index(method, ".")])
}

func (a *ArrayServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "ArrayBooleanValueClient":
		initServer(&a.trMu, &a.trArrayBooleanValueServer, func() *ArrayBooleanValueServerTransport {
			return NewArrayBooleanValueServerTransport(&a.srv.ArrayBooleanValueServer)
		})
		resp, err = a.trArrayBooleanValueServer.Do(req)
	case "ArrayDatetimeValueClient":
		initServer(&a.trMu, &a.trArrayDatetimeValueServer, func() *ArrayDatetimeValueServerTransport {
			return NewArrayDatetimeValueServerTransport(&a.srv.ArrayDatetimeValueServer)
		})
		resp, err = a.trArrayDatetimeValueServer.Do(req)
	case "ArrayDurationValueClient":
		initServer(&a.trMu, &a.trArrayDurationValueServer, func() *ArrayDurationValueServerTransport {
			return NewArrayDurationValueServerTransport(&a.srv.ArrayDurationValueServer)
		})
		resp, err = a.trArrayDurationValueServer.Do(req)
	case "ArrayFloat32ValueClient":
		initServer(&a.trMu, &a.trArrayFloat32ValueServer, func() *ArrayFloat32ValueServerTransport {
			return NewArrayFloat32ValueServerTransport(&a.srv.ArrayFloat32ValueServer)
		})
		resp, err = a.trArrayFloat32ValueServer.Do(req)
	case "ArrayInt32ValueClient":
		initServer(&a.trMu, &a.trArrayInt32ValueServer, func() *ArrayInt32ValueServerTransport {
			return NewArrayInt32ValueServerTransport(&a.srv.ArrayInt32ValueServer)
		})
		resp, err = a.trArrayInt32ValueServer.Do(req)
	case "ArrayInt64ValueClient":
		initServer(&a.trMu, &a.trArrayInt64ValueServer, func() *ArrayInt64ValueServerTransport {
			return NewArrayInt64ValueServerTransport(&a.srv.ArrayInt64ValueServer)
		})
		resp, err = a.trArrayInt64ValueServer.Do(req)
	case "ArrayModelValueClient":
		initServer(&a.trMu, &a.trArrayModelValueServer, func() *ArrayModelValueServerTransport {
			return NewArrayModelValueServerTransport(&a.srv.ArrayModelValueServer)
		})
		resp, err = a.trArrayModelValueServer.Do(req)
	case "ArrayNullableBooleanValueClient":
		initServer(&a.trMu, &a.trArrayNullableBooleanValueServer, func() *ArrayNullableBooleanValueServerTransport {
			return NewArrayNullableBooleanValueServerTransport(&a.srv.ArrayNullableBooleanValueServer)
		})
		resp, err = a.trArrayNullableBooleanValueServer.Do(req)
	case "ArrayNullableFloatValueClient":
		initServer(&a.trMu, &a.trArrayNullableFloatValueServer, func() *ArrayNullableFloatValueServerTransport {
			return NewArrayNullableFloatValueServerTransport(&a.srv.ArrayNullableFloatValueServer)
		})
		resp, err = a.trArrayNullableFloatValueServer.Do(req)
	case "ArrayNullableInt32ValueClient":
		initServer(&a.trMu, &a.trArrayNullableInt32ValueServer, func() *ArrayNullableInt32ValueServerTransport {
			return NewArrayNullableInt32ValueServerTransport(&a.srv.ArrayNullableInt32ValueServer)
		})
		resp, err = a.trArrayNullableInt32ValueServer.Do(req)
	case "ArrayNullableModelValueClient":
		initServer(&a.trMu, &a.trArrayNullableModelValueServer, func() *ArrayNullableModelValueServerTransport {
			return NewArrayNullableModelValueServerTransport(&a.srv.ArrayNullableModelValueServer)
		})
		resp, err = a.trArrayNullableModelValueServer.Do(req)
	case "ArrayNullableStringValueClient":
		initServer(&a.trMu, &a.trArrayNullableStringValueServer, func() *ArrayNullableStringValueServerTransport {
			return NewArrayNullableStringValueServerTransport(&a.srv.ArrayNullableStringValueServer)
		})
		resp, err = a.trArrayNullableStringValueServer.Do(req)
	case "ArrayStringValueClient":
		initServer(&a.trMu, &a.trArrayStringValueServer, func() *ArrayStringValueServerTransport {
			return NewArrayStringValueServerTransport(&a.srv.ArrayStringValueServer)
		})
		resp, err = a.trArrayStringValueServer.Do(req)
	case "ArrayUnknownValueClient":
		initServer(&a.trMu, &a.trArrayUnknownValueServer, func() *ArrayUnknownValueServerTransport {
			return NewArrayUnknownValueServerTransport(&a.srv.ArrayUnknownValueServer)
		})
		resp, err = a.trArrayUnknownValueServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}

// set this to conditionally intercept incoming requests to ArrayServerTransport
var arrayServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
