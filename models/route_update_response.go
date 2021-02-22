// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RouteUpdateResponse route update response
//
// swagger:model route.updateResponse
type RouteUpdateResponse struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// deleted at
	DeletedAt *string `json:"deletedAt,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// method
	Method string `json:"method,omitempty"`

	// namespace Id
	NamespaceID int64 `json:"namespaceId,omitempty"`

	// route
	Route string `json:"route,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this route update response
func (m *RouteUpdateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this route update response based on context it is used
func (m *RouteUpdateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RouteUpdateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RouteUpdateResponse) UnmarshalBinary(b []byte) error {
	var res RouteUpdateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
