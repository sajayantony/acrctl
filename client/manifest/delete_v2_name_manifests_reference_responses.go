// Code generated by go-swagger; DO NOT EDIT.

package manifest

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

// DeleteV2NameManifestsReferenceReader is a Reader for the DeleteV2NameManifestsReference structure.
type DeleteV2NameManifestsReferenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV2NameManifestsReferenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewDeleteV2NameManifestsReferenceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteV2NameManifestsReferenceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteV2NameManifestsReferenceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteV2NameManifestsReferenceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteV2NameManifestsReferenceAccepted creates a DeleteV2NameManifestsReferenceAccepted with default headers values
func NewDeleteV2NameManifestsReferenceAccepted() *DeleteV2NameManifestsReferenceAccepted {
	return &DeleteV2NameManifestsReferenceAccepted{}
}

/*DeleteV2NameManifestsReferenceAccepted handles this case with default header values.

The manifest has been accepted by the registry and is stored under the specified name and tag.
*/
type DeleteV2NameManifestsReferenceAccepted struct {
}

func (o *DeleteV2NameManifestsReferenceAccepted) Error() string {
	return fmt.Sprintf("[DELETE /v2/{name}/manifests/{reference}][%d] deleteV2NameManifestsReferenceAccepted ", 202)
}

func (o *DeleteV2NameManifestsReferenceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteV2NameManifestsReferenceBadRequest creates a DeleteV2NameManifestsReferenceBadRequest with default headers values
func NewDeleteV2NameManifestsReferenceBadRequest() *DeleteV2NameManifestsReferenceBadRequest {
	return &DeleteV2NameManifestsReferenceBadRequest{}
}

/*DeleteV2NameManifestsReferenceBadRequest handles this case with default header values.

On failure
*/
type DeleteV2NameManifestsReferenceBadRequest struct {
	Payload *DeleteV2NameManifestsReferenceBadRequestBody
}

func (o *DeleteV2NameManifestsReferenceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v2/{name}/manifests/{reference}][%d] deleteV2NameManifestsReferenceBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteV2NameManifestsReferenceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteV2NameManifestsReferenceBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteV2NameManifestsReferenceUnauthorized creates a DeleteV2NameManifestsReferenceUnauthorized with default headers values
func NewDeleteV2NameManifestsReferenceUnauthorized() *DeleteV2NameManifestsReferenceUnauthorized {
	return &DeleteV2NameManifestsReferenceUnauthorized{}
}

/*DeleteV2NameManifestsReferenceUnauthorized handles this case with default header values.

Unauthorized access
*/
type DeleteV2NameManifestsReferenceUnauthorized struct {
	Payload *DeleteV2NameManifestsReferenceUnauthorizedBody
}

func (o *DeleteV2NameManifestsReferenceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v2/{name}/manifests/{reference}][%d] deleteV2NameManifestsReferenceUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteV2NameManifestsReferenceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteV2NameManifestsReferenceUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteV2NameManifestsReferenceNotFound creates a DeleteV2NameManifestsReferenceNotFound with default headers values
func NewDeleteV2NameManifestsReferenceNotFound() *DeleteV2NameManifestsReferenceNotFound {
	return &DeleteV2NameManifestsReferenceNotFound{}
}

/*DeleteV2NameManifestsReferenceNotFound handles this case with default header values.

The named manifest could not be found in the Registry
*/
type DeleteV2NameManifestsReferenceNotFound struct {
	Payload *DeleteV2NameManifestsReferenceNotFoundBody
}

func (o *DeleteV2NameManifestsReferenceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v2/{name}/manifests/{reference}][%d] deleteV2NameManifestsReferenceNotFound  %+v", 404, o.Payload)
}

func (o *DeleteV2NameManifestsReferenceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteV2NameManifestsReferenceNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteV2NameManifestsReferenceBadRequestBody delete v2 name manifests reference bad request body
swagger:model DeleteV2NameManifestsReferenceBadRequestBody
*/
type DeleteV2NameManifestsReferenceBadRequestBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this delete v2 name manifests reference bad request body
func (o *DeleteV2NameManifestsReferenceBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteV2NameManifestsReferenceBadRequestBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteV2NameManifestsReferenceBadRequest" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteV2NameManifestsReferenceBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteV2NameManifestsReferenceBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DeleteV2NameManifestsReferenceBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteV2NameManifestsReferenceNotFoundBody delete v2 name manifests reference not found body
swagger:model DeleteV2NameManifestsReferenceNotFoundBody
*/
type DeleteV2NameManifestsReferenceNotFoundBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this delete v2 name manifests reference not found body
func (o *DeleteV2NameManifestsReferenceNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteV2NameManifestsReferenceNotFoundBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteV2NameManifestsReferenceNotFound" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteV2NameManifestsReferenceNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteV2NameManifestsReferenceNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteV2NameManifestsReferenceNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteV2NameManifestsReferenceUnauthorizedBody delete v2 name manifests reference unauthorized body
swagger:model DeleteV2NameManifestsReferenceUnauthorizedBody
*/
type DeleteV2NameManifestsReferenceUnauthorizedBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this delete v2 name manifests reference unauthorized body
func (o *DeleteV2NameManifestsReferenceUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteV2NameManifestsReferenceUnauthorizedBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteV2NameManifestsReferenceUnauthorized" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteV2NameManifestsReferenceUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteV2NameManifestsReferenceUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res DeleteV2NameManifestsReferenceUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
