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

// SystemUpdateReader is a Reader for the SystemUpdate structure.
type SystemUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSystemUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSystemUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemUpdateOK creates a SystemUpdateOK with default headers values
func NewSystemUpdateOK() *SystemUpdateOK {
	return &SystemUpdateOK{}
}

/* SystemUpdateOK describes a response with status code 200, with default header values.

OK
*/
type SystemUpdateOK struct {
	Payload *models.SystemGetResponse
}

func (o *SystemUpdateOK) Error() string {
	return fmt.Sprintf("[POST /v1/system/update][%d] systemUpdateOK  %+v", 200, o.Payload)
}
func (o *SystemUpdateOK) GetPayload() *models.SystemGetResponse {
	return o.Payload
}

func (o *SystemUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemUpdateBadRequest creates a SystemUpdateBadRequest with default headers values
func NewSystemUpdateBadRequest() *SystemUpdateBadRequest {
	return &SystemUpdateBadRequest{}
}

/* SystemUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SystemUpdateBadRequest struct {
	Payload *models.UtilResponseWithError
}

func (o *SystemUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/system/update][%d] systemUpdateBadRequest  %+v", 400, o.Payload)
}
func (o *SystemUpdateBadRequest) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *SystemUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemUpdateInternalServerError creates a SystemUpdateInternalServerError with default headers values
func NewSystemUpdateInternalServerError() *SystemUpdateInternalServerError {
	return &SystemUpdateInternalServerError{}
}

/* SystemUpdateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type SystemUpdateInternalServerError struct {
	Payload *models.UtilResponseWithError
}

func (o *SystemUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/system/update][%d] systemUpdateInternalServerError  %+v", 500, o.Payload)
}
func (o *SystemUpdateInternalServerError) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *SystemUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
