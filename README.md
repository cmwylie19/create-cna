# Create DDIS App
Create DDIS apps quickly and easily scaffold apps with a multitude of languages including `python`, `go`, `rust`, `frontend-js`, `node`, and `python-batch` (for workloads that variable scalable memory and CPU). The tool allows users to add support for metrics (`serviceMonitor`), kubernetes and scaffolds a `Makefile` for quick and easily building of images.  


[Getting Started](#getting-started) - How to create a new app.  
[User Guide](#user-guide) -  How to develop apps bootstrapped with Create DDIS App.

Create DDIS App works on macOS, Linux, and Windows.  
If something doesn't work, [file an issue](https://github.com/cmwylie19/create-ddis-app/issues/new?assignees=&labels=&template=bug_report.md&title=).  
If there is a feature you would like to see, [file a feature request](https://github.com/cmwylie19/create-ddis-app/issues/new?assignees=&labels=&template=feature_request.md&title=).

## Overview

```bash
go run main.go scaffold --name=go-server --port=8080 --type=go --verbose=true

create-ddis-app scaffold --name=sb-proxy --type=rust --telemetry --metrics --ci --gitops --port=8080 --controller=deployment

create-ddis-app scaffold --name=jdv-migration --type python-batch --telemetry=false --metrics=true --ci --port=8080 --controller=cronjob
```

The `scaffold` option creates:
- create a folder for the application
- `build` folder with a dockerfile
- `k8s` folder with Kubernetes manifests, including telemetry and prometheus artifacts if selected
- `gitops` folder is Argo manifests if selected

