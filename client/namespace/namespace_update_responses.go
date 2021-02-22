// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hardstylez72/bzdacs/models"
)

// NamespaceUpdateReader is a Reader for the NamespaceUpdate structure.
type NamespaceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NamespaceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNamespaceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNamespaceUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNamespaceUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNamespaceUpdateOK creates a NamespaceUpdateOK with default headers values
func NewNamespaceUpdateOK() *NamespaceUpdateOK {
	return &NamespaceUpdateOK{}
}

/* NamespaceUpdateOK describes a response with status code 200, with default header values.

OK
*/
type NamespaceUpdateOK struct {
	Payload *models.NamespaceGetResponse
}

func (o *NamespaceUpdateOK) Error() string {
	return fmt.Sprintf("[POST /v1/namespace/update][%d] namespaceUpdateOK  %+v", 200, o.Payload)
}
func (o *NamespaceUpdateOK) GetPayload() *models.NamespaceGetResponse {
	return o.Payload
}

func (o *NamespaceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NamespaceGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamespaceUpdateBadRequest creates a NamespaceUpdateBadRequest with default headers values
func NewNamespaceUpdateBadRequest() *NamespaceUpdateBadRequest {
	return &NamespaceUpdateBadRequest{}
}

/* NamespaceUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type NamespaceUpdateBadRequest struct {
	Payload *models.UtilResponseWithError
}

func (o *NamespaceUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/namespace/update][%d] namespaceUpdateBadRequest  %+v", 400, o.Payload)
}
func (o *NamespaceUpdateBadRequest) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *NamespaceUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamespaceUpdateInternalServerError creates a NamespaceUpdateInternalServerError with default headers values
func NewNamespaceUpdateInternalServerError() *NamespaceUpdateInternalServerError {
	return &NamespaceUpdateInternalServerError{}
}

/* NamespaceUpdateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type NamespaceUpdateInternalServerError struct {
	Payload *models.UtilResponseWithError
}

func (o *NamespaceUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/namespace/update][%d] namespaceUpdateInternalServerError  %+v", 500, o.Payload)
}
func (o *NamespaceUpdateInternalServerError) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *NamespaceUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
