package bacotell

import (
	"os"

	"github.com/hashicorp/go-hclog"
)

var (
	logger        hclog.Logger
	loaderLogger  hclog.Logger
	discordLogger hclog.Logger
)

func initLoggers(level hclog.Level) {
	logger = hclog.New(&hclog.LoggerOptions{
		Name:   "bacotell",
		Output: os.Stdout,
		Level:  level,
	})

	loaderLogger = logger.Named("loader")
	discordLogger = logger.Named("discord")
}
