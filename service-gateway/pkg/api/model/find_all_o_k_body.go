// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FindAllOKBody find all o k body
// swagger:model findAllOKBody
type FindAllOKBody struct {

	// pagination
	Pagination *Pagination `json:"pagination,omitempty"`

	// rows
	Rows FindAllOKBodyRows `json:"rows"`
}

// Validate validates this find all o k body
func (m *FindAllOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagination(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FindAllOKBody) validatePagination(formats strfmt.Registry) error {

	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {

		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FindAllOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FindAllOKBody) UnmarshalBinary(b []byte) error {
	var res FindAllOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
