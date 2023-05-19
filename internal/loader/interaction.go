package loader

import (
	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/hashicorp/go-plugin"
)

var (
	commands   = make(map[string][]provider.Command)
	components = make(map[string][]provider.Component)
)

func GetApplicationCommands() map[string][]provider.Command {
	return commands
}

func GetMessageComponents() map[string][]provider.Component {
	return components
}

func loadInteractionProvider(protocol plugin.ClientProtocol) {
	logger.Info("loading interaction provider")

	raw, err := protocol.Dispense(interactionProviderPlugin)
	if err != nil {
		logger.Info("plugin has none")
		return
	}

	ip, ok := raw.(provider.InteractionProvider)
	if !ok {
		logger.Warn("unexpected type", "raw", raw)
		return
	}

	prefix, err := ip.GetPrefix()
	if err != nil {
		logger.Warn("could not get prefix", "err", err)
		return
	}

	cmds, err := ip.GetApplicationCommands()
	if err != nil {
		logger.Warn("could not get commands", "err", err)
		return
	}

	cpts, err := ip.GetMessageComponents()
	if err != nil {
		logger.Warn("could not get components", "err", err)
		return
	}

	commands[prefix] = cmds
	components[prefix] = cpts
}
