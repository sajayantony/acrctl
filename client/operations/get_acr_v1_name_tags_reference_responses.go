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

// GetAcrV1NameTagsReferenceReader is a Reader for the GetAcrV1NameTagsReference structure.
type GetAcrV1NameTagsReferenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAcrV1NameTagsReferenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAcrV1NameTagsReferenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAcrV1NameTagsReferenceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAcrV1NameTagsReferenceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAcrV1NameTagsReferenceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAcrV1NameTagsReferenceOK creates a GetAcrV1NameTagsReferenceOK with default headers values
func NewGetAcrV1NameTagsReferenceOK() *GetAcrV1NameTagsReferenceOK {
	return &GetAcrV1NameTagsReferenceOK{}
}

/*GetAcrV1NameTagsReferenceOK handles this case with default header values.

Returns attributes of a tag
*/
type GetAcrV1NameTagsReferenceOK struct {
	Payload *GetAcrV1NameTagsReferenceOKBody
}

func (o *GetAcrV1NameTagsReferenceOK) Error() string {
	return fmt.Sprintf("[GET /acr/v1/{name}/_tags/{reference}][%d] getAcrV1NameTagsReferenceOK  %+v", 200, o.Payload)
}

func (o *GetAcrV1NameTagsReferenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAcrV1NameTagsReferenceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAcrV1NameTagsReferenceBadRequest creates a GetAcrV1NameTagsReferenceBadRequest with default headers values
func NewGetAcrV1NameTagsReferenceBadRequest() *GetAcrV1NameTagsReferenceBadRequest {
	return &GetAcrV1NameTagsReferenceBadRequest{}
}

/*GetAcrV1NameTagsReferenceBadRequest handles this case with default header values.

On failure
*/
type GetAcrV1NameTagsReferenceBadRequest struct {
	Payload *GetAcrV1NameTagsReferenceBadRequestBody
}

func (o *GetAcrV1NameTagsReferenceBadRequest) Error() string {
	return fmt.Sprintf("[GET /acr/v1/{name}/_tags/{reference}][%d] getAcrV1NameTagsReferenceBadRequest  %+v", 400, o.Payload)
}

func (o *GetAcrV1NameTagsReferenceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAcrV1NameTagsReferenceBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAcrV1NameTagsReferenceUnauthorized creates a GetAcrV1NameTagsReferenceUnauthorized with default headers values
func NewGetAcrV1NameTagsReferenceUnauthorized() *GetAcrV1NameTagsReferenceUnauthorized {
	return &GetAcrV1NameTagsReferenceUnauthorized{}
}

/*GetAcrV1NameTagsReferenceUnauthorized handles this case with default header values.

Unauthorized access
*/
type GetAcrV1NameTagsReferenceUnauthorized struct {
	Payload *GetAcrV1NameTagsReferenceUnauthorizedBody
}

func (o *GetAcrV1NameTagsReferenceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /acr/v1/{name}/_tags/{reference}][%d] getAcrV1NameTagsReferenceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAcrV1NameTagsReferenceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAcrV1NameTagsReferenceUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAcrV1NameTagsReferenceNotFound creates a GetAcrV1NameTagsReferenceNotFound with default headers values
func NewGetAcrV1NameTagsReferenceNotFound() *GetAcrV1NameTagsReferenceNotFound {
	return &GetAcrV1NameTagsReferenceNotFound{}
}

/*GetAcrV1NameTagsReferenceNotFound handles this case with default header values.

The repository or tag is unknown to the registry.
*/
type GetAcrV1NameTagsReferenceNotFound struct {
	Payload *GetAcrV1NameTagsReferenceNotFoundBody
}

func (o *GetAcrV1NameTagsReferenceNotFound) Error() string {
	return fmt.Sprintf("[GET /acr/v1/{name}/_tags/{reference}][%d] getAcrV1NameTagsReferenceNotFound  %+v", 404, o.Payload)
}

func (o *GetAcrV1NameTagsReferenceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAcrV1NameTagsReferenceNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAcrV1NameTagsReferenceBadRequestBody get acr v1 name tags reference bad request body
swagger:model GetAcrV1NameTagsReferenceBadRequestBody
*/
type GetAcrV1NameTagsReferenceBadRequestBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get acr v1 name tags reference bad request body
func (o *GetAcrV1NameTagsReferenceBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAcrV1NameTagsReferenceBadRequestBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAcrV1NameTagsReferenceBadRequest" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAcrV1NameTagsReferenceBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAcrV1NameTagsReferenceBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetAcrV1NameTagsReferenceBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAcrV1NameTagsReferenceNotFoundBody get acr v1 name tags reference not found body
swagger:model GetAcrV1NameTagsReferenceNotFoundBody
*/
type GetAcrV1NameTagsReferenceNotFoundBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get acr v1 name tags reference not found body
func (o *GetAcrV1NameTagsReferenceNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAcrV1NameTagsReferenceNotFoundBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAcrV1NameTagsReferenceNotFound" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAcrV1NameTagsReferenceNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAcrV1NameTagsReferenceNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetAcrV1NameTagsReferenceNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAcrV1NameTagsReferenceOKBody get acr v1 name tags reference o k body
swagger:model GetAcrV1NameTagsReferenceOKBody
*/
type GetAcrV1NameTagsReferenceOKBody struct {

	// data
	Data *models.TagAttributeList `json:"data,omitempty"`
}

// Validate validates this get acr v1 name tags reference o k body
func (o *GetAcrV1NameTagsReferenceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAcrV1NameTagsReferenceOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAcrV1NameTagsReferenceOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAcrV1NameTagsReferenceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAcrV1NameTagsReferenceOKBody) UnmarshalBinary(b []byte) error {
	var res GetAcrV1NameTagsReferenceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAcrV1NameTagsReferenceUnauthorizedBody get acr v1 name tags reference unauthorized body
swagger:model GetAcrV1NameTagsReferenceUnauthorizedBody
*/
type GetAcrV1NameTagsReferenceUnauthorizedBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get acr v1 name tags reference unauthorized body
func (o *GetAcrV1NameTagsReferenceUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAcrV1NameTagsReferenceUnauthorizedBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAcrV1NameTagsReferenceUnauthorized" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAcrV1NameTagsReferenceUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAcrV1NameTagsReferenceUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetAcrV1NameTagsReferenceUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
