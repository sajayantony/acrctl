// Code generated by go-swagger; DO NOT EDIT.

package tag

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

// GetV2NameTagsListReader is a Reader for the GetV2NameTagsList structure.
type GetV2NameTagsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV2NameTagsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetV2NameTagsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetV2NameTagsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetV2NameTagsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetV2NameTagsListOK creates a GetV2NameTagsListOK with default headers values
func NewGetV2NameTagsListOK() *GetV2NameTagsListOK {
	return &GetV2NameTagsListOK{}
}

/*GetV2NameTagsListOK handles this case with default header values.

Gives a list of tags for the names repository.
*/
type GetV2NameTagsListOK struct {
	Payload *GetV2NameTagsListOKBody
}

func (o *GetV2NameTagsListOK) Error() string {
	return fmt.Sprintf("[GET /v2/{name}/tags/list][%d] getV2NameTagsListOK  %+v", 200, o.Payload)
}

func (o *GetV2NameTagsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV2NameTagsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV2NameTagsListUnauthorized creates a GetV2NameTagsListUnauthorized with default headers values
func NewGetV2NameTagsListUnauthorized() *GetV2NameTagsListUnauthorized {
	return &GetV2NameTagsListUnauthorized{}
}

/*GetV2NameTagsListUnauthorized handles this case with default header values.

Unauthorized access
*/
type GetV2NameTagsListUnauthorized struct {
	Payload *GetV2NameTagsListUnauthorizedBody
}

func (o *GetV2NameTagsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/{name}/tags/list][%d] getV2NameTagsListUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV2NameTagsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV2NameTagsListUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV2NameTagsListNotFound creates a GetV2NameTagsListNotFound with default headers values
func NewGetV2NameTagsListNotFound() *GetV2NameTagsListNotFound {
	return &GetV2NameTagsListNotFound{}
}

/*GetV2NameTagsListNotFound handles this case with default header values.

The named manifest could not be found in the Registry
*/
type GetV2NameTagsListNotFound struct {
	Payload *GetV2NameTagsListNotFoundBody
}

func (o *GetV2NameTagsListNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/{name}/tags/list][%d] getV2NameTagsListNotFound  %+v", 404, o.Payload)
}

func (o *GetV2NameTagsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV2NameTagsListNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetV2NameTagsListNotFoundBody get v2 name tags list not found body
swagger:model GetV2NameTagsListNotFoundBody
*/
type GetV2NameTagsListNotFoundBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get v2 name tags list not found body
func (o *GetV2NameTagsListNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetV2NameTagsListNotFoundBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getV2NameTagsListNotFound" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV2NameTagsListNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV2NameTagsListNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetV2NameTagsListNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetV2NameTagsListOKBody get v2 name tags list o k body
swagger:model GetV2NameTagsListOKBody
*/
type GetV2NameTagsListOKBody struct {

	// name
	Name string `json:"name,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this get v2 name tags list o k body
func (o *GetV2NameTagsListOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetV2NameTagsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV2NameTagsListOKBody) UnmarshalBinary(b []byte) error {
	var res GetV2NameTagsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetV2NameTagsListUnauthorizedBody get v2 name tags list unauthorized body
swagger:model GetV2NameTagsListUnauthorizedBody
*/
type GetV2NameTagsListUnauthorizedBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get v2 name tags list unauthorized body
func (o *GetV2NameTagsListUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetV2NameTagsListUnauthorizedBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getV2NameTagsListUnauthorized" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV2NameTagsListUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV2NameTagsListUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetV2NameTagsListUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
