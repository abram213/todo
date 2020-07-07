package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "start",
	Short: "TODO App",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting todo app...")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(errors.Wrap(err, "unable to start app"))
		os.Exit(1)
	}
}

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is config.yaml)")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(errors.Wrap(err, "unable to read config from file"))
		os.Exit(1)
	}
}
