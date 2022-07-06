// Code generated by go-swagger; DO NOT EDIT.

package experiment_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V2beta1Experiment v2beta1 experiment
// swagger:model v2beta1Experiment
type V2beta1Experiment struct {

	// Output. The time that the experiment created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Optional input field. Describing the purpose of the experiment
	Description string `json:"description,omitempty"`

	// Output. Unique experiment ID. Generated by API server.
	ID string `json:"id,omitempty"`

	// Required input field. Unique experiment name provided by user.
	Name string `json:"name,omitempty"`

	// Optional input field. Specify the namespace this experiment belongs to.
	Namespace string `json:"namespace,omitempty"`

	// Output. Specifies whether this experiment is in archived or available state.
	StorageState ExperimentStorageState `json:"storage_state,omitempty"`
}

// Validate validates this v2beta1 experiment
func (m *V2beta1Experiment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2beta1Experiment) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V2beta1Experiment) validateStorageState(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageState) { // not required
		return nil
	}

	if err := m.StorageState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("storage_state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2beta1Experiment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2beta1Experiment) UnmarshalBinary(b []byte) error {
	var res V2beta1Experiment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
