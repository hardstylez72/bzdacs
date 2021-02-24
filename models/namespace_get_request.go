// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NamespaceGetRequest namespace get request
//
// swagger:model namespaceGetRequest
type NamespaceGetRequest struct {

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// system Id
	SystemID int64 `json:"systemId,omitempty"`
}

// Validate validates this namespace get request
func (m *NamespaceGetRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this namespace get request based on context it is used
func (m *NamespaceGetRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NamespaceGetRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamespaceGetRequest) UnmarshalBinary(b []byte) error {
	var res NamespaceGetRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
