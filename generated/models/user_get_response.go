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

// UserGetResponse user get response
//
// swagger:model userGetResponse
type UserGetResponse struct {

	// created at
	// Required: true
	CreatedAt *string `json:"createdAt"`

	// deleted at
	DeletedAt *string `json:"deletedAt,omitempty"`

	// external Id
	// Required: true
	ExternalID *string `json:"externalId"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// namespace Id
	// Required: true
	NamespaceID *int64 `json:"namespaceId"`

	// updated at
	// Required: true
	UpdatedAt *string `json:"updatedAt"`
}

// Validate validates this user get response
func (m *UserGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserGetResponse) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *UserGetResponse) validateExternalID(formats strfmt.Registry) error {

	if err := validate.Required("externalId", "body", m.ExternalID); err != nil {
		return err
	}

	return nil
}

func (m *UserGetResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *UserGetResponse) validateNamespaceID(formats strfmt.Registry) error {

	if err := validate.Required("namespaceId", "body", m.NamespaceID); err != nil {
		return err
	}

	return nil
}

func (m *UserGetResponse) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user get response based on context it is used
func (m *UserGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserGetResponse) UnmarshalBinary(b []byte) error {
	var res UserGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
