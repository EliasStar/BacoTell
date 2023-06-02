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
	modals     []common.Modal
}

var _ common.Plugin = defaultPlugin{}

func (p defaultPlugin) ID() (string, error) {
	return p.id, nil
}

func (p defaultPlugin) ApplicationCommands() ([]common.Command, error) {
	return p.commands, nil
}

func (p defaultPlugin) MessageComponents() ([]common.Component, error) {
	return p.components, nil
}

func (p defaultPlugin) Modals() ([]common.Modal, error) {
	return p.modals, nil
}

var internalPlugin defaultPlugin

func SetApplicationCommands(commands ...common.Command) {
	internalPlugin.commands = commands
}

func SetMessageComponents(components ...common.Component) {
	internalPlugin.components = components
}

func SetModals(modals ...common.Modal) {
	internalPlugin.modals = modals
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
