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

// NewDeleteAcrV1NameManifestsReferenceMetadataMetadataParams creates a new DeleteAcrV1NameManifestsReferenceMetadataMetadataParams object
// with the default values initialized.
func NewDeleteAcrV1NameManifestsReferenceMetadataMetadataParams() *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams {
	var ()
	return &DeleteAcrV1NameManifestsReferenceMetadataMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAcrV1NameManifestsReferenceMetadataMetadataParamsWithTimeout creates a new DeleteAcrV1NameManifestsReferenceMetadataMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAcrV1NameManifestsReferenceMetadataMetadataParamsWithTimeout(timeout time.Duration) *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams {
	var ()
	return &DeleteAcrV1NameManifestsReferenceMetadataMetadataParams{

		timeout: timeout,
	}
}

// NewDeleteAcrV1NameManifestsReferenceMetadataMetadataParamsWithContext creates a new DeleteAcrV1NameManifestsReferenceMetadataMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAcrV1NameManifestsReferenceMetadataMetadataParamsWithContext(ctx context.Context) *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams {
	var ()
	return &DeleteAcrV1NameManifestsReferenceMetadataMetadataParams{

		Context: ctx,
	}
}

// NewDeleteAcrV1NameManifestsReferenceMetadataMetadataParamsWithHTTPClient creates a new DeleteAcrV1NameManifestsReferenceMetadataMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAcrV1NameManifestsReferenceMetadataMetadataParamsWithHTTPClient(client *http.Client) *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams {
	var ()
	return &DeleteAcrV1NameManifestsReferenceMetadataMetadataParams{
		HTTPClient: client,
	}
}

/*DeleteAcrV1NameManifestsReferenceMetadataMetadataParams contains all the parameters to send to the API endpoint
for the delete acr v1 name manifests reference metadata metadata operation typically these are written to a http.Request
*/
type DeleteAcrV1NameManifestsReferenceMetadataMetadataParams struct {

	/*Metadata
	  Name of the metadata

	*/
	Metadata string
	/*Name
	  Name of the image (including the namespace)

	*/
	Name string
	/*Reference
	  A tag or a digest, pointing to a specific image

	*/
	Reference string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete acr v1 name manifests reference metadata metadata params
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams) WithTimeout(timeout time.Duration) *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete acr v1 name manifests reference metadata metadata params
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete acr v1 name manifests reference metadata metadata params
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams) WithContext(ctx context.Context) *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete acr v1 name manifests reference metadata metadata params
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete acr v1 name manifests reference metadata metadata params
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams) WithHTTPClient(client *http.Client) *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete acr v1 name manifests reference metadata metadata params
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetadata adds the metadata to the delete acr v1 name manifests reference metadata metadata params
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams) WithMetadata(metadata string) *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams {
	o.SetMetadata(metadata)
	return o
}

// SetMetadata adds the metadata to the delete acr v1 name manifests reference metadata metadata params
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams) SetMetadata(metadata string) {
	o.Metadata = metadata
}

// WithName adds the name to the delete acr v1 name manifests reference metadata metadata params
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams) WithName(name string) *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete acr v1 name manifests reference metadata metadata params
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams) SetName(name string) {
	o.Name = name
}

// WithReference adds the reference to the delete acr v1 name manifests reference metadata metadata params
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams) WithReference(reference string) *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams {
	o.SetReference(reference)
	return o
}

// SetReference adds the reference to the delete acr v1 name manifests reference metadata metadata params
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams) SetReference(reference string) {
	o.Reference = reference
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAcrV1NameManifestsReferenceMetadataMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param metadata
	if err := r.SetPathParam("metadata", o.Metadata); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param reference
	if err := r.SetPathParam("reference", o.Reference); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
