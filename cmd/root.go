/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func getRootCommand() *cobra.Command {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	return rootCmd
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "create-ddis-app",
	Short: "Simplify, scaffold and automate DDIS repos, applications and tasks.",
	Long: `A tool to simplify, scaffold, and automate DDIS repositories and applications
allowing users to focus on delivering value while avoiding cumbersome config.

Create DDIS App is a CLI app that scaffolds applications
based on industry best practices.
This application is a tool to generate the needed files
and folders to quickly create an  application.

usage:
create-ddis-app help 
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.create-ddis-app.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
