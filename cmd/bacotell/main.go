package main

import (
	"os"

	"github.com/EliasStar/BacoTell/internal/common"
	"github.com/EliasStar/BacoTell/internal/discord"
	"github.com/EliasStar/BacoTell/internal/loader"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configPath string

	cmd = &cobra.Command{
		Run:     run,
		Version: common.Version,

		Use:     "bacotell [-n bot_name] [-t bot_token] [-p plugin_dir] [-c config_path]",
		Short:   "BacoTell - Pluggable bot client for Discord",
		Long:    `DESCRIPTION WIP`,
		Example: "",
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	flags := cmd.Flags()

	flags.StringP("name", "n", "", "name of this bot instance")
	viper.BindPFlag(common.ConfigBotName, flags.Lookup("name"))

	flags.StringP("token", "t", "", "bot token for Discord login")
	viper.BindPFlag(common.ConfigBotToken, flags.Lookup("token"))

	flags.StringP("plugins", "p", "", "path to the plugin directory")
	viper.BindPFlag(common.ConfigPluginDir, flags.Lookup("plugins"))

	flags.StringVarP(&configPath, "config", "c", "bacotell.config.toml", "path to the config file")

	common.InitConfig()
}

func main() {
	if cmd.Execute() != nil {
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigFile(configPath)

	viper.SetEnvPrefix("BACOTELL")
	viper.AutomaticEnv()

	if viper.ReadInConfig() != nil {
		viper.WriteConfig()
	}
}

func run(cmd *cobra.Command, args []string) {
	if viper.GetString(common.ConfigBotToken) == "" {
		common.GetLogger().Error("no bot token provided, use 'token' flag or set 'bot_token' in config file")
		os.Exit(1)
	}

	loader.LoadFromFolder()
	discord.Connect()
	loader.Unload()
}
