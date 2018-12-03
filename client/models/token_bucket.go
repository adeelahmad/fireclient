// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TokenBucket Defines a token bucket with a maximum capacity (size), an initial burst size (one_time_burst) and an interval for refilling purposes (refill_time). The refill-rate is derived from size and refill_time, and it is the constant rate at which the tokens replenish. The refill process only starts happening after the initial burst budget is consumed. Consumption from the token bucket is unbounded in speed which allows for bursts bound in size by the amount of tokens available. Once the token bucket is empty, consumption speed is bound by the refill_rate.
// swagger:model TokenBucket
type TokenBucket struct {

	// The initial size of a token bucket.
	// Minimum: 0
	OneTimeBurst *int64 `json:"one_time_burst,omitempty"`

	// The amount of milliseconds it takes for the bucket to refill.
	// Minimum: 0
	RefillTime *int64 `json:"refill_time,omitempty"`

	// The total number of tokens this bucket can hold.
	// Minimum: 0
	Size *int64 `json:"size,omitempty"`
}

// Validate validates this token bucket
func (m *TokenBucket) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOneTimeBurst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefillTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenBucket) validateOneTimeBurst(formats strfmt.Registry) error {

	if swag.IsZero(m.OneTimeBurst) { // not required
		return nil
	}

	if err := validate.MinimumInt("one_time_burst", "body", int64(*m.OneTimeBurst), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *TokenBucket) validateRefillTime(formats strfmt.Registry) error {

	if swag.IsZero(m.RefillTime) { // not required
		return nil
	}

	if err := validate.MinimumInt("refill_time", "body", int64(*m.RefillTime), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *TokenBucket) validateSize(formats strfmt.Registry) error {

	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if err := validate.MinimumInt("size", "body", int64(*m.Size), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TokenBucket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenBucket) UnmarshalBinary(b []byte) error {
	var res TokenBucket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
