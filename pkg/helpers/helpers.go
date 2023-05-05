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

func CreateMainfile(path, _type string) error {
	var err error
	var mainfile []byte
	var mainfile_name string

	switch _type {
	case "go":
		mainfile = static.Main_go
		mainfile_name = "main.go"
	}

	err = os.WriteFile(path+"/"+mainfile_name, mainfile, 0644)
	if err != nil {
		return err
	}
	return err

}

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

func CreateServiceMonitor(path, name string) error {
	var err error
	var servicemonitor []byte

	servicemonitor = ByteReplace(static.Servicemonitor, "<APP_NAME>", name)

	err = os.WriteFile(path+"/servicemonitor.yaml", servicemonitor, 0644)
	if err != nil {
		return err
	}
	return err
}

func CreateKubernetesManifests(path, controller, port, name string) error {
	var err error
	var kubernetes []byte

	switch strings.ToLower(controller) {
	case "deployment":
		// replace <APP_NAME> with name and <APP_PORT> with port
		kubernetes = ByteReplace(ByteReplace(static.Deploy_app, "<APP_NAME>", name), "<APP_PORT>", port)
	case "daemonset":
		// replace <APP_NAME> with name and <APP_PORT> with port
		kubernetes = ByteReplace(ByteReplace(static.DS_app, "<APP_NAME>", name), "<APP_PORT>", port)
	case "pod":
		kubernetes = ByteReplace(static.PO_batch, "<APP_NAME>", name)
	}

	err = os.WriteFile(path+"/app.yaml", kubernetes, 0644)
	if err != nil {
		return err
	}
	return err
}
