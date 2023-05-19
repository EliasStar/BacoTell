package bacotell

import (
	"os"

	"github.com/hashicorp/go-hclog"
)

var logger = hclog.New(&hclog.LoggerOptions{
	Name:   "bacotell",
	Output: os.Stdout,
	Level:  hclog.Debug,
})

func GetLogger() hclog.Logger {
	return logger
}
