/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/cmwylie19/create-ddis-app/cmd"
	"github.com/cmwylie19/create-ddis-app/pkg/log"
)

var logger *log.Log

func main() {
	if err := cmd.GetRootCommand(logger).Execute(); err != nil {
		logger.Infof("Error executing command: %v", err)
		os.Exit(1)
	}
}
func init() {
	logger = &log.Log{
		Debug: true,
	}
}
