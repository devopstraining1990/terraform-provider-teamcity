// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cvbarros/go-teamcity-sdk/client/agent"
	"github.com/cvbarros/go-teamcity-sdk/client/agent_pool"
	"github.com/cvbarros/go-teamcity-sdk/client/build"
	"github.com/cvbarros/go-teamcity-sdk/client/build_queue"
	"github.com/cvbarros/go-teamcity-sdk/client/build_type"
	"github.com/cvbarros/go-teamcity-sdk/client/change"
	"github.com/cvbarros/go-teamcity-sdk/client/client"
	"github.com/cvbarros/go-teamcity-sdk/client/debug"
	"github.com/cvbarros/go-teamcity-sdk/client/federation"
	"github.com/cvbarros/go-teamcity-sdk/client/group"
	"github.com/cvbarros/go-teamcity-sdk/client/investigation"
	"github.com/cvbarros/go-teamcity-sdk/client/mute"
	"github.com/cvbarros/go-teamcity-sdk/client/problem"
	"github.com/cvbarros/go-teamcity-sdk/client/project"
	"github.com/cvbarros/go-teamcity-sdk/client/server"
	"github.com/cvbarros/go-teamcity-sdk/client/user"
	"github.com/cvbarros/go-teamcity-sdk/client/vcs_root"
	"github.com/cvbarros/go-teamcity-sdk/client/vcs_root_instance"
)

// Default team city r e s t HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "teamcity"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new team city r e s t HTTP client.
func NewHTTPClient(formats strfmt.Registry) *TeamCityREST {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new team city r e s t HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *TeamCityREST {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new team city r e s t client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *TeamCityREST {
	cli := new(TeamCityREST)
	cli.Transport = transport

	cli.Agent = agent.New(transport, formats)

	cli.AgentPool = agent_pool.New(transport, formats)

	cli.Build = build.New(transport, formats)

	cli.BuildQueue = build_queue.New(transport, formats)

	cli.BuildType = build_type.New(transport, formats)

	cli.Change = change.New(transport, formats)

	cli.Client = client.New(transport, formats)

	cli.Debug = debug.New(transport, formats)

	cli.Federation = federation.New(transport, formats)

	cli.Group = group.New(transport, formats)

	cli.Investigation = investigation.New(transport, formats)

	cli.Mute = mute.New(transport, formats)

	cli.Problem = problem.New(transport, formats)

	cli.Project = project.New(transport, formats)

	cli.Server = server.New(transport, formats)

	cli.User = user.New(transport, formats)

	cli.VcsRoot = vcs_root.New(transport, formats)

	cli.VcsRootInstance = vcs_root_instance.New(transport, formats)

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

// TeamCityREST is a client for team city r e s t
type TeamCityREST struct {
	Agent *agent.Client

	AgentPool *agent_pool.Client

	Build *build.Client

	BuildQueue *build_queue.Client

	BuildType *build_type.Client

	Change *change.Client

	Client *client.Client

	Debug *debug.Client

	Federation *federation.Client

	Group *group.Client

	Investigation *investigation.Client

	Mute *mute.Client

	Problem *problem.Client

	Project *project.Client

	Server *server.Client

	User *user.Client

	VcsRoot *vcs_root.Client

	VcsRootInstance *vcs_root_instance.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *TeamCityREST) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Agent.SetTransport(transport)

	c.AgentPool.SetTransport(transport)

	c.Build.SetTransport(transport)

	c.BuildQueue.SetTransport(transport)

	c.BuildType.SetTransport(transport)

	c.Change.SetTransport(transport)

	c.Client.SetTransport(transport)

	c.Debug.SetTransport(transport)

	c.Federation.SetTransport(transport)

	c.Group.SetTransport(transport)

	c.Investigation.SetTransport(transport)

	c.Mute.SetTransport(transport)

	c.Problem.SetTransport(transport)

	c.Project.SetTransport(transport)

	c.Server.SetTransport(transport)

	c.User.SetTransport(transport)

	c.VcsRoot.SetTransport(transport)

	c.VcsRootInstance.SetTransport(transport)

}