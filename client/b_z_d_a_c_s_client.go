// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/hardstylez72/bzdacs/client/group"
	"github.com/hardstylez72/bzdacs/client/namespace"
	"github.com/hardstylez72/bzdacs/client/route"
	"github.com/hardstylez72/bzdacs/client/system"
	"github.com/hardstylez72/bzdacs/client/tag"
)

// Default b z d a c s HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost:8081"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new b z d a c s HTTP client.
func NewHTTPClient(formats strfmt.Registry) *BZDACS {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new b z d a c s HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *BZDACS {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new b z d a c s client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *BZDACS {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(BZDACS)
	cli.Transport = transport
	cli.Group = group.New(transport, formats)
	cli.Namespace = namespace.New(transport, formats)
	cli.Route = route.New(transport, formats)
	cli.System = system.New(transport, formats)
	cli.Tag = tag.New(transport, formats)
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

// BZDACS is a client for b z d a c s
type BZDACS struct {
	Group group.ClientService

	Namespace namespace.ClientService

	Route route.ClientService

	System system.ClientService

	Tag tag.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *BZDACS) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Group.SetTransport(transport)
	c.Namespace.SetTransport(transport)
	c.Route.SetTransport(transport)
	c.System.SetTransport(transport)
	c.Tag.SetTransport(transport)
}
