# globals
BINARY_NAME?=video-game-api
BUILD_DIR?="./build"

# basic Go commands
GOCMD=go
GOBUILD=$(GOCMD) build

# images
KIND=kindest/node:v1.33.1
KIND_CLUSTER=video-game-cluster

VERSION=1.0.0
VIDEO_GAME_IMAGE=video-game-api:$(VERSION)

# kind stuff
NAMESPACE=video-game
VIDEO_GAME_SERVICE=video-game-api

# linting
define get_latest_lint_release
	curl -s "https://api.github.com/repos/golangci/golangci-lint/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/'
endef

LATEST_LINT_VERSION=$(shell $(call get_latest_lint_release))
INSTALLED_LINT_VERSION=$(shell golangci-lint --version 2>/dev/null | awk '{print "v"$$4}')

# get GOPATH according to OS
ifeq ($(OS),Windows_NT) # is Windows_NT on XP, 2000, 7, Vista, 10...
    GOPATH=$(go env GOPATH)
else
    GOPATH=$(shell go env GOPATH)
endif

.PHONY: install-linter
install-linter:
ifneq "$(INSTALLED_LINT_VERSION)" "$(LATEST_LINT_VERSION)"
	@echo "new golangci-lint version found:" $(LATEST_LINT_VERSION)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/main/install.sh | sh -s -- -b $(GOPATH)/bin latest
endif

# targets
build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $(MAKE) build

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(MAKE) build

build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(MAKE) build

build-windows-amd64:
	GOOS=windows GOARCH=amd64 $(MAKE) build-windows

# MCP targets
build-darwin-arm64-mcp:
	GOOS=darwin GOARCH=arm64 BINARY_NAME=mcp-server $(MAKE) build-mcp

.PHONY: build
build:
	CGO_ENABLED="0" GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -v \
		-o ${BUILD_DIR}/$(BINARY_NAME)-$(GOOS)-$(GOARCH) ./cmd/$(BINARY_NAME)/main.go

.PHONY: build-windows
build-windows:
	CGO_ENABLED="0" GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -v \
		-o ${BUILD_DIR}/$(BINARY_NAME)-$(GOOS)-$(GOARCH).exe ./cmd/$(BINARY_NAME)/main.go

.PHONY: build-mcp
build-mcp:
	CGO_ENABLED="0" GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -v \
		-o ${BUILD_DIR}/$(BINARY_NAME)-$(GOOS)-$(GOARCH) ./cmd/$(BINARY_NAME)/main.go

.PHONY: lint
lint: install-linter
	golangci-lint run ./...

.PHONY: test
test:
	go test -count=1 -race -covermode=atomic -coverprofile=coverage.out ./...

.PHONY: test-integration
test-integration:
	docker compose down && docker compose run --build --rm integration-test && docker compose down

.PHONY: mockgen
mockgen:
	mockery

proto:
	protoc \
    		--go_out=./pkg/grpc \
    		--go_opt=module=github.com/gandarez/video-game-api/pkg/grpc \
    		--go-grpc_out=./pkg/grpc \
			--go-grpc_opt=require_unimplemented_servers=false \
    		--go-grpc_opt=module=github.com/gandarez/video-game-api/pkg/grpc \
    		./api/protos/*.proto

create.migration:
	migrate create -ext sql -format unix -dir db/migrations $(name)

# ------------------------
# --------- kind ---------
# ------------------------

#1
kind-up:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config ./deployments/kind/kind-config.yaml
	
	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

#2
kind-load:
	kind load docker-image $(VIDEO_GAME_IMAGE) --name $(KIND_CLUSTER) & \
	wait;

#3..#4..#5
kind-apply:	
	kubectl apply -f ./deployments/app --recursive
	kubectl wait pods --namespace=$(NAMESPACE) --selector app=$(VIDEO_GAME_SERVICE) --timeout=120s --for=condition=Ready

kind-downsize-1:
	docker update --cpuset-cpus="0" video-game-cluster-control-plane
	docker restart video-game-cluster-control-plane

# let's see GOMAXPROCS for go 1.25 in action
# DEBUG
# docker exec -it video-game-cluster-control-plane bash
# run nproc inside container to check it) - it should show 12 for M2 Max processor
# after downsizing it should show 2
kind-downsize-2:
	docker update --cpuset-cpus="0-1" video-game-cluster-control-plane
	docker restart video-game-cluster-control-plane

.PHONY: load-test
load-test:
	hey -m GET -c 10 -n 1000 "http://localhost:17020/consoles\?page\=1&rows\=10"

# 	hey -m GET -c 10 -n 1000 "http://localhost:17020/consoles/0eee8295-9d2e-4c19-af43-1e5464c64eb6"

.PHONY: dev-status
dev-status:
	watch -n 2 kubectl get pods --all-namespaces

.PHONY: list-consoles
list-consoles:
	curl http://localhost:17020/consoles

.PHONY: dev-liveness
dev-liveness:
	curl http://localhost:17020/liveness

.PHONY: dev-readiness
dev-readiness:
	curl http://localhost:17020/readiness

build-docker-image:
	docker build . -t $(VIDEO_GAME_IMAGE)
