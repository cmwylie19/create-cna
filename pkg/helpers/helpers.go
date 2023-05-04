package helpers

import (
	"errors"
	"os"
)

//go:embed static/Makefile-go
var Makefile_go []byte

//go:embed static/Dockerfile-go
var Dockerfile_go []byte

//go:embed static/.dockerignore
var Dockerignore []byte

//go:embed static/.gitignore-go
var Gitignore []byte

//go:embed static/main.go
var Main_go []byte

// CreateDir creates a directory if it does not exist
func CreateDir(path string) error {
	var err error
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return err
}

func CreateDockerfile(path, nane, _type string) error {
	var dockerfile []byte

}
func CreateMainfile(path, nane, _type string) error {
	var mainfile []byte

}
func CreateMakefile(path, nane, _type string) error {
	var makefile []byte

}

func CreateServiceMonitor(path string) error {
	var servicemonitor []byte

}
