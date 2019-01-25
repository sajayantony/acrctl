// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sajayantony/acrctl/client/acr_repository"
	"github.com/sajayantony/acrctl/client/layer"
	"github.com/sajayantony/acrctl/client/manifest"
	"github.com/sajayantony/acrctl/client/operations"
	"github.com/sajayantony/acrctl/client/repository"
	"github.com/sajayantony/acrctl/client/tag"
	"github.com/sajayantony/acrctl/client/v2"
)

// Default azure container registry HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "yuwatest12-msft.azurecr-test.io"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new azure container registry HTTP client.
func NewHTTPClient(formats strfmt.Registry) *AzureContainerRegistry {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new azure container registry HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *AzureContainerRegistry {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new azure container registry client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *AzureContainerRegistry {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(AzureContainerRegistry)
	cli.Transport = transport

	cli.AcrRepository = acr_repository.New(transport, formats)

	cli.Layer = layer.New(transport, formats)

	cli.Manifest = manifest.New(transport, formats)

	cli.Operations = operations.New(transport, formats)

	cli.Repository = repository.New(transport, formats)

	cli.Tag = tag.New(transport, formats)

	cli.V2 = v2.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// AzureContainerRegistry is a client for azure container registry
type AzureContainerRegistry struct {
	AcrRepository *acr_repository.Client

	Layer *layer.Client

	Manifest *manifest.Client

	Operations *operations.Client

	Repository *repository.Client

	Tag *tag.Client

	V2 *v2.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *AzureContainerRegistry) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.AcrRepository.SetTransport(transport)

	c.Layer.SetTransport(transport)

	c.Manifest.SetTransport(transport)

	c.Operations.SetTransport(transport)

	c.Repository.SetTransport(transport)

	c.Tag.SetTransport(transport)

	c.V2.SetTransport(transport)

}
