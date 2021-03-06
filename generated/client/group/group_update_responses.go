// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hardstylez72/bzdacs/generated/models"
)

// GroupUpdateReader is a Reader for the GroupUpdate structure.
type GroupUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGroupUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGroupUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGroupUpdateOK creates a GroupUpdateOK with default headers values
func NewGroupUpdateOK() *GroupUpdateOK {
	return &GroupUpdateOK{}
}

/* GroupUpdateOK describes a response with status code 200, with default header values.

OK
*/
type GroupUpdateOK struct {
	Payload *models.GroupGetResponse
}

func (o *GroupUpdateOK) Error() string {
	return fmt.Sprintf("[POST /v1/group/update][%d] groupUpdateOK  %+v", 200, o.Payload)
}
func (o *GroupUpdateOK) GetPayload() *models.GroupGetResponse {
	return o.Payload
}

func (o *GroupUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupUpdateBadRequest creates a GroupUpdateBadRequest with default headers values
func NewGroupUpdateBadRequest() *GroupUpdateBadRequest {
	return &GroupUpdateBadRequest{}
}

/* GroupUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GroupUpdateBadRequest struct {
	Payload *models.UtilResponseWithError
}

func (o *GroupUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/group/update][%d] groupUpdateBadRequest  %+v", 400, o.Payload)
}
func (o *GroupUpdateBadRequest) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *GroupUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupUpdateInternalServerError creates a GroupUpdateInternalServerError with default headers values
func NewGroupUpdateInternalServerError() *GroupUpdateInternalServerError {
	return &GroupUpdateInternalServerError{}
}

/* GroupUpdateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GroupUpdateInternalServerError struct {
	Payload *models.UtilResponseWithError
}

func (o *GroupUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/group/update][%d] groupUpdateInternalServerError  %+v", 500, o.Payload)
}
func (o *GroupUpdateInternalServerError) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *GroupUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
