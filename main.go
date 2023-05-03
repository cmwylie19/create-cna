/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/cmwylie19/create-ddis-app/pkg/log"
	"github.com/cmwylie19/create-ddis-app/cmd"
)

var logger log.Logger

func main() {
	cmd.getRootCommand(logger).Execute(); err != nil {
		logger.Infof("Error executing command: %v", err)
		os.Exit(1)
	}
}

func init() {
	logger = log.Logger{
		Debug: true,
	}
}
