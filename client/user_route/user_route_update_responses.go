// Code generated by go-swagger; DO NOT EDIT.

package user_route

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hardstylez72/bzdacs/models"
)

// UserRouteUpdateReader is a Reader for the UserRouteUpdate structure.
type UserRouteUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserRouteUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserRouteUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserRouteUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserRouteUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserRouteUpdateOK creates a UserRouteUpdateOK with default headers values
func NewUserRouteUpdateOK() *UserRouteUpdateOK {
	return &UserRouteUpdateOK{}
}

/* UserRouteUpdateOK describes a response with status code 200, with default header values.

OK
*/
type UserRouteUpdateOK struct {
	Payload *models.UserRouteUpdateResponse
}

func (o *UserRouteUpdateOK) Error() string {
	return fmt.Sprintf("[POST /v1/user/route/update][%d] userRouteUpdateOK  %+v", 200, o.Payload)
}
func (o *UserRouteUpdateOK) GetPayload() *models.UserRouteUpdateResponse {
	return o.Payload
}

func (o *UserRouteUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserRouteUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserRouteUpdateBadRequest creates a UserRouteUpdateBadRequest with default headers values
func NewUserRouteUpdateBadRequest() *UserRouteUpdateBadRequest {
	return &UserRouteUpdateBadRequest{}
}

/* UserRouteUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserRouteUpdateBadRequest struct {
	Payload *models.UtilResponseWithError
}

func (o *UserRouteUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/user/route/update][%d] userRouteUpdateBadRequest  %+v", 400, o.Payload)
}
func (o *UserRouteUpdateBadRequest) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *UserRouteUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserRouteUpdateInternalServerError creates a UserRouteUpdateInternalServerError with default headers values
func NewUserRouteUpdateInternalServerError() *UserRouteUpdateInternalServerError {
	return &UserRouteUpdateInternalServerError{}
}

/* UserRouteUpdateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UserRouteUpdateInternalServerError struct {
	Payload *models.UtilResponseWithError
}

func (o *UserRouteUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/user/route/update][%d] userRouteUpdateInternalServerError  %+v", 500, o.Payload)
}
func (o *UserRouteUpdateInternalServerError) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *UserRouteUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}