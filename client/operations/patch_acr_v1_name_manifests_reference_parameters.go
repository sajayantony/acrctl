// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewPatchAcrV1NameManifestsReferenceParams creates a new PatchAcrV1NameManifestsReferenceParams object
// with the default values initialized.
func NewPatchAcrV1NameManifestsReferenceParams() *PatchAcrV1NameManifestsReferenceParams {
	var ()
	return &PatchAcrV1NameManifestsReferenceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAcrV1NameManifestsReferenceParamsWithTimeout creates a new PatchAcrV1NameManifestsReferenceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAcrV1NameManifestsReferenceParamsWithTimeout(timeout time.Duration) *PatchAcrV1NameManifestsReferenceParams {
	var ()
	return &PatchAcrV1NameManifestsReferenceParams{

		timeout: timeout,
	}
}

// NewPatchAcrV1NameManifestsReferenceParamsWithContext creates a new PatchAcrV1NameManifestsReferenceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAcrV1NameManifestsReferenceParamsWithContext(ctx context.Context) *PatchAcrV1NameManifestsReferenceParams {
	var ()
	return &PatchAcrV1NameManifestsReferenceParams{

		Context: ctx,
	}
}

// NewPatchAcrV1NameManifestsReferenceParamsWithHTTPClient creates a new PatchAcrV1NameManifestsReferenceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAcrV1NameManifestsReferenceParamsWithHTTPClient(client *http.Client) *PatchAcrV1NameManifestsReferenceParams {
	var ()
	return &PatchAcrV1NameManifestsReferenceParams{
		HTTPClient: client,
	}
}

/*PatchAcrV1NameManifestsReferenceParams contains all the parameters to send to the API endpoint
for the patch acr v1 name manifests reference operation typically these are written to a http.Request
*/
type PatchAcrV1NameManifestsReferenceParams struct {

	/*Name
	  Name of the image (including the namespace)

	*/
	Name string
	/*Reference
	  A tag or a digest, pointing to a specific image

	*/
	Reference string
	/*Value*/
	Value string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch acr v1 name manifests reference params
func (o *PatchAcrV1NameManifestsReferenceParams) WithTimeout(timeout time.Duration) *PatchAcrV1NameManifestsReferenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch acr v1 name manifests reference params
func (o *PatchAcrV1NameManifestsReferenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch acr v1 name manifests reference params
func (o *PatchAcrV1NameManifestsReferenceParams) WithContext(ctx context.Context) *PatchAcrV1NameManifestsReferenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch acr v1 name manifests reference params
func (o *PatchAcrV1NameManifestsReferenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch acr v1 name manifests reference params
func (o *PatchAcrV1NameManifestsReferenceParams) WithHTTPClient(client *http.Client) *PatchAcrV1NameManifestsReferenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch acr v1 name manifests reference params
func (o *PatchAcrV1NameManifestsReferenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the patch acr v1 name manifests reference params
func (o *PatchAcrV1NameManifestsReferenceParams) WithName(name string) *PatchAcrV1NameManifestsReferenceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch acr v1 name manifests reference params
func (o *PatchAcrV1NameManifestsReferenceParams) SetName(name string) {
	o.Name = name
}

// WithReference adds the reference to the patch acr v1 name manifests reference params
func (o *PatchAcrV1NameManifestsReferenceParams) WithReference(reference string) *PatchAcrV1NameManifestsReferenceParams {
	o.SetReference(reference)
	return o
}

// SetReference adds the reference to the patch acr v1 name manifests reference params
func (o *PatchAcrV1NameManifestsReferenceParams) SetReference(reference string) {
	o.Reference = reference
}

// WithValue adds the value to the patch acr v1 name manifests reference params
func (o *PatchAcrV1NameManifestsReferenceParams) WithValue(value string) *PatchAcrV1NameManifestsReferenceParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the patch acr v1 name manifests reference params
func (o *PatchAcrV1NameManifestsReferenceParams) SetValue(value string) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAcrV1NameManifestsReferenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param reference
	if err := r.SetPathParam("reference", o.Reference); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Value); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
