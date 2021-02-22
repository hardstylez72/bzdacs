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

// RouteDeleteRequest route delete request
//
// swagger:model route.deleteRequest
type RouteDeleteRequest struct {

	// id
	// Required: true
	ID *int64 `json:"id"`

	// namespace Id
	// Required: true
	NamespaceID *int64 `json:"namespaceId"`
}

// Validate validates this route delete request
func (m *RouteDeleteRequest) Validate(formats strfmt.Registry) error {
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

func (m *RouteDeleteRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RouteDeleteRequest) validateNamespaceID(formats strfmt.Registry) error {

	if err := validate.Required("namespaceId", "body", m.NamespaceID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this route delete request based on context it is used
func (m *RouteDeleteRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RouteDeleteRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RouteDeleteRequest) UnmarshalBinary(b []byte) error {
	var res RouteDeleteRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
