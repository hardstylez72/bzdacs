// Code generated by go-swagger; DO NOT EDIT.

package relations_group_route

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hardstylez72/bzdacs/models"
)

// GroupRoutesDeleteReader is a Reader for the GroupRoutesDelete structure.
type GroupRoutesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupRoutesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupRoutesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGroupRoutesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGroupRoutesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGroupRoutesDeleteOK creates a GroupRoutesDeleteOK with default headers values
func NewGroupRoutesDeleteOK() *GroupRoutesDeleteOK {
	return &GroupRoutesDeleteOK{}
}

/* GroupRoutesDeleteOK describes a response with status code 200, with default header values.

GroupRoutesDeleteOK group routes delete o k
*/
type GroupRoutesDeleteOK struct {
}

func (o *GroupRoutesDeleteOK) Error() string {
	return fmt.Sprintf("[POST /v1/group/route/delete][%d] groupRoutesDeleteOK ", 200)
}

func (o *GroupRoutesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGroupRoutesDeleteBadRequest creates a GroupRoutesDeleteBadRequest with default headers values
func NewGroupRoutesDeleteBadRequest() *GroupRoutesDeleteBadRequest {
	return &GroupRoutesDeleteBadRequest{}
}

/* GroupRoutesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GroupRoutesDeleteBadRequest struct {
	Payload *models.UtilResponseWithError
}

func (o *GroupRoutesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/group/route/delete][%d] groupRoutesDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *GroupRoutesDeleteBadRequest) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *GroupRoutesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupRoutesDeleteInternalServerError creates a GroupRoutesDeleteInternalServerError with default headers values
func NewGroupRoutesDeleteInternalServerError() *GroupRoutesDeleteInternalServerError {
	return &GroupRoutesDeleteInternalServerError{}
}

/* GroupRoutesDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GroupRoutesDeleteInternalServerError struct {
	Payload *models.UtilResponseWithError
}

func (o *GroupRoutesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/group/route/delete][%d] groupRoutesDeleteInternalServerError  %+v", 500, o.Payload)
}
func (o *GroupRoutesDeleteInternalServerError) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *GroupRoutesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}