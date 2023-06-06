// Package bacotell_plugin provides the plugin entrypoint.
package bacotell_plugin

import (
	"context"
	"os"

	"github.com/EliasStar/BacoTell/internal/bacotell"
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

// The default implementation for bacotell_common.Plugin which simply returns pre-set values.
type defaultPlugin struct {
	id         string
	commands   []common.Command
	components []common.Component
	modals     []common.Modal
}

// defaultPlugin implements bacotell_common.Plugin.
var _ common.Plugin = defaultPlugin{}

// ID implements bacotell_common.Plugin.
func (p defaultPlugin) ID() (string, error) {
	return p.id, nil
}

// ApplicationCommands implements bacotell_common.Plugin.
func (p defaultPlugin) ApplicationCommands() ([]common.Command, error) {
	return p.commands, nil
}

// MessageComponents implements bacotell_common.Plugin.
func (p defaultPlugin) MessageComponents() ([]common.Component, error) {
	return p.components, nil
}

// Modals implements bacotell_common.Plugin.
func (p defaultPlugin) Modals() ([]common.Modal, error) {
	return p.modals, nil
}

// Singleton instance of defaultPlugin.
var internalPlugin defaultPlugin

// SetApplicationCommands sets the commands provided by the current plugin.
func SetApplicationCommands(commands ...common.Command) {
	internalPlugin.commands = commands
}

// SetMessageComponents sets the components provided by the current plugin.
func SetMessageComponents(components ...common.Component) {
	internalPlugin.components = components
}

// SetModals sets the modals provided by the current plugin.
func SetModals(modals ...common.Modal) {
	internalPlugin.modals = modals
}

// Run starts the plugin with the provided id and a default plugin implementation.
//
// Before calling this all commands, component, etc. that should be provided by the plugin
// must be set using the functions SetApplicationCommands, SetMessageComponents, etc.
// The returned logger should be used by the plugin code to log messages to stdout/stderr.
// Because this function does not wait for the plugin to terminate again,
// the returned channel should be used to block until it gets closed, signaling a exit.
func Run(id string) (hclog.Logger, <-chan struct{}, error) {
	internalPlugin.id = id
	return RunCustom(internalPlugin)
}

// Debug starts the plugin in debug mode with the provided id, bot token and a default plugin implementation.
//
// In debug mode a BacoTell instance is started within the plugin process so that the plugin can be executed directly.
// This allows for easy debugging within an IDE. The embedded BacoTell only loads the current plugin.
//
// Before calling this all commands, component, etc. that should be provided by the plugin
// must be set using the functions SetApplicationCommands, SetMessageComponents, etc.
// The returned logger should be used by the plugin code to log messages to stdout/stderr.
// Because this function does not wait for the plugin to terminate again,
// the returned channel should be used to block until it gets closed, signaling a exit.
func Debug(id, token string) (hclog.Logger, <-chan struct{}, error) {
	internalPlugin.id = id
	return DebugCustom(internalPlugin, token)
}

// RunCustom starts the plugin with the provided id and custom plugin implementation.
//
// The returned logger should be used by the plugin code to log messages to stdout/stderr.
// Because this function does not wait for the plugin to terminate again,
// the returned channel should be used to block until it gets closed, signaling a exit.
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

// DebugCustom starts the plugin in debug mode with the provided id, bot token and custom plugin implementation.
//
// In debug mode a BacoTell instance is started within the plugin process so that the plugin can be executed directly.
// This allows for easy debugging within an IDE. The embedded BacoTell only loads the current plugin.
//
// The returned logger should be used by the plugin code to log messages to stdout/stderr.
// Because this function does not wait for the plugin to terminate again,
// the returned channel should be used to block until it gets closed, signaling a exit.
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
