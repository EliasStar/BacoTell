// Package bacotell provides the BacoTell entrypoint and the core subsystems of the framework.
package bacotell

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/spf13/viper"
)

const (
	Version         = "v1.0.0"
	ProtocolVersion = 1 // Must be equal to major version.

	ConfigBotName   = "bot_name"
	ConfigBotToken  = "bot_token"
	ConfigPluginDir = "plugin_dir"
	ConfigLogFile   = "log_file"
	ConfigLogLevel  = "log_level"
)

// InitConfig sets default values for viper config entries.
func InitConfig() {
	viper.SetDefault(ConfigBotName, "BacoTell")
	viper.SetDefault(ConfigBotToken, "")
	viper.SetDefault(ConfigPluginDir, "plugins")
	viper.SetDefault(ConfigLogFile, "")
	viper.SetDefault(ConfigLogLevel, "info")
}

// Run starts BacoTell.
func Run() {
	logOutput, err := os.OpenFile(viper.GetString(ConfigLogFile), os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		logOutput = os.Stdout
	}

	logLevel := hclog.LevelFromString(viper.GetString(ConfigLogLevel))
	if logLevel == hclog.NoLevel {
		logLevel = hclog.Info
	}

	initLoggers(logOutput, logLevel)

	if viper.GetString(ConfigBotToken) == "" {
		logger.Error("no bot token provided, set '" + ConfigBotToken + "' in config")
		return
	}

	loadAll()
	connect()
	unloadAll()
}

// Debug starts BacoTell in debug mode.
func Debug(token string, reattachConfig *plugin.ReattachConfig) {
	initLoggers(os.Stdout, hclog.Trace)

	if token == "" {
		logger.Error("no bot token provided")
		return
	}

	InitConfig()
	viper.Set(ConfigBotToken, token)

	loadSingle(reattachConfig)
	connect()
	unloadAll()
}
