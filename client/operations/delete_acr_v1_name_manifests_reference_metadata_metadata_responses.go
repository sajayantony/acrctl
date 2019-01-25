// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sajayantony/acrctl/models"
)

// DeleteAcrV1NameManifestsReferenceMetadataMetadataReader is a Reader for the DeleteAcrV1NameManifestsReferenceMetadataMetadata structure.
type DeleteAcrV1NameManifestsReferenceMetadataMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewDeleteAcrV1NameManifestsReferenceMetadataMetadataAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteAcrV1NameManifestsReferenceMetadataMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAcrV1NameManifestsReferenceMetadataMetadataAccepted creates a DeleteAcrV1NameManifestsReferenceMetadataMetadataAccepted with default headers values
func NewDeleteAcrV1NameManifestsReferenceMetadataMetadataAccepted() *DeleteAcrV1NameManifestsReferenceMetadataMetadataAccepted {
	return &DeleteAcrV1NameManifestsReferenceMetadataMetadataAccepted{}
}

/*DeleteAcrV1NameManifestsReferenceMetadataMetadataAccepted handles this case with default header values.

metadata is deleted
*/
type DeleteAcrV1NameManifestsReferenceMetadataMetadataAccepted struct {
}

func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataAccepted) Error() string {
	return fmt.Sprintf("[DELETE /acr/v1/{name}/_manifests/{reference}/_metadata/{metadata}][%d] deleteAcrV1NameManifestsReferenceMetadataMetadataAccepted ", 202)
}

func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequest creates a DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequest with default headers values
func NewDeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequest() *DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequest {
	return &DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequest{}
}

/*DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequest handles this case with default header values.

On failure
*/
type DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequest struct {
	Payload *DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequestBody
}

func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /acr/v1/{name}/_manifests/{reference}/_metadata/{metadata}][%d] deleteAcrV1NameManifestsReferenceMetadataMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorized creates a DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorized with default headers values
func NewDeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorized() *DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorized {
	return &DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorized{}
}

/*DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorized handles this case with default header values.

Unauthorized access
*/
type DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorized struct {
	Payload *DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorizedBody
}

func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /acr/v1/{name}/_manifests/{reference}/_metadata/{metadata}][%d] deleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAcrV1NameManifestsReferenceMetadataMetadataNotFound creates a DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFound with default headers values
func NewDeleteAcrV1NameManifestsReferenceMetadataMetadataNotFound() *DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFound {
	return &DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFound{}
}

/*DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFound handles this case with default header values.

The repository, manifest or metadata is unknown to the registry.
*/
type DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFound struct {
	Payload *DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFoundBody
}

func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFound) Error() string {
	return fmt.Sprintf("[DELETE /acr/v1/{name}/_manifests/{reference}/_metadata/{metadata}][%d] deleteAcrV1NameManifestsReferenceMetadataMetadataNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequestBody delete acr v1 name manifests reference metadata metadata bad request body
swagger:model DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequestBody
*/
type DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequestBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this delete acr v1 name manifests reference metadata metadata bad request body
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequestBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteAcrV1NameManifestsReferenceMetadataMetadataBadRequest" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DeleteAcrV1NameManifestsReferenceMetadataMetadataBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFoundBody delete acr v1 name manifests reference metadata metadata not found body
swagger:model DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFoundBody
*/
type DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFoundBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this delete acr v1 name manifests reference metadata metadata not found body
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFoundBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteAcrV1NameManifestsReferenceMetadataMetadataNotFound" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteAcrV1NameManifestsReferenceMetadataMetadataNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorizedBody delete acr v1 name manifests reference metadata metadata unauthorized body
swagger:model DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorizedBody
*/
type DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorizedBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this delete acr v1 name manifests reference metadata metadata unauthorized body
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorizedBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorized" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res DeleteAcrV1NameManifestsReferenceMetadataMetadataUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
