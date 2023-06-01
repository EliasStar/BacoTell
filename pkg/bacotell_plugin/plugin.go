package bacotell_plugin

import (
	"context"
	"os"

	"github.com/EliasStar/BacoTell/internal/bacotell"
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type defaultPlugin struct {
	id         string
	commands   []common.Command
	components []common.Component
}

var _ common.Plugin = defaultPlugin{}

func (i defaultPlugin) ID() (string, error) {
	return i.id, nil
}

func (i defaultPlugin) ApplicationCommands() ([]common.Command, error) {
	return i.commands, nil
}

func (i defaultPlugin) MessageComponents() ([]common.Component, error) {
	return i.components, nil
}

var internalPlugin defaultPlugin

func SetApplicationCommands(commands []common.Command) {
	internalPlugin.commands = commands
}

func SetMessageComponents(components []common.Component) {
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

func RunCustom(customPlugin common.Plugin) (hclog.Logger, <-chan struct{}, error) {
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
			HandshakeConfig: bacotell.HandshakeConfig(),
			Plugins:         bacotell.PluginMap(customPlugin),
			Logger:          logger,
		})

		close(closeChan)
	}()

	return logger, closeChan, nil
}

func DebugCustom(customPlugin common.Plugin, token string) (hclog.Logger, <-chan struct{}, error) {
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
		Level:  hclog.Debug,
	})

	go plugin.Serve(&plugin.ServeConfig{
		GRPCServer:      plugin.DefaultGRPCServer,
		HandshakeConfig: bacotell.HandshakeConfig(),
		Plugins:         bacotell.PluginMap(customPlugin),
		Logger:          logger,

		Test: &plugin.ServeTestConfig{
			Context:          ctx,
			CloseCh:          closeChan,
			ReattachConfigCh: reattachConfigChan,
		},
	})

	go func() {
		bacotell.Debug(token, <-reattachConfigChan)
		cancel()
	}()

	return logger, closeChan, nil
}
