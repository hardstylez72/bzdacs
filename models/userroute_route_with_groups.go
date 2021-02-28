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
	"github.com/go-openapi/validate"
)

// UserrouteRouteWithGroups userroute route with groups
//
// swagger:model userroute.RouteWithGroups
type UserrouteRouteWithGroups struct {

	// created at
	// Required: true
	CreatedAt *string `json:"createdAt"`

	// deleted at
	DeletedAt *string `json:"deletedAt,omitempty"`

	// description
	// Required: true
	Description *string `json:"description"`

	// groups
	Groups []*UserrouteGroup `json:"groups"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// is excluded
	IsExcluded bool `json:"isExcluded,omitempty"`

	// method
	// Required: true
	Method *string `json:"method"`

	// namespace Id
	// Required: true
	NamespaceID *int64 `json:"namespaceId"`

	// route
	// Required: true
	Route *string `json:"route"`

	// tags
	// Required: true
	Tags []string `json:"tags"`

	// updated at
	// Required: true
	UpdatedAt *string `json:"updatedAt"`
}

// Validate validates this userroute route with groups
func (m *UserrouteRouteWithGroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
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

func (m *UserrouteRouteWithGroups) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *UserrouteRouteWithGroups) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *UserrouteRouteWithGroups) validateGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserrouteRouteWithGroups) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *UserrouteRouteWithGroups) validateMethod(formats strfmt.Registry) error {

	if err := validate.Required("method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

func (m *UserrouteRouteWithGroups) validateNamespaceID(formats strfmt.Registry) error {

	if err := validate.Required("namespaceId", "body", m.NamespaceID); err != nil {
		return err
	}

	return nil
}

func (m *UserrouteRouteWithGroups) validateRoute(formats strfmt.Registry) error {

	if err := validate.Required("route", "body", m.Route); err != nil {
		return err
	}

	return nil
}

func (m *UserrouteRouteWithGroups) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

func (m *UserrouteRouteWithGroups) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this userroute route with groups based on the context it is used
func (m *UserrouteRouteWithGroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserrouteRouteWithGroups) contextValidateGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Groups); i++ {

		if m.Groups[i] != nil {
			if err := m.Groups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserrouteRouteWithGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserrouteRouteWithGroups) UnmarshalBinary(b []byte) error {
	var res UserrouteRouteWithGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
