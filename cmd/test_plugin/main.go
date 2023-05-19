package main

import (
	"os"

	"github.com/EliasStar/BacoTell/pkg/bacotell"
	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/hashicorp/go-hclog"
)

var logger = hclog.New(&hclog.LoggerOptions{
	Name:   "test_plugin",
	Output: os.Stdout,
	Level:  hclog.Debug,
})

var commands = []provider.Command{
	TestCommand{},
}

var components = []provider.Component{
	TestComponent{},
}

func main() {
	bacotell.SetInteractionProvider(provider.NewInteractionProvider("test_plugin", commands, components))
	bacotell.DebugPlugin(logger, "MTA4OTk3MjA5MTE0MDQ0NDIwMA.GD1PSP.3i4VfMmnHuPjenyFgUUxkpDoJqw_zC_pW1sMsQ")
}
