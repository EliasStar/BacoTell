package main

import (
	"os"

	"github.com/EliasStar/BacoTell/internal/bacotell"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	configPath string

	cmd = &cobra.Command{
		Run:     func(*cobra.Command, []string) { bacotell.Run() },
		Version: bacotell.Version,

		Use:   "bacotell [-n bot_name] [-t bot_token] [-p plugin_dir] [-c config_path]",
		Short: "BacoTell - Pluggable bot client for Discord",
	}
)

func init() {
	cobra.OnInitialize(func() {
		viper.SetConfigFile(configPath)

		viper.SetEnvPrefix("BACOTELL")
		viper.AutomaticEnv()

		if viper.ReadInConfig() != nil {
			viper.WriteConfig()
		}
	})

	flags := cmd.Flags()

	flags.StringP("name", "n", "", "name of this bot instance")
	viper.BindPFlag(bacotell.ConfigBotName, flags.Lookup("name"))

	flags.StringP("token", "t", "", "bot token for Discord login")
	viper.BindPFlag(bacotell.ConfigBotToken, flags.Lookup("token"))

	flags.StringP("plugins", "p", "", "path to the plugin directory")
	viper.BindPFlag(bacotell.ConfigPluginDir, flags.Lookup("plugins"))

	flags.StringP("log_file", "l", "", "path to log file")
	viper.BindPFlag(bacotell.ConfigLogFile, flags.Lookup("log_file"))

	flags.BoolP("verbose", "v", false, "set log level to DEBUG")
	flags.BoolP("quiet", "q", false, "set log level to OFF")
	viper.BindFlagValue(bacotell.ConfigLogLevel, logFlagHelper{flags.Lookup("verbose"), flags.Lookup("quiet")})

	flags.StringVarP(&configPath, "config", "c", "bacotell.config.toml", "path to the config file")

	bacotell.InitConfig()
}

func main() {
	if cmd.Execute() != nil {
		os.Exit(1)
	}
}

type logFlagHelper struct {
	verbose *pflag.Flag
	quiet   *pflag.Flag
}

var _ viper.FlagValue = logFlagHelper{}

// HasChanged implements viper.FlagValue.
func (h logFlagHelper) HasChanged() bool {
	return h.verbose.Changed || h.quiet.Changed
}

// Name implements viper.FlagValue.
func (h logFlagHelper) Name() string {
	return bacotell.ConfigLogLevel
}

// ValueType implements viper.FlagValue.
func (h logFlagHelper) ValueType() string {
	return "string"
}

// ValueString implements viper.FlagValue.
func (h logFlagHelper) ValueString() string {
	if h.verbose.Value.String() == "true" {
		return "debug"
	}

	if h.quiet.Value.String() == "true" {
		return "off"
	}

	return ""
}
