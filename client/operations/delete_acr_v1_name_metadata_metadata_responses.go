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

// DeleteAcrV1NameMetadataMetadataReader is a Reader for the DeleteAcrV1NameMetadataMetadata structure.
type DeleteAcrV1NameMetadataMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAcrV1NameMetadataMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewDeleteAcrV1NameMetadataMetadataAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteAcrV1NameMetadataMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteAcrV1NameMetadataMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteAcrV1NameMetadataMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAcrV1NameMetadataMetadataAccepted creates a DeleteAcrV1NameMetadataMetadataAccepted with default headers values
func NewDeleteAcrV1NameMetadataMetadataAccepted() *DeleteAcrV1NameMetadataMetadataAccepted {
	return &DeleteAcrV1NameMetadataMetadataAccepted{}
}

/*DeleteAcrV1NameMetadataMetadataAccepted handles this case with default header values.

metadata is deleted
*/
type DeleteAcrV1NameMetadataMetadataAccepted struct {
}

func (o *DeleteAcrV1NameMetadataMetadataAccepted) Error() string {
	return fmt.Sprintf("[DELETE /acr/v1/{name}/_metadata/{metadata}][%d] deleteAcrV1NameMetadataMetadataAccepted ", 202)
}

func (o *DeleteAcrV1NameMetadataMetadataAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAcrV1NameMetadataMetadataBadRequest creates a DeleteAcrV1NameMetadataMetadataBadRequest with default headers values
func NewDeleteAcrV1NameMetadataMetadataBadRequest() *DeleteAcrV1NameMetadataMetadataBadRequest {
	return &DeleteAcrV1NameMetadataMetadataBadRequest{}
}

/*DeleteAcrV1NameMetadataMetadataBadRequest handles this case with default header values.

On failure
*/
type DeleteAcrV1NameMetadataMetadataBadRequest struct {
	Payload *DeleteAcrV1NameMetadataMetadataBadRequestBody
}

func (o *DeleteAcrV1NameMetadataMetadataBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /acr/v1/{name}/_metadata/{metadata}][%d] deleteAcrV1NameMetadataMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAcrV1NameMetadataMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteAcrV1NameMetadataMetadataBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAcrV1NameMetadataMetadataUnauthorized creates a DeleteAcrV1NameMetadataMetadataUnauthorized with default headers values
func NewDeleteAcrV1NameMetadataMetadataUnauthorized() *DeleteAcrV1NameMetadataMetadataUnauthorized {
	return &DeleteAcrV1NameMetadataMetadataUnauthorized{}
}

/*DeleteAcrV1NameMetadataMetadataUnauthorized handles this case with default header values.

Unauthorized access
*/
type DeleteAcrV1NameMetadataMetadataUnauthorized struct {
	Payload *DeleteAcrV1NameMetadataMetadataUnauthorizedBody
}

func (o *DeleteAcrV1NameMetadataMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /acr/v1/{name}/_metadata/{metadata}][%d] deleteAcrV1NameMetadataMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAcrV1NameMetadataMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteAcrV1NameMetadataMetadataUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAcrV1NameMetadataMetadataNotFound creates a DeleteAcrV1NameMetadataMetadataNotFound with default headers values
func NewDeleteAcrV1NameMetadataMetadataNotFound() *DeleteAcrV1NameMetadataMetadataNotFound {
	return &DeleteAcrV1NameMetadataMetadataNotFound{}
}

/*DeleteAcrV1NameMetadataMetadataNotFound handles this case with default header values.

The repository or metadata is unknown to the registry.
*/
type DeleteAcrV1NameMetadataMetadataNotFound struct {
	Payload *DeleteAcrV1NameMetadataMetadataNotFoundBody
}

func (o *DeleteAcrV1NameMetadataMetadataNotFound) Error() string {
	return fmt.Sprintf("[DELETE /acr/v1/{name}/_metadata/{metadata}][%d] deleteAcrV1NameMetadataMetadataNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAcrV1NameMetadataMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteAcrV1NameMetadataMetadataNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteAcrV1NameMetadataMetadataBadRequestBody delete acr v1 name metadata metadata bad request body
swagger:model DeleteAcrV1NameMetadataMetadataBadRequestBody
*/
type DeleteAcrV1NameMetadataMetadataBadRequestBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this delete acr v1 name metadata metadata bad request body
func (o *DeleteAcrV1NameMetadataMetadataBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteAcrV1NameMetadataMetadataBadRequestBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteAcrV1NameMetadataMetadataBadRequest" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteAcrV1NameMetadataMetadataBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteAcrV1NameMetadataMetadataBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DeleteAcrV1NameMetadataMetadataBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteAcrV1NameMetadataMetadataNotFoundBody delete acr v1 name metadata metadata not found body
swagger:model DeleteAcrV1NameMetadataMetadataNotFoundBody
*/
type DeleteAcrV1NameMetadataMetadataNotFoundBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this delete acr v1 name metadata metadata not found body
func (o *DeleteAcrV1NameMetadataMetadataNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteAcrV1NameMetadataMetadataNotFoundBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteAcrV1NameMetadataMetadataNotFound" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteAcrV1NameMetadataMetadataNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteAcrV1NameMetadataMetadataNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteAcrV1NameMetadataMetadataNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteAcrV1NameMetadataMetadataUnauthorizedBody delete acr v1 name metadata metadata unauthorized body
swagger:model DeleteAcrV1NameMetadataMetadataUnauthorizedBody
*/
type DeleteAcrV1NameMetadataMetadataUnauthorizedBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this delete acr v1 name metadata metadata unauthorized body
func (o *DeleteAcrV1NameMetadataMetadataUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteAcrV1NameMetadataMetadataUnauthorizedBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteAcrV1NameMetadataMetadataUnauthorized" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteAcrV1NameMetadataMetadataUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteAcrV1NameMetadataMetadataUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res DeleteAcrV1NameMetadataMetadataUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
