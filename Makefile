##################################################
# Variables                                      #
##################################################
IMAGE_TAG      ?= master
IMAGE_REGISTRY ?= hub.docker.com
IMAGE_REPO     ?= turnoutt

IMAGE_CONTROLLER = $(IMAGE_REPO)/azure-keyvault-operator:$(IMAGE_TAG)


ARCH       ?=amd64
CGO        ?=0
TARGET_OS  ?=linux

GIT_VERSION = $(shell git describe --always --abbrev=7)
GIT_COMMIT  = $(shell git rev-list -1 HEAD)
DATE        = $(shell date -u +"%Y.%m.%d.%H.%M.%S")

##################################################
# All                                            #
##################################################
.PHONY: All
all: test build

##################################################
# Tests                                          #
##################################################
.PHONY: test
test:
	go test ./...

.PHONY: e2e-test
e2e-test:
	TERMINFO=/etc/terminfo
	TERM=linux
	@az login --service-principal -u $(AZURE_SP_ID) -p "$(AZURE_SP_KEY)" --tenant $(AZURE_SP_TENANT)
	@az aks get-credentials \
		--name azure-keyvault-operator-nightly-run \
		--subscription $(AZURE_SUBSCRIPTION) \
		--resource-group $(AZURE_RESOURCE_GROUP)
	npm install --prefix tests
	IMAGE_CONTROLLER=$(IMAGE_CONTROLLER) IMAGE_ADAPTER=$(IMAGE_ADAPTER) npm test --verbose --prefix tests

##################################################
# PUBLISH                                        #
##################################################
.PHONY: publish
publish: build
	docker push $(IMAGE_CONTROLLER)

##################################################
# Build                                          #
##################################################
GO_BUILD_VARS= GO111MODULE=on CGO_ENABLED=$(CGO) GOOS=$(TARGET_OS) GOARCH=$(ARCH)

.PHONY: build
build: build-controller

.PHONY: build-controller
build-controller: 
	$(GO_BUILD_VARS) operator-sdk build $(IMAGE_CONTROLLER) \
		--go-build-args "-ldflags -X=main.GitCommit=$(GIT_COMMIT) -o build/_output/bin/azure-keyvault-operator" --verbose

.PHONY: generate-api
generate-api:
	$(GO_BUILD_VARS) operator-sdk generate k8s
	$(GO_BUILD_VARS) operator-sdk generate openapi

.PHONY: setup-debug-environment
setup-debug-environment:
	curl -sL https://aka.ms/InstallAzureCLIDeb | sudo bash
	curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
	chmod +x ./kubectl
	sudo mv ./kubectl /usr/local/bin/kubectl
