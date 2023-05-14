# Makefile for building the go app + docker image.
SHELL=bash
DOCKER_USERNAME=<YOUR_USERNAME>
IMAGE_REGISTRY=<YOUR_IMAGE_REGISTRY>
APP_NAME=go-server
VERSION=0.0.1
ARCH=amd64



IMAGE ?= ${IMAGE_REGISTRY}/${DOCKER_USERNAME}/${APP_NAME}:${VERSION}
IMAGE_REPLACE ?= ${IMAGE_REGISTRY}\/${DOCKER_USERNAME}\/${APP_NAME}:${VERSION}

#---------------------------
# Build the secret-watcher binary

.PHONY: compile/app
compile/app:
	GOARCH=${ARCH} GOOS=linux go build -ldflags="-s -w" -o build/${APP_NAME} *.go


#---------------------------
# Build the docker image

.PHONY: build/image
build/image:
	docker build -t $(IMAGE) build/

#--------------------------------
# Push the docker image to dockerhub
.PHONY: push/image
push/image:
	docker push $(IMAGE)

#--------------------------------
# Update Kubernetes manifests with image
.PHONY: update/manifests
update/manifests:
	@cat manifests/k8s/app.yaml | sed 's/<APP_IMAGE>/${IMAGE_REPLACE}/g;' >  manifests/k8s/app.generated.yaml
	@mv manifests/k8s/app.generated.yaml manifests/k8s/app.yaml

#--------------------------------
all: compile/app build/image push/image
	@echo "Done building ${APP_NAME} for ${ARCH}"
