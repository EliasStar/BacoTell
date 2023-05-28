package bacotell_plugin

import (
	"context"
	"os"

	"github.com/EliasStar/BacoTell/internal/common"
	"github.com/EliasStar/BacoTell/internal/discord"
	"github.com/EliasStar/BacoTell/internal/loader"
	"github.com/EliasStar/BacoTell/pkg/bacotell"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/spf13/viper"
)

type defaultPlugin struct {
	id         string
	commands   []bacotell.Command
	components []bacotell.Component
}

var _ bacotell.Plugin = defaultPlugin{}

func (i defaultPlugin) ID() (string, error) {
	return i.id, nil
}

func (i defaultPlugin) ApplicationCommands() ([]bacotell.Command, error) {
	return i.commands, nil
}

func (i defaultPlugin) MessageComponents() ([]bacotell.Component, error) {
	return i.components, nil
}

var internalPlugin defaultPlugin

func SetApplicationCommands(commands []bacotell.Command) {
	internalPlugin.commands = commands
}

func SetMessageComponents(components []bacotell.Component) {
	internalPlugin.components = components
}

func Run(id string) (hclog.Logger, <-chan struct{}, error) {
	internalPlugin.id = id
	return RunCustom(internalPlugin)
}

func Debug(id, token string) (hclog.Logger, <-chan struct{}, error) {
	internalPlugin.id = id
	return DebugCustom(internalPlugin, token)
}

func RunCustom(customPlugin bacotell.Plugin) (hclog.Logger, <-chan struct{}, error) {
	id, err := customPlugin.ID()
	if err != nil {
		return nil, nil, err
	}

	closeChan := make(chan struct{})
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   id,
		Output: os.Stdout,
		Level:  hclog.Info,
	})

	go func() {
		plugin.Serve(&plugin.ServeConfig{
			GRPCServer:      plugin.DefaultGRPCServer,
			HandshakeConfig: loader.HandshakeConfig(),
			Plugins:         loader.PluginMap(customPlugin),
			Logger:          logger,
		})

		close(closeChan)
	}()

	return logger, closeChan, nil
}

func DebugCustom(customPlugin bacotell.Plugin, token string) (hclog.Logger, <-chan struct{}, error) {
	id, err := customPlugin.ID()
	if err != nil {
		return nil, nil, err
	}

	closeChan := make(chan struct{})
	reattachConfigChan := make(chan *plugin.ReattachConfig)
	ctx, cancel := context.WithCancel(context.Background())
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   id,
		Output: os.Stdout,
		Level:  hclog.Info,
	})

	go plugin.Serve(&plugin.ServeConfig{
		GRPCServer:      plugin.DefaultGRPCServer,
		HandshakeConfig: loader.HandshakeConfig(),
		Plugins:         loader.PluginMap(customPlugin),
		Logger:          logger,

		Test: &plugin.ServeTestConfig{
			Context:          ctx,
			CloseCh:          closeChan,
			ReattachConfigCh: reattachConfigChan,
		},
	})

	go func() {
		common.InitConfig()
		viper.Set(common.ConfigBotToken, token)

		loader.LoadFromRunning(<-reattachConfigChan)
		discord.Connect()
		loader.Unload()

		cancel()
	}()

	return logger, closeChan, nil
}
