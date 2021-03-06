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

// DeleteV2NameBlobsUploadsUUIDReader is a Reader for the DeleteV2NameBlobsUploadsUUID structure.
type DeleteV2NameBlobsUploadsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV2NameBlobsUploadsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteV2NameBlobsUploadsUUIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteV2NameBlobsUploadsUUIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteV2NameBlobsUploadsUUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteV2NameBlobsUploadsUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteV2NameBlobsUploadsUUIDNoContent creates a DeleteV2NameBlobsUploadsUUIDNoContent with default headers values
func NewDeleteV2NameBlobsUploadsUUIDNoContent() *DeleteV2NameBlobsUploadsUUIDNoContent {
	return &DeleteV2NameBlobsUploadsUUIDNoContent{}
}

/*DeleteV2NameBlobsUploadsUUIDNoContent handles this case with default header values.

The upload has been successfully deleted.
*/
type DeleteV2NameBlobsUploadsUUIDNoContent struct {
}

func (o *DeleteV2NameBlobsUploadsUUIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v2/{name}/blobs/uploads/{uuid}][%d] deleteV2NameBlobsUploadsUuidNoContent ", 204)
}

func (o *DeleteV2NameBlobsUploadsUUIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteV2NameBlobsUploadsUUIDBadRequest creates a DeleteV2NameBlobsUploadsUUIDBadRequest with default headers values
func NewDeleteV2NameBlobsUploadsUUIDBadRequest() *DeleteV2NameBlobsUploadsUUIDBadRequest {
	return &DeleteV2NameBlobsUploadsUUIDBadRequest{}
}

/*DeleteV2NameBlobsUploadsUUIDBadRequest handles this case with default header values.

On failure
*/
type DeleteV2NameBlobsUploadsUUIDBadRequest struct {
	Payload *DeleteV2NameBlobsUploadsUUIDBadRequestBody
}

func (o *DeleteV2NameBlobsUploadsUUIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v2/{name}/blobs/uploads/{uuid}][%d] deleteV2NameBlobsUploadsUuidBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteV2NameBlobsUploadsUUIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteV2NameBlobsUploadsUUIDBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteV2NameBlobsUploadsUUIDUnauthorized creates a DeleteV2NameBlobsUploadsUUIDUnauthorized with default headers values
func NewDeleteV2NameBlobsUploadsUUIDUnauthorized() *DeleteV2NameBlobsUploadsUUIDUnauthorized {
	return &DeleteV2NameBlobsUploadsUUIDUnauthorized{}
}

/*DeleteV2NameBlobsUploadsUUIDUnauthorized handles this case with default header values.

Unauthorized access
*/
type DeleteV2NameBlobsUploadsUUIDUnauthorized struct {
	Payload *DeleteV2NameBlobsUploadsUUIDUnauthorizedBody
}

func (o *DeleteV2NameBlobsUploadsUUIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v2/{name}/blobs/uploads/{uuid}][%d] deleteV2NameBlobsUploadsUuidUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteV2NameBlobsUploadsUUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteV2NameBlobsUploadsUUIDUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteV2NameBlobsUploadsUUIDNotFound creates a DeleteV2NameBlobsUploadsUUIDNotFound with default headers values
func NewDeleteV2NameBlobsUploadsUUIDNotFound() *DeleteV2NameBlobsUploadsUUIDNotFound {
	return &DeleteV2NameBlobsUploadsUUIDNotFound{}
}

/*DeleteV2NameBlobsUploadsUUIDNotFound handles this case with default header values.

The upload is unknown to the registry. The client may ignore this error and assume the upload has been deleted.
*/
type DeleteV2NameBlobsUploadsUUIDNotFound struct {
	Payload *DeleteV2NameBlobsUploadsUUIDNotFoundBody
}

func (o *DeleteV2NameBlobsUploadsUUIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v2/{name}/blobs/uploads/{uuid}][%d] deleteV2NameBlobsUploadsUuidNotFound  %+v", 404, o.Payload)
}

func (o *DeleteV2NameBlobsUploadsUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteV2NameBlobsUploadsUUIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteV2NameBlobsUploadsUUIDBadRequestBody delete v2 name blobs uploads UUID bad request body
swagger:model DeleteV2NameBlobsUploadsUUIDBadRequestBody
*/
type DeleteV2NameBlobsUploadsUUIDBadRequestBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this delete v2 name blobs uploads UUID bad request body
func (o *DeleteV2NameBlobsUploadsUUIDBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteV2NameBlobsUploadsUUIDBadRequestBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteV2NameBlobsUploadsUuidBadRequest" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteV2NameBlobsUploadsUUIDBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteV2NameBlobsUploadsUUIDBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DeleteV2NameBlobsUploadsUUIDBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteV2NameBlobsUploadsUUIDNotFoundBody delete v2 name blobs uploads UUID not found body
swagger:model DeleteV2NameBlobsUploadsUUIDNotFoundBody
*/
type DeleteV2NameBlobsUploadsUUIDNotFoundBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this delete v2 name blobs uploads UUID not found body
func (o *DeleteV2NameBlobsUploadsUUIDNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteV2NameBlobsUploadsUUIDNotFoundBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteV2NameBlobsUploadsUuidNotFound" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteV2NameBlobsUploadsUUIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteV2NameBlobsUploadsUUIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteV2NameBlobsUploadsUUIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteV2NameBlobsUploadsUUIDUnauthorizedBody delete v2 name blobs uploads UUID unauthorized body
swagger:model DeleteV2NameBlobsUploadsUUIDUnauthorizedBody
*/
type DeleteV2NameBlobsUploadsUUIDUnauthorizedBody struct {

	// data
	Data *models.Error `json:"data,omitempty"`
}

// Validate validates this delete v2 name blobs uploads UUID unauthorized body
func (o *DeleteV2NameBlobsUploadsUUIDUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteV2NameBlobsUploadsUUIDUnauthorizedBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteV2NameBlobsUploadsUuidUnauthorized" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteV2NameBlobsUploadsUUIDUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteV2NameBlobsUploadsUUIDUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res DeleteV2NameBlobsUploadsUUIDUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
