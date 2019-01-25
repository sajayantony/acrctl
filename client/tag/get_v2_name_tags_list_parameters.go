// Code generated by go-swagger; DO NOT EDIT.

package tag

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

// NewGetV2NameTagsListParams creates a new GetV2NameTagsListParams object
// with the default values initialized.
func NewGetV2NameTagsListParams() *GetV2NameTagsListParams {
	var ()
	return &GetV2NameTagsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV2NameTagsListParamsWithTimeout creates a new GetV2NameTagsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV2NameTagsListParamsWithTimeout(timeout time.Duration) *GetV2NameTagsListParams {
	var ()
	return &GetV2NameTagsListParams{

		timeout: timeout,
	}
}

// NewGetV2NameTagsListParamsWithContext creates a new GetV2NameTagsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV2NameTagsListParamsWithContext(ctx context.Context) *GetV2NameTagsListParams {
	var ()
	return &GetV2NameTagsListParams{

		Context: ctx,
	}
}

// NewGetV2NameTagsListParamsWithHTTPClient creates a new GetV2NameTagsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV2NameTagsListParamsWithHTTPClient(client *http.Client) *GetV2NameTagsListParams {
	var ()
	return &GetV2NameTagsListParams{
		HTTPClient: client,
	}
}

/*GetV2NameTagsListParams contains all the parameters to send to the API endpoint
for the get v2 name tags list operation typically these are written to a http.Request
*/
type GetV2NameTagsListParams struct {

	/*Name
	  Name of the image (including the namespace)

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v2 name tags list params
func (o *GetV2NameTagsListParams) WithTimeout(timeout time.Duration) *GetV2NameTagsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v2 name tags list params
func (o *GetV2NameTagsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v2 name tags list params
func (o *GetV2NameTagsListParams) WithContext(ctx context.Context) *GetV2NameTagsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v2 name tags list params
func (o *GetV2NameTagsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v2 name tags list params
func (o *GetV2NameTagsListParams) WithHTTPClient(client *http.Client) *GetV2NameTagsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v2 name tags list params
func (o *GetV2NameTagsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get v2 name tags list params
func (o *GetV2NameTagsListParams) WithName(name string) *GetV2NameTagsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get v2 name tags list params
func (o *GetV2NameTagsListParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetV2NameTagsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
