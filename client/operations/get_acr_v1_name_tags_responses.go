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

// GetAcrV1NameTagsReader is a Reader for the GetAcrV1NameTags structure.
type GetAcrV1NameTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAcrV1NameTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAcrV1NameTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAcrV1NameTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAcrV1NameTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAcrV1NameTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAcrV1NameTagsOK creates a GetAcrV1NameTagsOK with default headers values
func NewGetAcrV1NameTagsOK() *GetAcrV1NameTagsOK {
	return &GetAcrV1NameTagsOK{}
}

/*GetAcrV1NameTagsOK handles this case with default header values.

Returns a list of tags
*/
type GetAcrV1NameTagsOK struct {
	Payload *GetAcrV1NameTagsOKBody
}

func (o *GetAcrV1NameTagsOK) Error() string {
	return fmt.Sprintf("[GET /acr/v1/{name}/_tags][%d] getAcrV1NameTagsOK  %+v", 200, o.Payload)
}

func (o *GetAcrV1NameTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAcrV1NameTagsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAcrV1NameTagsBadRequest creates a GetAcrV1NameTagsBadRequest with default headers values
func NewGetAcrV1NameTagsBadRequest() *GetAcrV1NameTagsBadRequest {
	return &GetAcrV1NameTagsBadRequest{}
}

/*GetAcrV1NameTagsBadRequest handles this case with default header values.

On failure
*/
type GetAcrV1NameTagsBadRequest struct {
	Payload *GetAcrV1NameTagsBadRequestBody
}

func (o *GetAcrV1NameTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /acr/v1/{name}/_tags][%d] getAcrV1NameTagsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAcrV1NameTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAcrV1NameTagsBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAcrV1NameTagsUnauthorized creates a GetAcrV1NameTagsUnauthorized with default headers values
func NewGetAcrV1NameTagsUnauthorized() *GetAcrV1NameTagsUnauthorized {
	return &GetAcrV1NameTagsUnauthorized{}
}

/*GetAcrV1NameTagsUnauthorized handles this case with default header values.

Unauthorized access
*/
type GetAcrV1NameTagsUnauthorized struct {
	Payload *GetAcrV1NameTagsUnauthorizedBody
}

func (o *GetAcrV1NameTagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /acr/v1/{name}/_tags][%d] getAcrV1NameTagsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAcrV1NameTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAcrV1NameTagsUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAcrV1NameTagsNotFound creates a GetAcrV1NameTagsNotFound with default headers values
func NewGetAcrV1NameTagsNotFound() *GetAcrV1NameTagsNotFound {
	return &GetAcrV1NameTagsNotFound{}
}

/*GetAcrV1NameTagsNotFound handles this case with default header values.

The repository is unknown to the registry.
*/
type GetAcrV1NameTagsNotFound struct {
	Payload *GetAcrV1NameTagsNotFoundBody
}

func (o *GetAcrV1NameTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /acr/v1/{name}/_tags][%d] getAcrV1NameTagsNotFound  %+v", 404, o.Payload)
}

func (o *GetAcrV1NameTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAcrV1NameTagsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAcrV1NameTagsBadRequestBody get acr v1 name tags bad request body
swagger:model GetAcrV1NameTagsBadRequestBody
*/
type GetAcrV1NameTagsBadRequestBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get acr v1 name tags bad request body
func (o *GetAcrV1NameTagsBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAcrV1NameTagsBadRequestBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAcrV1NameTagsBadRequest" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAcrV1NameTagsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAcrV1NameTagsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetAcrV1NameTagsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAcrV1NameTagsNotFoundBody get acr v1 name tags not found body
swagger:model GetAcrV1NameTagsNotFoundBody
*/
type GetAcrV1NameTagsNotFoundBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get acr v1 name tags not found body
func (o *GetAcrV1NameTagsNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAcrV1NameTagsNotFoundBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAcrV1NameTagsNotFound" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAcrV1NameTagsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAcrV1NameTagsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetAcrV1NameTagsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAcrV1NameTagsOKBody get acr v1 name tags o k body
swagger:model GetAcrV1NameTagsOKBody
*/
type GetAcrV1NameTagsOKBody struct {

	// data
	Data *models.TagAttributeList `json:"data,omitempty"`
}

// Validate validates this get acr v1 name tags o k body
func (o *GetAcrV1NameTagsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAcrV1NameTagsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAcrV1NameTagsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAcrV1NameTagsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAcrV1NameTagsOKBody) UnmarshalBinary(b []byte) error {
	var res GetAcrV1NameTagsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAcrV1NameTagsUnauthorizedBody get acr v1 name tags unauthorized body
swagger:model GetAcrV1NameTagsUnauthorizedBody
*/
type GetAcrV1NameTagsUnauthorizedBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get acr v1 name tags unauthorized body
func (o *GetAcrV1NameTagsUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAcrV1NameTagsUnauthorizedBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAcrV1NameTagsUnauthorized" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAcrV1NameTagsUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAcrV1NameTagsUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetAcrV1NameTagsUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
