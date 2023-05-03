/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/cmwylie19/create-ddis-app/internal/log"

	"github.com/spf13/cobra"
)

func getRootCommand(logger log.Logger) *cobra.Command {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(getScaffoldCommand(logger))
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
