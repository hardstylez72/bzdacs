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

// UserroutePair userroute pair
//
// swagger:model userroute.Pair
type UserroutePair struct {

	// is excluded
	// Required: true
	IsExcluded *bool `json:"isExcluded"`

	// route Id
	// Required: true
	RouteID *int64 `json:"routeId"`

	// user Id
	// Required: true
	UserID *int64 `json:"userId"`
}

// Validate validates this userroute pair
func (m *UserroutePair) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsExcluded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserroutePair) validateIsExcluded(formats strfmt.Registry) error {

	if err := validate.Required("isExcluded", "body", m.IsExcluded); err != nil {
		return err
	}

	return nil
}

func (m *UserroutePair) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("routeId", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *UserroutePair) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this userroute pair based on context it is used
func (m *UserroutePair) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserroutePair) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserroutePair) UnmarshalBinary(b []byte) error {
	var res UserroutePair
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
