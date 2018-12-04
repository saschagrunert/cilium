// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DebugInfo groups some debugging related information on the agent
// swagger:model DebugInfo

type DebugInfo struct {

	// cilium memory map
	CiliumMemoryMap string `json:"cilium-memory-map,omitempty"`

	// cilium nodemonitor memory map
	CiliumNodemonitorMemoryMap string `json:"cilium-nodemonitor-memory-map,omitempty"`

	// cilium status
	CiliumStatus *StatusResponse `json:"cilium-status,omitempty"`

	// cilium version
	CiliumVersion string `json:"cilium-version,omitempty"`

	// endpoint list
	EndpointList []*Endpoint `json:"endpoint-list"`

	// environment variables
	EnvironmentVariables []string `json:"environment-variables"`

	// kernel version
	KernelVersion string `json:"kernel-version,omitempty"`

	// policy
	Policy *Policy `json:"policy,omitempty"`

	// service list
	ServiceList []*Service `json:"service-list"`

	// subsystem
	Subsystem map[string]string `json:"subsystem,omitempty"`
}

/* polymorph DebugInfo cilium-memory-map false */

/* polymorph DebugInfo cilium-nodemonitor-memory-map false */

/* polymorph DebugInfo cilium-status false */

/* polymorph DebugInfo cilium-version false */

/* polymorph DebugInfo endpoint-list false */

/* polymorph DebugInfo environment-variables false */

/* polymorph DebugInfo kernel-version false */

/* polymorph DebugInfo policy false */

/* polymorph DebugInfo service-list false */

/* polymorph DebugInfo subsystem false */

// Validate validates this debug info
func (m *DebugInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCiliumStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEndpointList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEnvironmentVariables(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServiceList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DebugInfo) validateCiliumStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.CiliumStatus) { // not required
		return nil
	}

	if m.CiliumStatus != nil {

		if err := m.CiliumStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cilium-status")
			}
			return err
		}
	}

	return nil
}

func (m *DebugInfo) validateEndpointList(formats strfmt.Registry) error {

	if swag.IsZero(m.EndpointList) { // not required
		return nil
	}

	for i := 0; i < len(m.EndpointList); i++ {

		if swag.IsZero(m.EndpointList[i]) { // not required
			continue
		}

		if m.EndpointList[i] != nil {

			if err := m.EndpointList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("endpoint-list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DebugInfo) validateEnvironmentVariables(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvironmentVariables) { // not required
		return nil
	}

	return nil
}

func (m *DebugInfo) validatePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if m.Policy != nil {

		if err := m.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

func (m *DebugInfo) validateServiceList(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceList) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceList); i++ {

		if swag.IsZero(m.ServiceList[i]) { // not required
			continue
		}

		if m.ServiceList[i] != nil {

			if err := m.ServiceList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service-list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DebugInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DebugInfo) UnmarshalBinary(b []byte) error {
	var res DebugInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
