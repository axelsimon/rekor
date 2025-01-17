// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package timestamp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sigstore/rekor/pkg/generated/models"
)

// GetTimestampResponseReader is a Reader for the GetTimestampResponse structure.
type GetTimestampResponseReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetTimestampResponseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTimestampResponseOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTimestampResponseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewGetTimestampResponseNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetTimestampResponseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTimestampResponseOK creates a GetTimestampResponseOK with default headers values
func NewGetTimestampResponseOK(writer io.Writer) *GetTimestampResponseOK {
	return &GetTimestampResponseOK{

		Payload: writer,
	}
}

/* GetTimestampResponseOK describes a response with status code 200, with default header values.

Returns a timestamp response
*/
type GetTimestampResponseOK struct {
	Payload io.Writer
}

func (o *GetTimestampResponseOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/timestamp][%d] getTimestampResponseOK  %+v", 200, o.Payload)
}
func (o *GetTimestampResponseOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetTimestampResponseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTimestampResponseBadRequest creates a GetTimestampResponseBadRequest with default headers values
func NewGetTimestampResponseBadRequest() *GetTimestampResponseBadRequest {
	return &GetTimestampResponseBadRequest{}
}

/* GetTimestampResponseBadRequest describes a response with status code 400, with default header values.

The content supplied to the server was invalid
*/
type GetTimestampResponseBadRequest struct {
	Payload *models.Error
}

func (o *GetTimestampResponseBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/timestamp][%d] getTimestampResponseBadRequest  %+v", 400, o.Payload)
}
func (o *GetTimestampResponseBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTimestampResponseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTimestampResponseNotImplemented creates a GetTimestampResponseNotImplemented with default headers values
func NewGetTimestampResponseNotImplemented() *GetTimestampResponseNotImplemented {
	return &GetTimestampResponseNotImplemented{}
}

/* GetTimestampResponseNotImplemented describes a response with status code 501, with default header values.

The content requested is not implemented
*/
type GetTimestampResponseNotImplemented struct {
}

func (o *GetTimestampResponseNotImplemented) Error() string {
	return fmt.Sprintf("[POST /api/v1/timestamp][%d] getTimestampResponseNotImplemented ", 501)
}

func (o *GetTimestampResponseNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTimestampResponseDefault creates a GetTimestampResponseDefault with default headers values
func NewGetTimestampResponseDefault(code int) *GetTimestampResponseDefault {
	return &GetTimestampResponseDefault{
		_statusCode: code,
	}
}

/* GetTimestampResponseDefault describes a response with status code -1, with default header values.

There was an internal error in the server while processing the request
*/
type GetTimestampResponseDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get timestamp response default response
func (o *GetTimestampResponseDefault) Code() int {
	return o._statusCode
}

func (o *GetTimestampResponseDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/timestamp][%d] getTimestampResponse default  %+v", o._statusCode, o.Payload)
}
func (o *GetTimestampResponseDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTimestampResponseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
