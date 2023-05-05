package static

import (
	_ "embed"
)

//go:embed build/Dockerfile-go
var Dockerfile_go []byte

//go:embed make/Makefile-go
var Makefile_go []byte

//go:embed .gitignore
var Gitignore []byte

//go:embed apps/main.txt
var Main_go []byte
