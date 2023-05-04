package static

import (
	_ "embed"
)

//go:embed Dockerfile-go
var Dockerfile_go []byte
