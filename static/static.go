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

//go:embed manifests/servicemonitor.yaml
var Servicemonitor []byte

//go:embed manifests/deploy-app.yaml
var Deploy_app []byte

//go:embed manifests/deploy-app-binary.yaml
var Deploy_app_binary []byte

//go:embed manifests/ds-app-binary.yaml
var DS_app_binary []byte

//go:embed manifests/ds-app.yaml
var DS_app []byte

//go:embed manifests/po-batch.yaml
var PO_batch []byte
