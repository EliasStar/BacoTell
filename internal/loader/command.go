package loader

import (
	"github.com/EliasStar/BacoTell/pkg/bacotell"
	"github.com/hashicorp/go-hclog"
)

var applicationCommands = make(map[string][]bacotell.Command)

func ApplicationCommands() map[string][]bacotell.Command {
	return applicationCommands
}

func loadApplicationCommands(id string, plugin bacotell.Plugin, logger hclog.Logger) {
	logger.Info("loading commands")

	commands, err := plugin.ApplicationCommands()
	if err != nil {
		logger.Warn("could not get commands", "err", err)
		return
	}

	if len(commands) == 0 {
		logger.Info("plugin has no commands")
		return
	}

	applicationCommands[id] = commands
}
