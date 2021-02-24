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

// GrouprouteListRequest grouproute list request
//
// swagger:model grouproute.listRequest
type GrouprouteListRequest struct {

	// belong to group
	BelongToGroup bool `json:"belongToGroup,omitempty"`

	// group Id
	// Required: true
	GroupID *int64 `json:"groupId"`
}

// Validate validates this grouproute list request
func (m *GrouprouteListRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GrouprouteListRequest) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("groupId", "body", m.GroupID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this grouproute list request based on context it is used
func (m *GrouprouteListRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GrouprouteListRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GrouprouteListRequest) UnmarshalBinary(b []byte) error {
	var res GrouprouteListRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
