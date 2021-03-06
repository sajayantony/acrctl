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

// PutAcrV1NameTagsReferenceMetadataMetadataReader is a Reader for the PutAcrV1NameTagsReferenceMetadataMetadata structure.
type PutAcrV1NameTagsReferenceMetadataMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAcrV1NameTagsReferenceMetadataMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPutAcrV1NameTagsReferenceMetadataMetadataCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutAcrV1NameTagsReferenceMetadataMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPutAcrV1NameTagsReferenceMetadataMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutAcrV1NameTagsReferenceMetadataMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutAcrV1NameTagsReferenceMetadataMetadataCreated creates a PutAcrV1NameTagsReferenceMetadataMetadataCreated with default headers values
func NewPutAcrV1NameTagsReferenceMetadataMetadataCreated() *PutAcrV1NameTagsReferenceMetadataMetadataCreated {
	return &PutAcrV1NameTagsReferenceMetadataMetadataCreated{}
}

/*PutAcrV1NameTagsReferenceMetadataMetadataCreated handles this case with default header values.

The metadata is added/updated
*/
type PutAcrV1NameTagsReferenceMetadataMetadataCreated struct {
}

func (o *PutAcrV1NameTagsReferenceMetadataMetadataCreated) Error() string {
	return fmt.Sprintf("[PUT /acr/v1/{name}/_tags/{reference}/_metadata/{metadata}][%d] putAcrV1NameTagsReferenceMetadataMetadataCreated ", 201)
}

func (o *PutAcrV1NameTagsReferenceMetadataMetadataCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAcrV1NameTagsReferenceMetadataMetadataBadRequest creates a PutAcrV1NameTagsReferenceMetadataMetadataBadRequest with default headers values
func NewPutAcrV1NameTagsReferenceMetadataMetadataBadRequest() *PutAcrV1NameTagsReferenceMetadataMetadataBadRequest {
	return &PutAcrV1NameTagsReferenceMetadataMetadataBadRequest{}
}

/*PutAcrV1NameTagsReferenceMetadataMetadataBadRequest handles this case with default header values.

On failure
*/
type PutAcrV1NameTagsReferenceMetadataMetadataBadRequest struct {
	Payload *PutAcrV1NameTagsReferenceMetadataMetadataBadRequestBody
}

func (o *PutAcrV1NameTagsReferenceMetadataMetadataBadRequest) Error() string {
	return fmt.Sprintf("[PUT /acr/v1/{name}/_tags/{reference}/_metadata/{metadata}][%d] putAcrV1NameTagsReferenceMetadataMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *PutAcrV1NameTagsReferenceMetadataMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutAcrV1NameTagsReferenceMetadataMetadataBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAcrV1NameTagsReferenceMetadataMetadataUnauthorized creates a PutAcrV1NameTagsReferenceMetadataMetadataUnauthorized with default headers values
func NewPutAcrV1NameTagsReferenceMetadataMetadataUnauthorized() *PutAcrV1NameTagsReferenceMetadataMetadataUnauthorized {
	return &PutAcrV1NameTagsReferenceMetadataMetadataUnauthorized{}
}

/*PutAcrV1NameTagsReferenceMetadataMetadataUnauthorized handles this case with default header values.

Unauthorized access
*/
type PutAcrV1NameTagsReferenceMetadataMetadataUnauthorized struct {
	Payload *PutAcrV1NameTagsReferenceMetadataMetadataUnauthorizedBody
}

func (o *PutAcrV1NameTagsReferenceMetadataMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /acr/v1/{name}/_tags/{reference}/_metadata/{metadata}][%d] putAcrV1NameTagsReferenceMetadataMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *PutAcrV1NameTagsReferenceMetadataMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutAcrV1NameTagsReferenceMetadataMetadataUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAcrV1NameTagsReferenceMetadataMetadataNotFound creates a PutAcrV1NameTagsReferenceMetadataMetadataNotFound with default headers values
func NewPutAcrV1NameTagsReferenceMetadataMetadataNotFound() *PutAcrV1NameTagsReferenceMetadataMetadataNotFound {
	return &PutAcrV1NameTagsReferenceMetadataMetadataNotFound{}
}

/*PutAcrV1NameTagsReferenceMetadataMetadataNotFound handles this case with default header values.

The repository, tag or metadata is unknown
*/
type PutAcrV1NameTagsReferenceMetadataMetadataNotFound struct {
	Payload *PutAcrV1NameTagsReferenceMetadataMetadataNotFoundBody
}

func (o *PutAcrV1NameTagsReferenceMetadataMetadataNotFound) Error() string {
	return fmt.Sprintf("[PUT /acr/v1/{name}/_tags/{reference}/_metadata/{metadata}][%d] putAcrV1NameTagsReferenceMetadataMetadataNotFound  %+v", 404, o.Payload)
}

func (o *PutAcrV1NameTagsReferenceMetadataMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutAcrV1NameTagsReferenceMetadataMetadataNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutAcrV1NameTagsReferenceMetadataMetadataBadRequestBody put acr v1 name tags reference metadata metadata bad request body
swagger:model PutAcrV1NameTagsReferenceMetadataMetadataBadRequestBody
*/
type PutAcrV1NameTagsReferenceMetadataMetadataBadRequestBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this put acr v1 name tags reference metadata metadata bad request body
func (o *PutAcrV1NameTagsReferenceMetadataMetadataBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutAcrV1NameTagsReferenceMetadataMetadataBadRequestBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putAcrV1NameTagsReferenceMetadataMetadataBadRequest" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutAcrV1NameTagsReferenceMetadataMetadataBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutAcrV1NameTagsReferenceMetadataMetadataBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PutAcrV1NameTagsReferenceMetadataMetadataBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PutAcrV1NameTagsReferenceMetadataMetadataNotFoundBody put acr v1 name tags reference metadata metadata not found body
swagger:model PutAcrV1NameTagsReferenceMetadataMetadataNotFoundBody
*/
type PutAcrV1NameTagsReferenceMetadataMetadataNotFoundBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this put acr v1 name tags reference metadata metadata not found body
func (o *PutAcrV1NameTagsReferenceMetadataMetadataNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutAcrV1NameTagsReferenceMetadataMetadataNotFoundBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putAcrV1NameTagsReferenceMetadataMetadataNotFound" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutAcrV1NameTagsReferenceMetadataMetadataNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutAcrV1NameTagsReferenceMetadataMetadataNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PutAcrV1NameTagsReferenceMetadataMetadataNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PutAcrV1NameTagsReferenceMetadataMetadataUnauthorizedBody put acr v1 name tags reference metadata metadata unauthorized body
swagger:model PutAcrV1NameTagsReferenceMetadataMetadataUnauthorizedBody
*/
type PutAcrV1NameTagsReferenceMetadataMetadataUnauthorizedBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this put acr v1 name tags reference metadata metadata unauthorized body
func (o *PutAcrV1NameTagsReferenceMetadataMetadataUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutAcrV1NameTagsReferenceMetadataMetadataUnauthorizedBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putAcrV1NameTagsReferenceMetadataMetadataUnauthorized" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutAcrV1NameTagsReferenceMetadataMetadataUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutAcrV1NameTagsReferenceMetadataMetadataUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res PutAcrV1NameTagsReferenceMetadataMetadataUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
