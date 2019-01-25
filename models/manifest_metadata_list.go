// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ManifestMetadataList manifest metadata list
// swagger:model ManifestMetadataList
type ManifestMetadataList struct {

	// digest
	Digest string `json:"digest,omitempty"`

	// image name
	ImageName string `json:"imageName,omitempty"`

	// metadata
	Metadata []string `json:"metadata"`

	// registry
	Registry string `json:"registry,omitempty"`
}

// Validate validates this manifest metadata list
func (m *ManifestMetadataList) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ManifestMetadataList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManifestMetadataList) UnmarshalBinary(b []byte) error {
	var res ManifestMetadataList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}