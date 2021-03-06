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

// RouteInsertRequest route insert request
//
// swagger:model routeInsertRequest
type RouteInsertRequest struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// method
	// Required: true
	Method *string `json:"method"`

	// namespace Id
	// Required: true
	NamespaceID *int64 `json:"namespaceId"`

	// route
	// Required: true
	Route *string `json:"route"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this route insert request
func (m *RouteInsertRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoute(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RouteInsertRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *RouteInsertRequest) validateMethod(formats strfmt.Registry) error {

	if err := validate.Required("method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

func (m *RouteInsertRequest) validateNamespaceID(formats strfmt.Registry) error {

	if err := validate.Required("namespaceId", "body", m.NamespaceID); err != nil {
		return err
	}

	return nil
}

func (m *RouteInsertRequest) validateRoute(formats strfmt.Registry) error {

	if err := validate.Required("route", "body", m.Route); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this route insert request based on context it is used
func (m *RouteInsertRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RouteInsertRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RouteInsertRequest) UnmarshalBinary(b []byte) error {
	var res RouteInsertRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
