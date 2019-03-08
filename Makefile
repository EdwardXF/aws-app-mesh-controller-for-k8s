PKG=github.com/aws/aws-app-mesh-controller-for-k8s
IMAGE=amazon/app-mesh-controller
REPO=$(AWS_ACCOUNT).dkr.ecr.$(AWS_REGION).amazonaws.com/$(IMAGE)
VERSION=0.1.0-alpha
GIT_COMMIT?=$(shell git rev-parse HEAD)
BUILD_DATE?=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS?="-X main.version=${VERSION} -X main.gitCommit=${GIT_COMMIT} -X main.buildDate=${BUILD_DATE} -s -w"
GO111MODULE=on

ifndef AWS_REGION
$(error AWS_REGION is not set)
endif

ifndef AWS_ACCOUNT
$(error AWS_ACCOUNT is not set)
endif

.PHONY: eks-appmesh-controller
eks-appmesh-controller:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux go build -ldflags ${LDFLAGS} -o bin/app-mesh-controller ./cmd/app-mesh-controller

.PHONY: darwin
darwin:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=darwin go build -ldflags ${LDFLAGS} -o bin/app-mesh-controller-osx ./cmd/app-mesh-controller

.PHONY: code-gen
code-gen:
	./hack/update-codegen.sh

.PHONY: image
image:
	docker build -t $(IMAGE):latest .

.PHONY: image-release
image-release:
	docker build -t $(IMAGE):$(VERSION) .

.PHONY: push
push:
	docker tag $(IMAGE):latest $(REPO):latest
	docker push $(REPO):latest

.PHONY: push-release
push-release:
	docker tag $(IMAGE):$(VERSION) $(REPO):$(VERSION)
	docker push $(REPO):$(VERSION)

.PHONY: kube-deploy
kube-deploy:
	./hack/deploy.sh
