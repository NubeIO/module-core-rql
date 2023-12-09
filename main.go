package main

import (
	"github.com/NubeIO/lib-module-go/shared"
	"github.com/NubeIO/module-core-rql/pkg"
	"github.com/hashicorp/go-plugin"
)

func ServePlugin() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.HandshakeConfig,
		Plugins:         plugin.PluginSet{"module-core-rql": &shared.NubeModule{Impl: &pkg.Module{}}},
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}

func main() {
	ServePlugin()
}
