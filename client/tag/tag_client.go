// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sajayantony/acrctl/client"
)

// New creates a new tag API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tag API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetV2NameTagsList Fetch the tags under the repository identified by 'name'
*/
func (a *Client) GetV2NameTagsList(params *GetV2NameTagsListParams, authInfo runtime.ClientAuthInfoWriter) (*GetV2NameTagsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV2NameTagsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetV2NameTagsList",
		Method:             "GET",
		PathPattern:        "/v2/{name}/tags/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV2NameTagsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetV2NameTagsListOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
