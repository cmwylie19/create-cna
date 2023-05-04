# Create DDIS App

Create DDIS app is meant to make life easier for developers and data scientist. The CLI scaffolds new DDIS applications complete with a `Makefile` for quick compilations and image building/pushing, a manifests folder containing the Kubernetes required to deploy the app, and a `README.md` explaining how to use the repo.

`create-ddis-app` supports a multitude of languages including `python`, `go`, `rust`, `frontend-js`, `node`, and `python-batch` that are meant to be deployed in Kubernetes with instrumentation.

[Download](#download)  - How to download the `create-ddis-app` tool.  
[Getting Started](#getting-started) - How to create a new app.  
[User Guide](#user-guide) -  How to develop apps bootstrapped with Create DDIS App.

Create DDIS App works on macOS, Linux, and Windows. The artifacts produced by `create-ddis-app` include will require make to be installed on the host machine. 

If something doesn't work, [file an issue](https://github.com/cmwylie19/create-ddis-app/issues/new?assignees=&labels=&template=bug_report.md&title=).  
If there is a feature you would like to see, [file a feature request](https://github.com/cmwylie19/create-ddis-app/issues/new?assignees=&labels=&template=feature_request.md&title=).


## Download

Download the latest release for your operator system from the [repo](https://github.com/cmwylie19/create-ddis-app).

- [macOS]()
- [Linux]()
- [Windows]() 

## Getting Started

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
- `Makefile` to drive compilation and image building/pushing

