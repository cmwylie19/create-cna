package log // import "github.com/create-ddis-app/cli/pkg/log"

import (
	"fmt"
	"log"
)

type Logger struct {
	Debug bool
}

// Log to stdout without timestamp
func (logger Logger) Printf(format string, args ...interface{}) {
	fmt.Printf("create-ddis-app: "+format, args...)
}

func (logger Logger) Print(format string) {
	fmt.Printf("create-ddis-app: " + format)
}

// Log to stdout only if Debug is true.
func (logger Logger) Infof(format string, args ...interface{}) {
	if logger.Debug {
		log.Printf(format+"\n", args...)
	}
}

// Log to stdout only if Debug is true.
func (logger Logger) Info(format string) {
	if logger.Debug {
		log.Println(format)
	}
}

// Log to stdout always.
func (logger Logger) Warnf(format string, args ...interface{}) {
	fmt.Printf("create-ddis-app: warning - %s\n"+format+"\n", args...)
}

// Log to stdout always.
func (logger Logger) Warn(format string) {
	fmt.Println("create-ddis-app: warning\n" + format)
}

// Log to stdout always.
func (logger Logger) Errorf(format string, args ...interface{}) {
	fmt.Printf("create-ddis-app: error - "+format, args...)
}

// Log to stdout always.
func (logger Logger) Error(format string) {
	fmt.Println("create-ddis-app: error - " + format)
}
