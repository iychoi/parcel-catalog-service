# Copyright 2020 CyVerse
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

PKG=github.com/iychoi/parcel-catalog-service
CATALOG_SERVICE_BUILD_IMAGE=parcel_catalog_service_build
CATALOG_SERVICE_BUILD_DOCKERFILE=deploy/image/parcel_catalog_service_build.dockerfile
CATALOG_SERVICE_IMAGE?=iychoi/parcel-catalog-service
CATALOG_SERVICE_DOCKERFILE=deploy/image/parcel_catalog_service_image.dockerfile
VERSION=v0.1.0
GIT_COMMIT?=$(shell git rev-parse HEAD)
BUILD_DATE?=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS?="-X ${PKG}/pkg/service.serviceVersion=${VERSION} -X ${PKG}/pkg/service.gitCommit=${GIT_COMMIT} -X ${PKG}/pkg/service.buildDate=${BUILD_DATE}"
GO111MODULE=on
GOPROXY=direct
GOPATH=$(shell go env GOPATH)

.EXPORT_ALL_VARIABLES:

.PHONY: parcel-catalog-service
parcel-catalog-service:
	mkdir -p bin
	CGO_ENABLED=1 GOOS=linux go build -ldflags ${LDFLAGS} -o bin/parcel-catalog-service ./cmd/

.PHONY: parcel_catalog_service_build
service_build:
	docker build -t $(CATALOG_SERVICE_BUILD_IMAGE):latest -f $(CATALOG_SERVICE_BUILD_DOCKERFILE) .

.PHONY: image
image: parcel_catalog_service_build
	docker build -t $(CATALOG_SERVICE_IMAGE):latest -f $(CATALOG_SERVICE_DOCKERFILE) .

.PHONY: push
push: image
	docker push $(CATALOG_SERVICE_IMAGE):latest

.PHONY: image-release
image-release:
	docker build -t $(CATALOG_SERVICE_IMAGE):$(VERSION) -f $(CATALOG_SERVICE_DOCKERFILE) .

.PHONY: push-release
push-release:
	docker push $(CATALOG_SERVICE_IMAGE):$(VERSION)
