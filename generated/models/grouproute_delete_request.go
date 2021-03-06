// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GrouprouteDeleteRequest grouproute delete request
//
// swagger:model grouproute.deleteRequest
type GrouprouteDeleteRequest struct {

	// pairs
	Pairs []*GrouproutePair `json:"pairs"`
}

// Validate validates this grouproute delete request
func (m *GrouprouteDeleteRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePairs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GrouprouteDeleteRequest) validatePairs(formats strfmt.Registry) error {
	if swag.IsZero(m.Pairs) { // not required
		return nil
	}

	for i := 0; i < len(m.Pairs); i++ {
		if swag.IsZero(m.Pairs[i]) { // not required
			continue
		}

		if m.Pairs[i] != nil {
			if err := m.Pairs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pairs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this grouproute delete request based on the context it is used
func (m *GrouprouteDeleteRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePairs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GrouprouteDeleteRequest) contextValidatePairs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Pairs); i++ {

		if m.Pairs[i] != nil {
			if err := m.Pairs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pairs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GrouprouteDeleteRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GrouprouteDeleteRequest) UnmarshalBinary(b []byte) error {
	var res GrouprouteDeleteRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
