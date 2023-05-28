package loader

import (
	"github.com/EliasStar/BacoTell/pkg/bacotell"
	"github.com/hashicorp/go-hclog"
)

var messageComponents = make(map[string][]bacotell.Component)

func MessageComponents() map[string][]bacotell.Component {
	return messageComponents
}

func loadMessageComponents(id string, plugin bacotell.Plugin, logger hclog.Logger) {
	logger.Info("loading components")

	components, err := plugin.MessageComponents()
	if err != nil {
		logger.Warn("could not get components", "err", err)
		return
	}

	if len(components) == 0 {
		logger.Info("plugin has no components")
		return
	}

	messageComponents[id] = components
}
