package loader

import (
	"github.com/EliasStar/BacoTell/internal/codec"
	"github.com/EliasStar/BacoTell/internal/common"
	"github.com/EliasStar/BacoTell/pkg/bacotell"
	"github.com/hashicorp/go-plugin"
)

const bacotellPlugin = "bacotell_plugin"

var clients []*plugin.Client

func HandshakeConfig() plugin.HandshakeConfig {
	return plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "BacoTell",
		MagicCookieValue: "d950af91-326f-4151-bbf0-2484ebd6a4fa",
	}
}

func PluginMap(pluginImpl bacotell.Plugin) plugin.PluginSet {
	return plugin.PluginSet{bacotellPlugin: codec.NewBacoTellPlugin(pluginImpl)}
}

func LoadFromFolder() {
	// TODO
}

func LoadFromRunning(reattachConfig *plugin.ReattachConfig) {
	load(plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  HandshakeConfig(),
		Plugins:          PluginMap(nil),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		Logger:           common.GetLogger(),
		Reattach:         reattachConfig,
	}))
}

func Unload() {
	for _, client := range clients {
		client.Kill()
	}
}

func load(client *plugin.Client) {
	logger := common.GetLogger().Named("loader")

	protocol, err := client.Client()
	if err != nil {
		logger.Warn("could not connect to plugin process", "error", err)
		client.Kill()
		return
	}

	raw, err := protocol.Dispense(bacotellPlugin)
	if err != nil {
		logger.Warn("could not dispense plugin", "error", err)
		client.Kill()
		return
	}

	pluginImpl, ok := raw.(bacotell.Plugin)
	if !ok {
		logger.Warn("unexpected plugin type", "raw", raw)
		client.Kill()
		return
	}

	id, err := pluginImpl.ID()
	if err != nil {
		logger.Warn("could not get plugin id", "err", err)
		client.Kill()
		return
	}

	clients = append(clients, client)
	pluginLogger := logger.With("plugin", id)

	loadApplicationCommands(id, pluginImpl, pluginLogger)
	loadMessageComponents(id, pluginImpl, pluginLogger)
}
