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

// GroupRoutePair group route pair
//
// swagger:model groupRoutePair
type GroupRoutePair struct {

	// group Id
	// Required: true
	GroupID *int64 `json:"groupId"`

	// route Id
	// Required: true
	RouteID *int64 `json:"routeId"`
}

// Validate validates this group route pair
func (m *GroupRoutePair) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupRoutePair) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("groupId", "body", m.GroupID); err != nil {
		return err
	}

	return nil
}

func (m *GroupRoutePair) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("routeId", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this group route pair based on context it is used
func (m *GroupRoutePair) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GroupRoutePair) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupRoutePair) UnmarshalBinary(b []byte) error {
	var res GroupRoutePair
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
