package bacotell

import "github.com/spf13/viper"

const (
	ConfigBotName   = "bot_name"
	ConfigBotToken  = "bot_token"
	ConfigPluginDir = "plugin_dir"
)

func InitConfig() {
	viper.SetDefault(ConfigBotName, "BacoTell")
	viper.SetDefault(ConfigBotToken, "")
	viper.SetDefault(ConfigPluginDir, "plugins")
}
