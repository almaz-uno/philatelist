// Copyright © 2017 Maxim Kovrov <maksim.kovrov@gmail.com>

package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	homedir "github.com/mitchellh/go-homedir"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

const executionName = "philatelist"

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   executionName,
	Short: "This program collects images from Google",
	Long: `Geoimage is a RESTful server for searching images based on location. 
It uses Google API for do this.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/."+executionName+".yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	log.SetOutput(os.Stdout)
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName("." + executionName)
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file:", viper.ConfigFileUsed())
	}
}

// exitHook waits for signals `syscall.SIGINT` and `syscall.SIGTERM` and exits from process when caught one of them
func exitHook() {
	// awaiting for termination
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	log.Println("Waiting exit signal...")
	sig := <-sigs
	log.Infof("Caught signal '%v'. Exiting...", sig)
	os.Exit(0)
}
