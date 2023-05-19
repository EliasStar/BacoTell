package loader

import (
	"github.com/EliasStar/BacoTell/internal/bacotell"
	"github.com/EliasStar/BacoTell/internal/plugin/provider/interaction"
	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/hashicorp/go-plugin"
)

const (
	interactionProviderPlugin = "interaction_provider"
)

var (
	InteractionProvider provider.InteractionProvider

	clients []*plugin.Client

	logger          = bacotell.GetLogger().Named("loader")
	providerLoaders = []func(plugin.ClientProtocol){
		loadInteractionProvider,
	}
)

func HandshakeConfig() plugin.HandshakeConfig {
	return plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "BacoTell",
		MagicCookieValue: "d950af91-326f-4151-bbf0-2484ebd6a4fa",
	}
}

func PluginMap() plugin.PluginSet {
	return plugin.PluginSet{
		interactionProviderPlugin: interaction.NewProviderPlugin(InteractionProvider),
	}
}

func LoadFromFolder() {
	// TODO
}

func LoadFromRunning(reattachConfig *plugin.ReattachConfig) {
	load(plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  HandshakeConfig(),
		Plugins:          PluginMap(),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		Logger:           bacotell.GetLogger(),
		Reattach:         reattachConfig,
	}))
}

func Unload() {
	for _, client := range clients {
		client.Kill()
	}
}

func load(client *plugin.Client) {
	protocol, err := client.Client()
	if err != nil {
		logger.Warn("could not connect to plugin", "error", err)
		client.Kill()
		return
	}

	clients = append(clients, client)

	for _, loadProvider := range providerLoaders {
		loadProvider(protocol)
	}
}
