// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hardstylez72/bzdacs/models"
)

// GroupGetByIDReader is a Reader for the GroupGetByID structure.
type GroupGetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupGetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupGetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGroupGetByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGroupGetByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGroupGetByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGroupGetByIDOK creates a GroupGetByIDOK with default headers values
func NewGroupGetByIDOK() *GroupGetByIDOK {
	return &GroupGetByIDOK{}
}

/* GroupGetByIDOK describes a response with status code 200, with default header values.

OK
*/
type GroupGetByIDOK struct {
	Payload *models.GroupGetResponse
}

func (o *GroupGetByIDOK) Error() string {
	return fmt.Sprintf("[POST /v1/group/getById][%d] groupGetByIdOK  %+v", 200, o.Payload)
}
func (o *GroupGetByIDOK) GetPayload() *models.GroupGetResponse {
	return o.Payload
}

func (o *GroupGetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupGetByIDBadRequest creates a GroupGetByIDBadRequest with default headers values
func NewGroupGetByIDBadRequest() *GroupGetByIDBadRequest {
	return &GroupGetByIDBadRequest{}
}

/* GroupGetByIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GroupGetByIDBadRequest struct {
	Payload *models.UtilResponseWithError
}

func (o *GroupGetByIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/group/getById][%d] groupGetByIdBadRequest  %+v", 400, o.Payload)
}
func (o *GroupGetByIDBadRequest) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *GroupGetByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupGetByIDNotFound creates a GroupGetByIDNotFound with default headers values
func NewGroupGetByIDNotFound() *GroupGetByIDNotFound {
	return &GroupGetByIDNotFound{}
}

/* GroupGetByIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GroupGetByIDNotFound struct {
	Payload *models.UtilResponseWithError
}

func (o *GroupGetByIDNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/group/getById][%d] groupGetByIdNotFound  %+v", 404, o.Payload)
}
func (o *GroupGetByIDNotFound) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *GroupGetByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupGetByIDInternalServerError creates a GroupGetByIDInternalServerError with default headers values
func NewGroupGetByIDInternalServerError() *GroupGetByIDInternalServerError {
	return &GroupGetByIDInternalServerError{}
}

/* GroupGetByIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GroupGetByIDInternalServerError struct {
	Payload *models.UtilResponseWithError
}

func (o *GroupGetByIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/group/getById][%d] groupGetByIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GroupGetByIDInternalServerError) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *GroupGetByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
