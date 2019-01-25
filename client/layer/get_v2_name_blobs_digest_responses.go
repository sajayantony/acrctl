// Code generated by go-swagger; DO NOT EDIT.

package layer

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

// GetV2NameBlobsDigestReader is a Reader for the GetV2NameBlobsDigest structure.
type GetV2NameBlobsDigestReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetV2NameBlobsDigestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetV2NameBlobsDigestOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 206:
		result := NewGetV2NameBlobsDigestPartialContent(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 307:
		result := NewGetV2NameBlobsDigestTemporaryRedirect()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetV2NameBlobsDigestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetV2NameBlobsDigestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetV2NameBlobsDigestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 416:
		result := NewGetV2NameBlobsDigestRequestRangeNotSatisfiable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetV2NameBlobsDigestOK creates a GetV2NameBlobsDigestOK with default headers values
func NewGetV2NameBlobsDigestOK(writer io.Writer) *GetV2NameBlobsDigestOK {
	return &GetV2NameBlobsDigestOK{
		Payload: writer,
	}
}

/*GetV2NameBlobsDigestOK handles this case with default header values.

The blob identified by digest is available. The blob content will be present in the body of the response.
*/
type GetV2NameBlobsDigestOK struct {
	Payload io.Writer
}

func (o *GetV2NameBlobsDigestOK) Error() string {
	return fmt.Sprintf("[GET /v2/{name}/blobs/{digest}][%d] getV2NameBlobsDigestOK  %+v", 200, o.Payload)
}

func (o *GetV2NameBlobsDigestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV2NameBlobsDigestPartialContent creates a GetV2NameBlobsDigestPartialContent with default headers values
func NewGetV2NameBlobsDigestPartialContent(writer io.Writer) *GetV2NameBlobsDigestPartialContent {
	return &GetV2NameBlobsDigestPartialContent{
		Payload: writer,
	}
}

/*GetV2NameBlobsDigestPartialContent handles this case with default header values.

The blob identified by digest is available. The specified chunk of blob content will be present in the body of the request.
*/
type GetV2NameBlobsDigestPartialContent struct {
	Payload io.Writer
}

func (o *GetV2NameBlobsDigestPartialContent) Error() string {
	return fmt.Sprintf("[GET /v2/{name}/blobs/{digest}][%d] getV2NameBlobsDigestPartialContent  %+v", 206, o.Payload)
}

func (o *GetV2NameBlobsDigestPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV2NameBlobsDigestTemporaryRedirect creates a GetV2NameBlobsDigestTemporaryRedirect with default headers values
func NewGetV2NameBlobsDigestTemporaryRedirect() *GetV2NameBlobsDigestTemporaryRedirect {
	return &GetV2NameBlobsDigestTemporaryRedirect{}
}

/*GetV2NameBlobsDigestTemporaryRedirect handles this case with default header values.

The blob identified by digest is available at the provided location.
*/
type GetV2NameBlobsDigestTemporaryRedirect struct {
}

func (o *GetV2NameBlobsDigestTemporaryRedirect) Error() string {
	return fmt.Sprintf("[GET /v2/{name}/blobs/{digest}][%d] getV2NameBlobsDigestTemporaryRedirect ", 307)
}

func (o *GetV2NameBlobsDigestTemporaryRedirect) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetV2NameBlobsDigestBadRequest creates a GetV2NameBlobsDigestBadRequest with default headers values
func NewGetV2NameBlobsDigestBadRequest() *GetV2NameBlobsDigestBadRequest {
	return &GetV2NameBlobsDigestBadRequest{}
}

/*GetV2NameBlobsDigestBadRequest handles this case with default header values.

On failure
*/
type GetV2NameBlobsDigestBadRequest struct {
	Payload *GetV2NameBlobsDigestBadRequestBody
}

func (o *GetV2NameBlobsDigestBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/{name}/blobs/{digest}][%d] getV2NameBlobsDigestBadRequest  %+v", 400, o.Payload)
}

func (o *GetV2NameBlobsDigestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV2NameBlobsDigestBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV2NameBlobsDigestUnauthorized creates a GetV2NameBlobsDigestUnauthorized with default headers values
func NewGetV2NameBlobsDigestUnauthorized() *GetV2NameBlobsDigestUnauthorized {
	return &GetV2NameBlobsDigestUnauthorized{}
}

/*GetV2NameBlobsDigestUnauthorized handles this case with default header values.

Unauthorized access
*/
type GetV2NameBlobsDigestUnauthorized struct {
	Payload *GetV2NameBlobsDigestUnauthorizedBody
}

func (o *GetV2NameBlobsDigestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/{name}/blobs/{digest}][%d] getV2NameBlobsDigestUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV2NameBlobsDigestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV2NameBlobsDigestUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV2NameBlobsDigestNotFound creates a GetV2NameBlobsDigestNotFound with default headers values
func NewGetV2NameBlobsDigestNotFound() *GetV2NameBlobsDigestNotFound {
	return &GetV2NameBlobsDigestNotFound{}
}

/*GetV2NameBlobsDigestNotFound handles this case with default header values.

The blob, identified by name and digest, is unknown to the registry.
*/
type GetV2NameBlobsDigestNotFound struct {
	Payload *GetV2NameBlobsDigestNotFoundBody
}

func (o *GetV2NameBlobsDigestNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/{name}/blobs/{digest}][%d] getV2NameBlobsDigestNotFound  %+v", 404, o.Payload)
}

func (o *GetV2NameBlobsDigestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV2NameBlobsDigestNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV2NameBlobsDigestRequestRangeNotSatisfiable creates a GetV2NameBlobsDigestRequestRangeNotSatisfiable with default headers values
func NewGetV2NameBlobsDigestRequestRangeNotSatisfiable() *GetV2NameBlobsDigestRequestRangeNotSatisfiable {
	return &GetV2NameBlobsDigestRequestRangeNotSatisfiable{}
}

/*GetV2NameBlobsDigestRequestRangeNotSatisfiable handles this case with default header values.

The range specification cannot be satisfied for the requested content.
*/
type GetV2NameBlobsDigestRequestRangeNotSatisfiable struct {
}

func (o *GetV2NameBlobsDigestRequestRangeNotSatisfiable) Error() string {
	return fmt.Sprintf("[GET /v2/{name}/blobs/{digest}][%d] getV2NameBlobsDigestRequestRangeNotSatisfiable ", 416)
}

func (o *GetV2NameBlobsDigestRequestRangeNotSatisfiable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetV2NameBlobsDigestBadRequestBody get v2 name blobs digest bad request body
swagger:model GetV2NameBlobsDigestBadRequestBody
*/
type GetV2NameBlobsDigestBadRequestBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get v2 name blobs digest bad request body
func (o *GetV2NameBlobsDigestBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetV2NameBlobsDigestBadRequestBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getV2NameBlobsDigestBadRequest" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV2NameBlobsDigestBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV2NameBlobsDigestBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetV2NameBlobsDigestBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetV2NameBlobsDigestNotFoundBody get v2 name blobs digest not found body
swagger:model GetV2NameBlobsDigestNotFoundBody
*/
type GetV2NameBlobsDigestNotFoundBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get v2 name blobs digest not found body
func (o *GetV2NameBlobsDigestNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetV2NameBlobsDigestNotFoundBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getV2NameBlobsDigestNotFound" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV2NameBlobsDigestNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV2NameBlobsDigestNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetV2NameBlobsDigestNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetV2NameBlobsDigestUnauthorizedBody get v2 name blobs digest unauthorized body
swagger:model GetV2NameBlobsDigestUnauthorizedBody
*/
type GetV2NameBlobsDigestUnauthorizedBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this get v2 name blobs digest unauthorized body
func (o *GetV2NameBlobsDigestUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetV2NameBlobsDigestUnauthorizedBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getV2NameBlobsDigestUnauthorized" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV2NameBlobsDigestUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV2NameBlobsDigestUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetV2NameBlobsDigestUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}