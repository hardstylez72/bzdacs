// Code generated by go-swagger; DO NOT EDIT.

package sys_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hardstylez72/bzdacs/generated/models"
)

// SysUserSessionReader is a Reader for the SysUserSession structure.
type SysUserSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SysUserSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSysUserSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSysUserSessionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSysUserSessionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSysUserSessionOK creates a SysUserSessionOK with default headers values
func NewSysUserSessionOK() *SysUserSessionOK {
	return &SysUserSessionOK{}
}

/* SysUserSessionOK describes a response with status code 200, with default header values.

OK
*/
type SysUserSessionOK struct {
	Payload *models.UserSessionResponse
}

func (o *SysUserSessionOK) Error() string {
	return fmt.Sprintf("[POST /v1/sys-user/session][%d] sysUserSessionOK  %+v", 200, o.Payload)
}
func (o *SysUserSessionOK) GetPayload() *models.UserSessionResponse {
	return o.Payload
}

func (o *SysUserSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserSessionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSysUserSessionUnauthorized creates a SysUserSessionUnauthorized with default headers values
func NewSysUserSessionUnauthorized() *SysUserSessionUnauthorized {
	return &SysUserSessionUnauthorized{}
}

/* SysUserSessionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SysUserSessionUnauthorized struct {
	Payload *models.UtilResponseWithError
}

func (o *SysUserSessionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/sys-user/session][%d] sysUserSessionUnauthorized  %+v", 401, o.Payload)
}
func (o *SysUserSessionUnauthorized) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *SysUserSessionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSysUserSessionInternalServerError creates a SysUserSessionInternalServerError with default headers values
func NewSysUserSessionInternalServerError() *SysUserSessionInternalServerError {
	return &SysUserSessionInternalServerError{}
}

/* SysUserSessionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type SysUserSessionInternalServerError struct {
	Payload *models.UtilResponseWithError
}

func (o *SysUserSessionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/sys-user/session][%d] sysUserSessionInternalServerError  %+v", 500, o.Payload)
}
func (o *SysUserSessionInternalServerError) GetPayload() *models.UtilResponseWithError {
	return o.Payload
}

func (o *SysUserSessionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UtilResponseWithError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
