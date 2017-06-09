// Copyright © 2017 Maxim Kovrov <maksim.kovrov@gmail.com>

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// this parameters will be set via -ldflags "-X", see Makefile for additional info
var (
	buildApplicationName = executionName
	buildVersion         = "n/a"
	buildTime            = "n/a"
	buildCommit          = ""
	buildBranch          = ""
	showFullVersion      bool
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version of the application",
	Long:  `Shows version of the application and exits`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(buildApplicationName, "version", buildVersion)
		if showFullVersion {
			showParam("commit    ", buildCommit)
			showParam("branch    ", buildBranch)
			showParam("build time", buildTime)
		}
	},
}

func showParam(name, value string) {
	if value == "" {
		return
	}
	fmt.Println(name, value)
}

func init() {
	RootCmd.AddCommand(versionCmd)
	versionCmd.PersistentFlags().BoolVarP(&showFullVersion, "--all", "a", false, "Output all available information about build")
}
