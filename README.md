# Create Cloud Native App
`create-cna` is meant to make life easier for developers and data scientist who just want to write code and not spend too much time in configuration. The CLI scaffolds new cloud native applications complete with a `Makefile` for quick compilations and image building/pushing, a manifests folder containing the Kubernetes objects required to deploy the app, and a `README.md` explaining how to use the repo.

`create-cna` is intended to support a multitude of languages including `python`, `go`, `rust`, `frontend-js`, `node`, and `python-batch` that are meant to be deployed in Kubernetes with instrumentation from the day 1. However, for now, only `go` is supported as this project is not actively being developed.

[Compile App](#compile-app)
[Getting Started](#getting-started) - How to create a new app.  

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

Scaffolding an application

```bash
create-cna scaffold --help
Scaffold a new project based on the langauge and type of project you want to create.

Usage:
  create-cna scaffold --name=<project-name> --type=<project-type> --port=<port> --metrics=<metrics>

Example:
  create-cna scaffold --name=go-server --type=go --port=8080 --metrics=true

  create-cna scaffold --name=migration-job --type=go --port=9191 --metrics=false --verbose=false

Usage:
  create-cna scaffold [flags]

Flags:
  -h, --help          help for scaffold
      --metrics       Metrics of the project: true or false
      --name string   Name of the project
      --port string   Port of the project: 8080
      --type string   Type of the project: go, rust, python-batch, python, frontend
      --verbose       Verbose logs: true (default) or false (default true)
```

Example scaffolding a go app  

```bash
create-cna scaffold --name=go-server --port=8080 --type=go  --metrics=true --verbose=true

➜  go-server git:(main) ✗ tree .
.
├── Makefile
├── build
│   └── Dockerfile
├── main.go
└── manifests
    ├── k8s
    │   └── app.yaml
    └── observability
        └── servicemonitor.yaml

5 directories, 5 files

➜  cat manifests/k8s/app.yaml  
.
apiVersion: v1
kind: Namespace
metadata:
  creationTimestamp: null
  name: go-server
  namespace: go-server
spec: {}
---
apiVersion: v1
kind: LimitRange
metadata:
  name: cpu-mem-resource-constraint
  namespace: go-server
spec:
  limits:
  - default: # this section defines default limits
      cpu: "500m"
      memory: "128Mi"
    defaultRequest: 
      cpu: "250m"
      memory: "64Mi"
    max: 
      cpu: "1"
      memory: "800Mi"
    min:
      cpu: "100m"
      memory: "100Mi"
    type: Container
---
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: go-server
  name: go-server
spec:
  replicas: 1
  selector:
    matchLabels:
      app: go-server
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: go-server
    spec:
      serviceAccountName: go-server
      containers:
      - image: <APP_IMAGE>
        name: go-server
        command: ["./go-server"]
      resources:
        requests:
          memory: "64Mi"
          cpu: "250m"
        limits:
          memory: "128Mi"
          cpu: "500m"
        securityContext:
          allowPrivilegeEscalation: false
        ports:
        - containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: null
  labels:
    app: go-server
  name: go-server
  namespace: default
spec:
  ports:
  - port: 8080
    protocol: TCP
    targetPort: 8080
    name: http
  selector:
    app: go-server
---
apiVersion: v1
kind: ServiceAccount
metadata:
  creationTimestamp: null
  name: go-server

➜  cat manifests/observability/servicemonitor.yaml 
.
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: go-server
  namespace: go-server
  labels:
    app: go-server
spec:
  selector:
    matchLabels:
      app: go-server
  endpoints:
  - port: http% 
```

The `scaffold` option creates:
- create a folder for the application
- `build` folder with a dockerfile
- `k8s` folder with Kubernetes manifests, including telemetry and prometheus artifacts if selected
- `observability` folder is for ServiceMonitor if selected
- `Makefile` to drive compilation and image building/pushing
- `README.md` with instructions on how to compile/build/push images

