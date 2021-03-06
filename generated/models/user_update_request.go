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

// UserUpdateRequest user update request
//
// swagger:model userUpdateRequest
type UserUpdateRequest struct {

	// external Id
	// Required: true
	ExternalID *string `json:"externalId"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// namespace Id
	// Required: true
	NamespaceID *int64 `json:"namespaceId"`
}

// Validate validates this user update request
func (m *UserUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

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

func (m *UserUpdateRequest) validateExternalID(formats strfmt.Registry) error {

	if err := validate.Required("externalId", "body", m.ExternalID); err != nil {
		return err
	}

	return nil
}

func (m *UserUpdateRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *UserUpdateRequest) validateNamespaceID(formats strfmt.Registry) error {

	if err := validate.Required("namespaceId", "body", m.NamespaceID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user update request based on context it is used
func (m *UserUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserUpdateRequest) UnmarshalBinary(b []byte) error {
	var res UserUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
