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

// GroupInsertRequest group insert request
//
// swagger:model groupInsertRequest
type GroupInsertRequest struct {

	// base group Id
	BaseGroupID int64 `json:"baseGroupId,omitempty"`

	// code
	// Required: true
	Code *string `json:"code"`

	// description
	// Required: true
	Description *string `json:"description"`

	// namespace Id
	// Required: true
	NamespaceID *int64 `json:"namespaceId"`
}

// Validate validates this group insert request
func (m *GroupInsertRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
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

func (m *GroupInsertRequest) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *GroupInsertRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *GroupInsertRequest) validateNamespaceID(formats strfmt.Registry) error {

	if err := validate.Required("namespaceId", "body", m.NamespaceID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this group insert request based on context it is used
func (m *GroupInsertRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GroupInsertRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupInsertRequest) UnmarshalBinary(b []byte) error {
	var res GroupInsertRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
