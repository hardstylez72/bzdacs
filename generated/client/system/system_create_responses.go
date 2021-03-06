// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hardstylez72/bzdacs/generated/models"
)

// SystemCreateReader is a Reader for the SystemCreate structure.
type SystemCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSystemCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSystemCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemCreateOK creates a SystemCreateOK with default headers values
func NewSystemCreateOK() *SystemCreateOK {
	return &SystemCreateOK{}
}

/* SystemCreateOK describes a response with status code 200, with default header values.

OK
*/
type SystemCreateOK struct {
	Payload *models.SystemGetResponse
}

func (o *SystemCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1/system/create][%d] systemCreateOK  %+v", 200, o.Payload)
}
func (o *SystemCreateOK) GetPayload() *models.SystemGetResponse {
	return o.Payload
}

func (o *SystemCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemCreateBadRequest creates a SystemCreateBadRequest with default headers values
func NewSystemCreateBadRequest() *SystemCreateBadRequest {
	return &SystemCreateBadRequest{}
}

/* SystemCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SystemCreateBadRequest struct {
	Payload *models.UtilResponseWithError
}

func (o *SystemCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/system/create][%d] systemCreateBadRequest  %+v", 400, o.Payload)
}
func (o *SystemCreateBadRequest) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *SystemCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemCreateInternalServerError creates a SystemCreateInternalServerError with default headers values
func NewSystemCreateInternalServerError() *SystemCreateInternalServerError {
	return &SystemCreateInternalServerError{}
}

/* SystemCreateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type SystemCreateInternalServerError struct {
	Payload *models.UtilResponseWithError
}

func (o *SystemCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/system/create][%d] systemCreateInternalServerError  %+v", 500, o.Payload)
}
func (o *SystemCreateInternalServerError) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *SystemCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
