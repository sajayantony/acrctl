// Code generated by go-swagger; DO NOT EDIT.

package layer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sajayantony/acrctl/client"
)

// New creates a new layer API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for layer API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteV2NameBlobsUploadsUUID Cancel outstanding upload processes, releasing associated resources. If this is not called, the unfinished uploads will eventually timeout.
*/
func (a *Client) DeleteV2NameBlobsUploadsUUID(params *DeleteV2NameBlobsUploadsUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteV2NameBlobsUploadsUUIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV2NameBlobsUploadsUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteV2NameBlobsUploadsUUID",
		Method:             "DELETE",
		PathPattern:        "/v2/{name}/blobs/uploads/{uuid}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV2NameBlobsUploadsUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteV2NameBlobsUploadsUUIDNoContent), nil

}

/*
GetV2NameBlobsDigest Retrieve the blob from the registry identified by digest.
*/
func (a *Client) GetV2NameBlobsDigest(params *GetV2NameBlobsDigestParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*GetV2NameBlobsDigestOK, *GetV2NameBlobsDigestPartialContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV2NameBlobsDigestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetV2NameBlobsDigest",
		Method:             "GET",
		PathPattern:        "/v2/{name}/blobs/{digest}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV2NameBlobsDigestReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetV2NameBlobsDigestOK:
		return value, nil, nil
	case *GetV2NameBlobsDigestPartialContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetV2NameBlobsUploadsUUID Retrieve status of upload identified by uuid. The primary purpose of this endpoint is to resolve the current status of a resumable upload.
*/
func (a *Client) GetV2NameBlobsUploadsUUID(params *GetV2NameBlobsUploadsUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetV2NameBlobsUploadsUUIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV2NameBlobsUploadsUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetV2NameBlobsUploadsUUID",
		Method:             "GET",
		PathPattern:        "/v2/{name}/blobs/uploads/{uuid}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV2NameBlobsUploadsUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetV2NameBlobsUploadsUUIDNoContent), nil

}

/*
HeadV2NameBlobsDigest Same as GET, except only the headers are returned.
*/
func (a *Client) HeadV2NameBlobsDigest(params *HeadV2NameBlobsDigestParams, authInfo runtime.ClientAuthInfoWriter) (*HeadV2NameBlobsDigestOK, *HeadV2NameBlobsDigestPartialContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeadV2NameBlobsDigestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "HeadV2NameBlobsDigest",
		Method:             "HEAD",
		PathPattern:        "/v2/{name}/blobs/{digest}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HeadV2NameBlobsDigestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *HeadV2NameBlobsDigestOK:
		return value, nil, nil
	case *HeadV2NameBlobsDigestPartialContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
PatchV2NameBlobsUploadsUUID Upload a stream of data to upload without completing the upload.
*/
func (a *Client) PatchV2NameBlobsUploadsUUID(params *PatchV2NameBlobsUploadsUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchV2NameBlobsUploadsUUIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV2NameBlobsUploadsUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchV2NameBlobsUploadsUUID",
		Method:             "PATCH",
		PathPattern:        "/v2/{name}/blobs/uploads/{uuid}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV2NameBlobsUploadsUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchV2NameBlobsUploadsUUIDNoContent), nil

}

/*
PostV2NameBlobsUploads Upload a blob identified by the digest parameter in single request. This upload will not be resumable unless a recoverable error is returned.
*/
func (a *Client) PostV2NameBlobsUploads(params *PostV2NameBlobsUploadsParams, authInfo runtime.ClientAuthInfoWriter) (*PostV2NameBlobsUploadsCreated, *PostV2NameBlobsUploadsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV2NameBlobsUploadsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostV2NameBlobsUploads",
		Method:             "POST",
		PathPattern:        "/v2/{name}/blobs/uploads",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV2NameBlobsUploadsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostV2NameBlobsUploadsCreated:
		return value, nil, nil
	case *PostV2NameBlobsUploadsAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
PutV2NameBlobsUploadsUUID Complete the upload, providing all the data in the body, if necessary. A request without a body will just complete the upload with previously uploaded content.
*/
func (a *Client) PutV2NameBlobsUploadsUUID(params *PutV2NameBlobsUploadsUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutV2NameBlobsUploadsUUIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutV2NameBlobsUploadsUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutV2NameBlobsUploadsUUID",
		Method:             "PUT",
		PathPattern:        "/v2/{name}/blobs/uploads/{uuid}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutV2NameBlobsUploadsUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutV2NameBlobsUploadsUUIDNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
