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

// GetAcrV1NameMetadataMetadataReader is a Reader for the GetAcrV1NameMetadataMetadata structure.
type GetAcrV1NameMetadataMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAcrV1NameMetadataMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAcrV1NameMetadataMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAcrV1NameMetadataMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAcrV1NameMetadataMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAcrV1NameMetadataMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAcrV1NameMetadataMetadataOK creates a GetAcrV1NameMetadataMetadataOK with default headers values
func NewGetAcrV1NameMetadataMetadataOK() *GetAcrV1NameMetadataMetadataOK {
	return &GetAcrV1NameMetadataMetadataOK{}
}

/*GetAcrV1NameMetadataMetadataOK handles this case with default header values.

Returns the metadata value
*/
type GetAcrV1NameMetadataMetadataOK struct {
	Payload string
}

func (o *GetAcrV1NameMetadataMetadataOK) Error() string {
	return fmt.Sprintf("[GET /acr/v1/{name}/_metadata/{metadata}][%d] getAcrV1NameMetadataMetadataOK  %+v", 200, o.Payload)
}

func (o *GetAcrV1NameMetadataMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAcrV1NameMetadataMetadataBadRequest creates a GetAcrV1NameMetadataMetadataBadRequest with default headers values
func NewGetAcrV1NameMetadataMetadataBadRequest() *GetAcrV1NameMetadataMetadataBadRequest {
	return &GetAcrV1NameMetadataMetadataBadRequest{}
}

/*GetAcrV1NameMetadataMetadataBadRequest handles this case with default header values.

On failure
*/
type GetAcrV1NameMetadataMetadataBadRequest struct {
	Payload *GetAcrV1NameMetadataMetadataBadRequestBody
}

func (o *GetAcrV1NameMetadataMetadataBadRequest) Error() string {
	return fmt.Sprintf("[GET /acr/v1/{name}/_metadata/{metadata}][%d] getAcrV1NameMetadataMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *GetAcrV1NameMetadataMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAcrV1NameMetadataMetadataBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAcrV1NameMetadataMetadataUnauthorized creates a GetAcrV1NameMetadataMetadataUnauthorized with default headers values
func NewGetAcrV1NameMetadataMetadataUnauthorized() *GetAcrV1NameMetadataMetadataUnauthorized {
	return &GetAcrV1NameMetadataMetadataUnauthorized{}
}

/*GetAcrV1NameMetadataMetadataUnauthorized handles this case with default header values.

Unauthorized access
*/
type GetAcrV1NameMetadataMetadataUnauthorized struct {
	Payload *GetAcrV1NameMetadataMetadataUnauthorizedBody
}

func (o *GetAcrV1NameMetadataMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[GET /acr/v1/{name}/_metadata/{metadata}][%d] getAcrV1NameMetadataMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAcrV1NameMetadataMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAcrV1NameMetadataMetadataUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAcrV1NameMetadataMetadataNotFound creates a GetAcrV1NameMetadataMetadataNotFound with default headers values
func NewGetAcrV1NameMetadataMetadataNotFound() *GetAcrV1NameMetadataMetadataNotFound {
	return &GetAcrV1NameMetadataMetadataNotFound{}
}

/*GetAcrV1NameMetadataMetadataNotFound handles this case with default header values.

The repository or the metadata is unknown to the registry.
*/
type GetAcrV1NameMetadataMetadataNotFound struct {
	Payload *GetAcrV1NameMetadataMetadataNotFoundBody
}

func (o *GetAcrV1NameMetadataMetadataNotFound) Error() string {
	return fmt.Sprintf("[GET /acr/v1/{name}/_metadata/{metadata}][%d] getAcrV1NameMetadataMetadataNotFound  %+v", 404, o.Payload)
}

func (o *GetAcrV1NameMetadataMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAcrV1NameMetadataMetadataNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAcrV1NameMetadataMetadataBadRequestBody get acr v1 name metadata metadata bad request body
swagger:model GetAcrV1NameMetadataMetadataBadRequestBody
*/
type GetAcrV1NameMetadataMetadataBadRequestBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get acr v1 name metadata metadata bad request body
func (o *GetAcrV1NameMetadataMetadataBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAcrV1NameMetadataMetadataBadRequestBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAcrV1NameMetadataMetadataBadRequest" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAcrV1NameMetadataMetadataBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAcrV1NameMetadataMetadataBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetAcrV1NameMetadataMetadataBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAcrV1NameMetadataMetadataNotFoundBody get acr v1 name metadata metadata not found body
swagger:model GetAcrV1NameMetadataMetadataNotFoundBody
*/
type GetAcrV1NameMetadataMetadataNotFoundBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get acr v1 name metadata metadata not found body
func (o *GetAcrV1NameMetadataMetadataNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAcrV1NameMetadataMetadataNotFoundBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAcrV1NameMetadataMetadataNotFound" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAcrV1NameMetadataMetadataNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAcrV1NameMetadataMetadataNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetAcrV1NameMetadataMetadataNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAcrV1NameMetadataMetadataUnauthorizedBody get acr v1 name metadata metadata unauthorized body
swagger:model GetAcrV1NameMetadataMetadataUnauthorizedBody
*/
type GetAcrV1NameMetadataMetadataUnauthorizedBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get acr v1 name metadata metadata unauthorized body
func (o *GetAcrV1NameMetadataMetadataUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAcrV1NameMetadataMetadataUnauthorizedBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAcrV1NameMetadataMetadataUnauthorized" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAcrV1NameMetadataMetadataUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAcrV1NameMetadataMetadataUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetAcrV1NameMetadataMetadataUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}