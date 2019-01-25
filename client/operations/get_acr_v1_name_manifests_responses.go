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

// GetAcrV1NameManifestsReader is a Reader for the GetAcrV1NameManifests structure.
type GetAcrV1NameManifestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAcrV1NameManifestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAcrV1NameManifestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAcrV1NameManifestsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAcrV1NameManifestsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAcrV1NameManifestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAcrV1NameManifestsOK creates a GetAcrV1NameManifestsOK with default headers values
func NewGetAcrV1NameManifestsOK() *GetAcrV1NameManifestsOK {
	return &GetAcrV1NameManifestsOK{}
}

/*GetAcrV1NameManifestsOK handles this case with default header values.

Returns a list of manifests
*/
type GetAcrV1NameManifestsOK struct {
	Payload *GetAcrV1NameManifestsOKBody
}

func (o *GetAcrV1NameManifestsOK) Error() string {
	return fmt.Sprintf("[GET /acr/v1/{name}/_manifests][%d] getAcrV1NameManifestsOK  %+v", 200, o.Payload)
}

func (o *GetAcrV1NameManifestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAcrV1NameManifestsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAcrV1NameManifestsBadRequest creates a GetAcrV1NameManifestsBadRequest with default headers values
func NewGetAcrV1NameManifestsBadRequest() *GetAcrV1NameManifestsBadRequest {
	return &GetAcrV1NameManifestsBadRequest{}
}

/*GetAcrV1NameManifestsBadRequest handles this case with default header values.

On failure
*/
type GetAcrV1NameManifestsBadRequest struct {
	Payload *GetAcrV1NameManifestsBadRequestBody
}

func (o *GetAcrV1NameManifestsBadRequest) Error() string {
	return fmt.Sprintf("[GET /acr/v1/{name}/_manifests][%d] getAcrV1NameManifestsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAcrV1NameManifestsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAcrV1NameManifestsBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAcrV1NameManifestsUnauthorized creates a GetAcrV1NameManifestsUnauthorized with default headers values
func NewGetAcrV1NameManifestsUnauthorized() *GetAcrV1NameManifestsUnauthorized {
	return &GetAcrV1NameManifestsUnauthorized{}
}

/*GetAcrV1NameManifestsUnauthorized handles this case with default header values.

Unauthorized access
*/
type GetAcrV1NameManifestsUnauthorized struct {
	Payload *GetAcrV1NameManifestsUnauthorizedBody
}

func (o *GetAcrV1NameManifestsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /acr/v1/{name}/_manifests][%d] getAcrV1NameManifestsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAcrV1NameManifestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAcrV1NameManifestsUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAcrV1NameManifestsNotFound creates a GetAcrV1NameManifestsNotFound with default headers values
func NewGetAcrV1NameManifestsNotFound() *GetAcrV1NameManifestsNotFound {
	return &GetAcrV1NameManifestsNotFound{}
}

/*GetAcrV1NameManifestsNotFound handles this case with default header values.

The repository is unknown to the registry.
*/
type GetAcrV1NameManifestsNotFound struct {
	Payload *GetAcrV1NameManifestsNotFoundBody
}

func (o *GetAcrV1NameManifestsNotFound) Error() string {
	return fmt.Sprintf("[GET /acr/v1/{name}/_manifests][%d] getAcrV1NameManifestsNotFound  %+v", 404, o.Payload)
}

func (o *GetAcrV1NameManifestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAcrV1NameManifestsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAcrV1NameManifestsBadRequestBody get acr v1 name manifests bad request body
swagger:model GetAcrV1NameManifestsBadRequestBody
*/
type GetAcrV1NameManifestsBadRequestBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get acr v1 name manifests bad request body
func (o *GetAcrV1NameManifestsBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAcrV1NameManifestsBadRequestBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAcrV1NameManifestsBadRequest" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAcrV1NameManifestsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAcrV1NameManifestsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetAcrV1NameManifestsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAcrV1NameManifestsNotFoundBody get acr v1 name manifests not found body
swagger:model GetAcrV1NameManifestsNotFoundBody
*/
type GetAcrV1NameManifestsNotFoundBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get acr v1 name manifests not found body
func (o *GetAcrV1NameManifestsNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAcrV1NameManifestsNotFoundBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAcrV1NameManifestsNotFound" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAcrV1NameManifestsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAcrV1NameManifestsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetAcrV1NameManifestsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAcrV1NameManifestsOKBody get acr v1 name manifests o k body
swagger:model GetAcrV1NameManifestsOKBody
*/
type GetAcrV1NameManifestsOKBody struct {

	// data
	Data *models.ManifestAttributeList `json:"data,omitempty"`
}

// Validate validates this get acr v1 name manifests o k body
func (o *GetAcrV1NameManifestsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAcrV1NameManifestsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAcrV1NameManifestsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAcrV1NameManifestsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAcrV1NameManifestsOKBody) UnmarshalBinary(b []byte) error {
	var res GetAcrV1NameManifestsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAcrV1NameManifestsUnauthorizedBody get acr v1 name manifests unauthorized body
swagger:model GetAcrV1NameManifestsUnauthorizedBody
*/
type GetAcrV1NameManifestsUnauthorizedBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get acr v1 name manifests unauthorized body
func (o *GetAcrV1NameManifestsUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAcrV1NameManifestsUnauthorizedBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAcrV1NameManifestsUnauthorized" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAcrV1NameManifestsUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAcrV1NameManifestsUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetAcrV1NameManifestsUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}