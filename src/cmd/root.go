package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	. "./utils"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "fidipides",
	Short: "Fidípides",
	Long: `Fidípides (IMAP Client to export e-mails).

Fidípides is the IMAP client that allows you to list email boxes and export emails to your location
through an easily executed and intuitive client

Required flags:
  --server   	Server name or IP to connect
  --port    	TCP port to use when communicating with the server
  --tls   		Use TLS when communicating with the server
  --username    e-mail to login with server
  --password   password to authorize access

`,
}

func init() {
	// the config is only loaded if the command is valid,
	// that is why we use OnInitialize
	cobra.OnInitialize(initConfig)
	// using this so i will check manually for strange behavior of the cli
	RootCmd.SilenceErrors = true
	RootCmd.SilenceUsage = true

	// change the suggestion distance of the commands
	RootCmd.SuggestionsMinimumDistance = 3
	RootCmd.PersistentFlags().BoolVar(&DebugFlag, "debug", false, "debug mode")

	RootCmd.PersistentFlags().StringVar(&CfgServerAddress, "server", "", "Server name or IP to connect")
	RootCmd.PersistentFlags().StringVar(&CfgServerPort, "port", "993", "TCP port to use when communicating with the server")
	RootCmd.PersistentFlags().BoolVar(&CfgServerUseTLS, "tls", false, "Use TLS when communicating with the server")

	RootCmd.PersistentFlags().StringVar(&CfgUsername, "username", "", "Username to login with server")
	RootCmd.PersistentFlags().StringVar(&CfgPassword, "password", "", "Password to authorize access")

	RootCmd.PersistentFlags().MarkHidden("debug")
}

func initConfig() {
	// defaults
	viper.SetDefault("debug", DebugFlag)
	// get from ENV
	viper.AutomaticEnv()
}
