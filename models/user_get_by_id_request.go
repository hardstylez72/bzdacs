// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserGetByIDRequest user get by Id request
//
// swagger:model userGetByIdRequest
type UserGetByIDRequest struct {

	// id
	// Required: true
	ID *int64 `json:"id"`

	// namespace Id
	// Required: true
	NamespaceID *int64 `json:"namespaceId"`
}

// Validate validates this user get by Id request
func (m *UserGetByIDRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserGetByIDRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *UserGetByIDRequest) validateNamespaceID(formats strfmt.Registry) error {

	if err := validate.Required("namespaceId", "body", m.NamespaceID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user get by Id request based on context it is used
func (m *UserGetByIDRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserGetByIDRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserGetByIDRequest) UnmarshalBinary(b []byte) error {
	var res UserGetByIDRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
