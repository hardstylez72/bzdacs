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

// GrouprouteFilter grouproute filter
//
// swagger:model grouproute.Filter
type GrouprouteFilter struct {

	// belong to group
	BelongToGroup bool `json:"belongToGroup,omitempty"`

	// group Id
	// Required: true
	GroupID *int64 `json:"groupId"`

	// namespace Id
	// Required: true
	NamespaceID *int64 `json:"namespaceId"`

	// page
	// Required: true
	Page *int64 `json:"page"`

	// page size
	PageSize int64 `json:"pageSize,omitempty"`

	// route pattern
	RoutePattern string `json:"routePattern,omitempty"`
}

// Validate validates this grouproute filter
func (m *GrouprouteFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GrouprouteFilter) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("groupId", "body", m.GroupID); err != nil {
		return err
	}

	return nil
}

func (m *GrouprouteFilter) validateNamespaceID(formats strfmt.Registry) error {

	if err := validate.Required("namespaceId", "body", m.NamespaceID); err != nil {
		return err
	}

	return nil
}

func (m *GrouprouteFilter) validatePage(formats strfmt.Registry) error {

	if err := validate.Required("page", "body", m.Page); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this grouproute filter based on context it is used
func (m *GrouprouteFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GrouprouteFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GrouprouteFilter) UnmarshalBinary(b []byte) error {
	var res GrouprouteFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
