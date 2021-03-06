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

// UserRouteFilter user route filter
//
// swagger:model userRouteFilter
type UserRouteFilter struct {

	// belong to user
	BelongToUser bool `json:"belongToUser,omitempty"`

	// namespace Id
	// Required: true
	NamespaceID *int64 `json:"namespaceId"`

	// page
	// Required: true
	Page *int64 `json:"page"`

	// page size
	PageSize int64 `json:"pageSize,omitempty"`

	// pattern
	Pattern string `json:"pattern,omitempty"`

	// user Id
	// Required: true
	UserID *int64 `json:"userId"`
}

// Validate validates this user route filter
func (m *UserRouteFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNamespaceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePage(formats); err != nil {
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

func (m *UserRouteFilter) validateNamespaceID(formats strfmt.Registry) error {

	if err := validate.Required("namespaceId", "body", m.NamespaceID); err != nil {
		return err
	}

	return nil
}

func (m *UserRouteFilter) validatePage(formats strfmt.Registry) error {

	if err := validate.Required("page", "body", m.Page); err != nil {
		return err
	}

	return nil
}

func (m *UserRouteFilter) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user route filter based on context it is used
func (m *UserRouteFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserRouteFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserRouteFilter) UnmarshalBinary(b []byte) error {
	var res UserRouteFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
