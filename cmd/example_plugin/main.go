package main

import (
	plugin "github.com/EliasStar/BacoTell/pkg/bacotell_plugin"
	"github.com/hashicorp/go-hclog"
)

const pluginID = "example"

var (
	logger    hclog.Logger
	closeChan <-chan struct{}
)

func main() {
	plugin.SetApplicationCommands(ExampleCommand{})
	plugin.SetMessageComponents(ExampleComponent{})
	plugin.SetModals(ExampleModal{})

	logger, closeChan, _ = plugin.Run(pluginID)
	//logger, closeChan, _ = plugin.Debug(pluginID, "bot_token")

	<-closeChan
}
