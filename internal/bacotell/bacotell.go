package bacotell

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/spf13/viper"
)

const (
	Version = "v0.1.0"

	ConfigBotName   = "bot_name"
	ConfigBotToken  = "bot_token"
	ConfigPluginDir = "plugin_dir"
)

func InitConfig() {
	viper.SetDefault(ConfigBotName, "BacoTell")
	viper.SetDefault(ConfigBotToken, "")
	viper.SetDefault(ConfigPluginDir, "plugins")
}

func Run() {
	initLoggers(hclog.Info)

	if viper.GetString(ConfigBotToken) == "" {
		logger.Error("no bot token provided, set '" + ConfigBotToken + "' in config")
		return
	}

	loadAll()
	connect()
	unloadAll()
}

func Debug(token string, reattachConfig *plugin.ReattachConfig) {
	initLoggers(hclog.Debug)

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
