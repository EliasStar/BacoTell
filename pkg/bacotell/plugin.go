package bacotell

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/bacotell"
	"github.com/EliasStar/BacoTell/internal/discord"
	"github.com/EliasStar/BacoTell/internal/loader"
	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/spf13/viper"
)

func SetInteractionProvider(provider provider.InteractionProvider) {
	loader.InteractionProvider = provider
}

func RunPlugin(logger hclog.Logger) {
	plugin.Serve(&plugin.ServeConfig{
		GRPCServer:      plugin.DefaultGRPCServer,
		HandshakeConfig: loader.HandshakeConfig(),
		Plugins:         loader.PluginMap(),
		Logger:          logger,
	})
}

func DebugPlugin(logger hclog.Logger, token string) {
	closeChan := make(chan struct{})
	reattachConfigChan := make(chan *plugin.ReattachConfig)

	ctx, cancel := context.WithCancel(context.Background())

	go plugin.Serve(&plugin.ServeConfig{
		GRPCServer:      plugin.DefaultGRPCServer,
		HandshakeConfig: loader.HandshakeConfig(),
		Plugins:         loader.PluginMap(),
		Logger:          logger,

		Test: &plugin.ServeTestConfig{
			Context:          ctx,
			CloseCh:          closeChan,
			ReattachConfigCh: reattachConfigChan,
		},
	})

	bacotell.InitConfig()
	viper.Set(bacotell.ConfigBotToken, token)

	loader.LoadFromRunning(<-reattachConfigChan)
	discord.Connect()
	loader.Unload()

	cancel()
	<-closeChan
}
