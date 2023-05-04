package log // import "github.com/create-ddis-app/cli/pkg/log"

import (
	"fmt"
	"log"
)

// Iterface for logging.
type Logger interface {
	// Infof may be used to write messages to the log Printf style
	Infof(format string, args ...interface{})
	// Info may be used to write messages to the log.
	Info(format string)
	// Warn may be used to write warning messages to the log Printf style
	Warn(format string)
	// Warnf may be used to write warning messages to the log Printf style
	Warnf(format string, args ...interface{})
	// Errorf may be used to write error messages to the log Printf style
	Errorf(format string, args ...interface{})
	// Error may be used to write error messages to the log.
	Error(format string)
	// Printf may be used to write messages to the log Printf style without timestamp
	Printf(format string, args ...interface{})
	// Print may be used to write messages to the log without timestamp
	Print(format string)
	// Set log verbosity
	V(v bool)
}

// implementation of Logger interface
type Log struct {
	Debug bool
}

// Log to stdout only if Debug is true.
func (logger *Log) V(verbose bool) {
	logger.Debug = verbose
}

// Log to stdout without timestamp
func (logger Log) Printf(format string, args ...interface{}) {
	fmt.Printf("create-ddis-app: "+format, args...)
}

func (logger Log) Print(format string) {
	fmt.Printf("create-ddis-app: " + format)
}

// Log to stdout only if Debug is true.
func (logger *Log) Infof(format string, args ...interface{}) {
	if logger.Debug {
		log.Printf(format+"\n", args...)
	}
}

// Log to stdout only if Debug is true.
func (logger *Log) Info(format string) {
	if logger.Debug {
		log.Println(format)
	}
}

// Log to stdout always.
func (logger Log) Warnf(format string, args ...interface{}) {
	fmt.Printf("create-ddis-app: warning - %s\n"+format+"\n", args...)
}

// Log to stdout always.
func (logger Log) Warn(format string) {
	fmt.Println("create-ddis-app: warning\n" + format)
}

// Log to stdout always.
func (logger Log) Errorf(format string, args ...interface{}) {
	fmt.Printf("create-ddis-app: error - "+format, args...)
}

// Log to stdout always.
func (logger Log) Error(format string) {
	fmt.Println("create-ddis-app: error - " + format)
}
