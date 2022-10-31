// Code generated by go-swagger; DO NOT EDIT.

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firecracker-microvm/firecracker-go-sdk/client/models"
)

// PutMmdsConfigReader is a Reader for the PutMmdsConfig structure.
type PutMmdsConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutMmdsConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutMmdsConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutMmdsConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutMmdsConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutMmdsConfigNoContent creates a PutMmdsConfigNoContent with default headers values
func NewPutMmdsConfigNoContent() *PutMmdsConfigNoContent {
	return &PutMmdsConfigNoContent{}
}

/*PutMmdsConfigNoContent handles this case with default header values.

MMDS configuration was created/updated.
*/
type PutMmdsConfigNoContent struct {
}

func (o *PutMmdsConfigNoContent) Error() string {
	return fmt.Sprintf("[PUT /mmds/config][%d] putMmdsConfigNoContent ", 204)
}

func (o *PutMmdsConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutMmdsConfigBadRequest creates a PutMmdsConfigBadRequest with default headers values
func NewPutMmdsConfigBadRequest() *PutMmdsConfigBadRequest {
	return &PutMmdsConfigBadRequest{}
}

/*PutMmdsConfigBadRequest handles this case with default header values.

MMDS configuration cannot be updated due to bad input.
*/
type PutMmdsConfigBadRequest struct {
	Payload *models.Error
}

func (o *PutMmdsConfigBadRequest) Error() string {
	return fmt.Sprintf("[PUT /mmds/config][%d] putMmdsConfigBadRequest  %+v", 400, o.Payload)
}

func (o *PutMmdsConfigBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutMmdsConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutMmdsConfigDefault creates a PutMmdsConfigDefault with default headers values
func NewPutMmdsConfigDefault(code int) *PutMmdsConfigDefault {
	return &PutMmdsConfigDefault{
		_statusCode: code,
	}
}

/*PutMmdsConfigDefault handles this case with default header values.

Internal server error
*/
type PutMmdsConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put mmds config default response
func (o *PutMmdsConfigDefault) Code() int {
	return o._statusCode
}

func (o *PutMmdsConfigDefault) Error() string {
	return fmt.Sprintf("[PUT /mmds/config][%d] putMmdsConfig default  %+v", o._statusCode, o.Payload)
}

func (o *PutMmdsConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutMmdsConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
