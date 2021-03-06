// Code generated by go-swagger; DO NOT EDIT.

package route

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hardstylez72/bzdacs/generated/models"
)

// RouteGetByIDReader is a Reader for the RouteGetByID structure.
type RouteGetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteGetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouteGetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRouteGetByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRouteGetByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRouteGetByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRouteGetByIDOK creates a RouteGetByIDOK with default headers values
func NewRouteGetByIDOK() *RouteGetByIDOK {
	return &RouteGetByIDOK{}
}

/* RouteGetByIDOK describes a response with status code 200, with default header values.

OK
*/
type RouteGetByIDOK struct {
	Payload *models.RouteGetResponse
}

func (o *RouteGetByIDOK) Error() string {
	return fmt.Sprintf("[POST /v1/route/getById][%d] routeGetByIdOK  %+v", 200, o.Payload)
}
func (o *RouteGetByIDOK) GetPayload() *models.RouteGetResponse {
	return o.Payload
}

func (o *RouteGetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouteGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouteGetByIDBadRequest creates a RouteGetByIDBadRequest with default headers values
func NewRouteGetByIDBadRequest() *RouteGetByIDBadRequest {
	return &RouteGetByIDBadRequest{}
}

/* RouteGetByIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RouteGetByIDBadRequest struct {
	Payload *models.UtilResponseWithError
}

func (o *RouteGetByIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/route/getById][%d] routeGetByIdBadRequest  %+v", 400, o.Payload)
}
func (o *RouteGetByIDBadRequest) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *RouteGetByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouteGetByIDNotFound creates a RouteGetByIDNotFound with default headers values
func NewRouteGetByIDNotFound() *RouteGetByIDNotFound {
	return &RouteGetByIDNotFound{}
}

/* RouteGetByIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RouteGetByIDNotFound struct {
	Payload *models.UtilResponseWithError
}

func (o *RouteGetByIDNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/route/getById][%d] routeGetByIdNotFound  %+v", 404, o.Payload)
}
func (o *RouteGetByIDNotFound) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *RouteGetByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouteGetByIDInternalServerError creates a RouteGetByIDInternalServerError with default headers values
func NewRouteGetByIDInternalServerError() *RouteGetByIDInternalServerError {
	return &RouteGetByIDInternalServerError{}
}

/* RouteGetByIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RouteGetByIDInternalServerError struct {
	Payload *models.UtilResponseWithError
}

func (o *RouteGetByIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/route/getById][%d] routeGetByIdInternalServerError  %+v", 500, o.Payload)
}
func (o *RouteGetByIDInternalServerError) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *RouteGetByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
