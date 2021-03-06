// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hardstylez72/bzdacs/generated/models"
)

// NamespaceCreateReader is a Reader for the NamespaceCreate structure.
type NamespaceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NamespaceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNamespaceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNamespaceCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNamespaceCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNamespaceCreateOK creates a NamespaceCreateOK with default headers values
func NewNamespaceCreateOK() *NamespaceCreateOK {
	return &NamespaceCreateOK{}
}

/* NamespaceCreateOK describes a response with status code 200, with default header values.

OK
*/
type NamespaceCreateOK struct {
	Payload *models.NamespaceGetResponse
}

func (o *NamespaceCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1/namespace/create][%d] namespaceCreateOK  %+v", 200, o.Payload)
}
func (o *NamespaceCreateOK) GetPayload() *models.NamespaceGetResponse {
	return o.Payload
}

func (o *NamespaceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NamespaceGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamespaceCreateBadRequest creates a NamespaceCreateBadRequest with default headers values
func NewNamespaceCreateBadRequest() *NamespaceCreateBadRequest {
	return &NamespaceCreateBadRequest{}
}

/* NamespaceCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type NamespaceCreateBadRequest struct {
	Payload *models.UtilResponseWithError
}

func (o *NamespaceCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/namespace/create][%d] namespaceCreateBadRequest  %+v", 400, o.Payload)
}
func (o *NamespaceCreateBadRequest) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *NamespaceCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamespaceCreateInternalServerError creates a NamespaceCreateInternalServerError with default headers values
func NewNamespaceCreateInternalServerError() *NamespaceCreateInternalServerError {
	return &NamespaceCreateInternalServerError{}
}

/* NamespaceCreateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type NamespaceCreateInternalServerError struct {
	Payload *models.UtilResponseWithError
}

func (o *NamespaceCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/namespace/create][%d] namespaceCreateInternalServerError  %+v", 500, o.Payload)
}
func (o *NamespaceCreateInternalServerError) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *NamespaceCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
