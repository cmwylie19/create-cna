# Create Cloud Native App
`create-cna` is meant to make life easier for developers and data scientist who just want to write code and not spend too much time in configuration. The CLI scaffolds new cloud native applications complete with a `Makefile` for quick compilations and image building/pushing, a manifests folder containing the Kubernetes objects required to deploy the app, and a `README.md` explaining how to use the repo.

`create-cna` is intended to support a multitude of languages including `python`, `go`, `rust`, `frontend-js`, `node`, and `python-batch` that are meant to be deployed in Kubernetes with instrumentation from the day 1. However, for now, only `go` is supported as this project is not actively being developed.

[Compile App](#compile-app)
[Getting Started](#getting-started) - How to create a new app.  
[User Guide](#user-guide) -  How to develop apps bootstrapped with `create-cna`.

`create-cna` works on macOS and Linux. The artifacts produced by the cli include will require [`Make`](https://www.gnu.org/software/make/) to be installed on the host machine. 

If something doesn't work, [file an issue](https://github.com/cmwylie19/create-ddis-app/issues/new?assignees=&labels=&template=bug_report.md&title=).  
If there is a feature you would like to see, [file a feature request](https://github.com/cmwylie19/create-ddis-app/issues/new?assignees=&labels=&template=feature_request.md&title=).

## Compile App

Linux/Mac    

```bash
go build -o create-cna && chmod 711 create-cna && sudo mv create-cna /usr/local/bin
```
## Getting Started

Getting help

```bash
create-cna --help

A tool to simplify, scaffold, and automate cloud native repositories and applications
allowing users to focus on delivering value while avoiding cumbersome config.

Create cloud native app is a CLI that scaffolds applications
based on industry best practices in order to quickly
deploy cloud native apps instrumented for Kubernetes, 
metric, and telemetry collection.

usage:
create-cna help

Usage:
  create-cna [command]

Available Commands:
  help        Help about any command
  scaffold    Scaffold a new project

Flags:
  -h, --help   help for create-cna

Use "create-cna [command] --help" for more information about a command.
```

```bash
create-cna scaffold --name=go-server --port=8080 --type=go  --metrics=true --verbose=true
```

The `scaffold` option creates:
- create a folder for the application
- `build` folder with a dockerfile
- `k8s` folder with Kubernetes manifests, including telemetry and prometheus artifacts if selected
- `observability` folder is for ServiceMonitor if selected
- `Makefile` to drive compilation and image building/pushing
- `README.md` with instructions on how to compile/build/push images

