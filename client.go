package sdk

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// PluginManifest is a test
type PluginManifest struct {
	ID, Name, Author, Version string
}

// StarportPlugin is the interface that we're exposing as a plugin.
type StarportPlugin interface {
	GetManifest() PluginManifest
}

// StarportPluginRPC is what the server is using to communicate to the plugin over RPC
type StarportPluginRPC struct {
	client *rpc.Client
}

func (i *StarportPluginRPC) GetManifest() PluginManifest {
	rep := PluginManifest{}
	err := i.client.Call("Plugin.GetManifest", new(interface{}), &rep)
	if err != nil {
		panic(err)
	}
	return rep
}

// This is the implementation of hashicorp plugin stuff
type StartPortPluginSystem struct {
	Impl StarportPlugin
}

func (p *StartPortPluginSystem) Server(*plugin.MuxBroker) (interface{}, error) {
	return &StartPortPluginSystem{Impl: p.Impl}, nil
}

func (StartPortPluginSystem) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &StarportPluginRPC{client: c}, nil
}
