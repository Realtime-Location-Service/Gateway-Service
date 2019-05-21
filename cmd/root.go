package cmd

import (
	"os"

	"github.com/rls/gateway-service/pkg/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd is the root command of Gateway Service
var RootCmd = &cobra.Command{
	Use:   "gs",
	Short: "Gateway Service will call to an appropriate API with necessary data ",
	Long:  "Gateway Service will call to an appropriate API with necessary data ",
}

// Execute executes the root command
func Execute() {
	config.Init()

	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yml", "config file")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		return
	}
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
}
