// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemSystem system system
//
// swagger:model system.System
type SystemSystem struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// deleted at
	DeletedAt *string `json:"deletedAt,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this system system
func (m *SystemSystem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this system system based on context it is used
func (m *SystemSystem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemSystem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemSystem) UnmarshalBinary(b []byte) error {
	var res SystemSystem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
