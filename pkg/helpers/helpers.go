package helpers

import (
	"errors"
	"os"
	"strings"

	"github.com/cmwylie19/create-ddis-app/static"
)

// //go:embed static/.dockerignore
// var Dockerignore []byte

// //go:embed static/.gitignore-go
// var Gitignore []byte

// //go:embed static/main.go
// var Main_go []byte

// CreateDir creates a directory if it does not exist
func CreateDir(path string) error {
	var err error
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return err
}

func CreateGitIgnore(path string) error {
	var err error
	err = os.WriteFile(path+"/.gitignore", static.Gitignore, 0644)
	if err != nil {
		return err
	}
	return err
}

func ByteReplace(file []byte, old, new string) []byte {
	return []byte(strings.ReplaceAll(string(file), old, new))
}
func CreateDockerfile(path, _type, name string) error {
	var err error
	var dockerfile []byte

	switch _type {
	case "go":
		// Convert to String, Replace <APP_NAME> with name, convert back to []byte
		dockerfile = ByteReplace(static.Dockerfile_go, "<APP_NAME>", name)
	}
	err = os.WriteFile(path+"/Dockerfile", dockerfile, 0644)
	if err != nil {
		return err
	}
	return err
}

// func CreateMainfile(path, _type string) error {
// 	var mainfile []byte

// }

func CreateMakefile(path, _type string) error {

	var err error
	var makefile []byte

	switch _type {
	case "go":
		// Convert to String, Replace <APP_NAME> with name, convert back to []byte
		makefile = ByteReplace(static.Makefile_go, "<APP_NAME>", path)
	}
	err = os.WriteFile(path+"/Makefile", makefile, 0644)
	if err != nil {
		return err
	}
	return err

}

// func CreateServiceMonitor(path string) error {
// 	var servicemonitor []byte

// }
