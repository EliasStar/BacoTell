package bacotell

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/EliasStar/BacoTell/internal/codec"
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/hashicorp/go-plugin"
	"github.com/spf13/viper"
)

const bacotellPlugin = "bacotell_plugin"

var (
	clients = make(map[string]*plugin.Client)

	commands   = make(map[string]common.Command)
	components = make(map[string]common.Component)
	modals     = make(map[string]common.Modal)
)

// HandshakeConfig returns the handshake configuration for the plugin.
func HandshakeConfig() plugin.HandshakeConfig {
	return plugin.HandshakeConfig{
		ProtocolVersion:  ProtocolVersion,
		MagicCookieKey:   "BacoTell",
		MagicCookieValue: "d950af91-326f-4151-bbf0-2484ebd6a4fa",
	}
}

// PluginMap returns the plugin map for the given plugin implementation (may be nil).
func PluginMap(pluginImpl common.Plugin) plugin.PluginSet {
	return plugin.PluginSet{bacotellPlugin: codec.NewPlugin(pluginImpl)}
}

// loadAll loads all the plugins found in the plugin directory.
func loadAll() {
	pluginDir := viper.GetString(ConfigPluginDir)
	absPluginDir, err := filepath.Abs(pluginDir)
	if err != nil {
		loaderLogger.Error("cannot determine absolute path of plugin directory", ConfigPluginDir, pluginDir, "err", err)
		return
	}

	if err := os.MkdirAll(absPluginDir, os.ModePerm); err != nil {
		loaderLogger.Error("cannot create plugin directory", ConfigPluginDir, pluginDir, "err", err)
		return
	}

	err = filepath.Walk(absPluginDir, func(path string, info os.FileInfo, err error) error {
		pathLogger := loaderLogger.With("path", path)

		if err != nil {
			pathLogger.Warn("failed to walk path", "err", err)
			return err
		}

		if info.IsDir() || !strings.HasSuffix(strings.ToLower(info.Name()), ".btp") {
			pathLogger.Debug("skipping non-plugin path")
			return nil
		}

		pathLogger.Debug("loading plugin")
		_load(plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig:  HandshakeConfig(),
			Plugins:          PluginMap(nil),
			AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
			Logger:           loaderLogger,
			Cmd:              exec.Command(path),
		}))

		return nil
	})

	if err != nil {
		loaderLogger.Error("could not load plugins", "path", absPluginDir, "err", err)
	}
}

// loadSingle loads a single plugin using the provided reattach configuration.
func loadSingle(reattachConfig *plugin.ReattachConfig) {
	loaderLogger.Debug("loading plugin")
	_load(plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  HandshakeConfig(),
		Plugins:          PluginMap(nil),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		Logger:           loaderLogger,
		Reattach:         reattachConfig,
	}))
}

// unloadAll unloads all the loaded plugins.
func unloadAll() {
	for _, client := range clients {
		client.Kill()
	}

	clients = nil
}

// _load loads all commands, components, etc. from the provided plugin client.
func _load(client *plugin.Client) {
	protocol, err := client.Client()
	if err != nil {
		loaderLogger.Warn("could not connect to plugin process", "error", err)
		client.Kill()
		return
	}

	raw, err := protocol.Dispense(bacotellPlugin)
	if err != nil {
		loaderLogger.Warn("could not dispense plugin", "error", err)
		client.Kill()
		return
	}

	pluginImpl, ok := raw.(common.Plugin)
	if !ok {
		loaderLogger.Warn("unexpected plugin type", "raw", raw)
		client.Kill()
		return
	}

	id, err := pluginImpl.ID()
	if err != nil {
		loaderLogger.Warn("could not get plugin id", "err", err)
		client.Kill()
		return
	}

	if _, ok := clients[id]; ok {
		loaderLogger.Warn("plugin with same id is already loaded", "id", id)
		client.Kill()
		return
	}

	clients[id] = client
	pluginLogger := loaderLogger.With("plugin", id)

	pluginLogger.Info("loading commands")
	cmds, err := pluginImpl.ApplicationCommands()
	if err != nil {
		pluginLogger.Warn("could not get commands", "err", err)
	} else if len(cmds) == 0 {
		pluginLogger.Info("plugin has no commands")
	} else {
		for i, cmd := range cmds {
			data, err := cmd.Data()
			if err != nil {
				pluginLogger.Warn("cannot get command data", "command", i, "err", err)
				continue
			}

			if !strings.HasPrefix(data.Name, id) {
				pluginLogger.Warn("command name must be prefixed with the plugin id", "command", data.Name)
				continue
			}

			commands[data.Name] = cmd
		}
	}

	pluginLogger.Info("loading components")
	cpts, err := pluginImpl.MessageComponents()
	if err != nil {
		pluginLogger.Warn("could not get components", "err", err)
	} else if len(cpts) == 0 {
		pluginLogger.Info("plugin has no components")
	} else {
		for i, cpt := range cpts {
			cid, err := cpt.CustomID()
			if err != nil {
				pluginLogger.Warn("cannot get custom id", "component", i, "err", err)
				continue
			}

			if !strings.HasPrefix(cid, id) {
				pluginLogger.Warn("custom id must be prefixed with the plugin id", "component", cid)
				continue
			}

			components[cid] = cpt
		}
	}

	pluginLogger.Info("loading modals")
	mods, err := pluginImpl.Modals()
	if err != nil {
		pluginLogger.Warn("could not get modals", "err", err)
	} else if len(mods) == 0 {
		pluginLogger.Info("plugin has no modals")
	} else {
		for i, mod := range mods {
			cid, err := mod.CustomID()
			if err != nil {
				pluginLogger.Warn("cannot get custom id", "modal", i, "err", err)
				continue
			}

			if !strings.HasPrefix(cid, id) {
				pluginLogger.Warn("custom id must be prefixed with the plugin id", "modal", cid)
				continue
			}

			modals[cid] = mod
		}
	}
}
