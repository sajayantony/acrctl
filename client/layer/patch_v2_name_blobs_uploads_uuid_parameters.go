// Code generated by go-swagger; DO NOT EDIT.

package layer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPatchV2NameBlobsUploadsUUIDParams creates a new PatchV2NameBlobsUploadsUUIDParams object
// with the default values initialized.
func NewPatchV2NameBlobsUploadsUUIDParams() *PatchV2NameBlobsUploadsUUIDParams {
	var ()
	return &PatchV2NameBlobsUploadsUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV2NameBlobsUploadsUUIDParamsWithTimeout creates a new PatchV2NameBlobsUploadsUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchV2NameBlobsUploadsUUIDParamsWithTimeout(timeout time.Duration) *PatchV2NameBlobsUploadsUUIDParams {
	var ()
	return &PatchV2NameBlobsUploadsUUIDParams{

		timeout: timeout,
	}
}

// NewPatchV2NameBlobsUploadsUUIDParamsWithContext creates a new PatchV2NameBlobsUploadsUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchV2NameBlobsUploadsUUIDParamsWithContext(ctx context.Context) *PatchV2NameBlobsUploadsUUIDParams {
	var ()
	return &PatchV2NameBlobsUploadsUUIDParams{

		Context: ctx,
	}
}

// NewPatchV2NameBlobsUploadsUUIDParamsWithHTTPClient creates a new PatchV2NameBlobsUploadsUUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchV2NameBlobsUploadsUUIDParamsWithHTTPClient(client *http.Client) *PatchV2NameBlobsUploadsUUIDParams {
	var ()
	return &PatchV2NameBlobsUploadsUUIDParams{
		HTTPClient: client,
	}
}

/*PatchV2NameBlobsUploadsUUIDParams contains all the parameters to send to the API endpoint
for the patch v2 name blobs uploads UUID operation typically these are written to a http.Request
*/
type PatchV2NameBlobsUploadsUUIDParams struct {

	/*Name
	  Name of the image (including the namespace)

	*/
	Name string
	/*UUID
	  A uuid identifying the upload.

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch v2 name blobs uploads UUID params
func (o *PatchV2NameBlobsUploadsUUIDParams) WithTimeout(timeout time.Duration) *PatchV2NameBlobsUploadsUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v2 name blobs uploads UUID params
func (o *PatchV2NameBlobsUploadsUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v2 name blobs uploads UUID params
func (o *PatchV2NameBlobsUploadsUUIDParams) WithContext(ctx context.Context) *PatchV2NameBlobsUploadsUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v2 name blobs uploads UUID params
func (o *PatchV2NameBlobsUploadsUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v2 name blobs uploads UUID params
func (o *PatchV2NameBlobsUploadsUUIDParams) WithHTTPClient(client *http.Client) *PatchV2NameBlobsUploadsUUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v2 name blobs uploads UUID params
func (o *PatchV2NameBlobsUploadsUUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the patch v2 name blobs uploads UUID params
func (o *PatchV2NameBlobsUploadsUUIDParams) WithName(name string) *PatchV2NameBlobsUploadsUUIDParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch v2 name blobs uploads UUID params
func (o *PatchV2NameBlobsUploadsUUIDParams) SetName(name string) {
	o.Name = name
}

// WithUUID adds the uuid to the patch v2 name blobs uploads UUID params
func (o *PatchV2NameBlobsUploadsUUIDParams) WithUUID(uuid string) *PatchV2NameBlobsUploadsUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the patch v2 name blobs uploads UUID params
func (o *PatchV2NameBlobsUploadsUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV2NameBlobsUploadsUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
