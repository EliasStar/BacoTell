package bacotell

import (
	"io"

	"github.com/hashicorp/go-hclog"
)

var (
	logger        hclog.Logger
	loaderLogger  hclog.Logger
	discordLogger hclog.Logger
)

// initLoggers initializes all (sub)loggers for the major subsystems
func initLoggers(output io.Writer, level hclog.Level) {
	logger = hclog.New(&hclog.LoggerOptions{
		Name:   "bacotell",
		Output: output,
		Level:  level,
	})

	loaderLogger = logger.Named("loader")
	discordLogger = logger.Named("discord")
}
