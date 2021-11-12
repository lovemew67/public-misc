PKGPATH=github.com/lovemew67/public-misc/confluence-hater/command
SOURCE=./...
DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
REVISION=$(shell git rev-list -1 HEAD)
TAG=$(shell git tag -l --points-at HEAD)
ifeq ($(TAG),)
TAG=$(REVISION)
endif

export GO111MODULE=on

test:
	@echo "start unit tests & vet"
	go vet $(SOURCE)
	go test -mod=vendor -race -cover $(SOURCE)

modvendor:
	go mod tidy
	go mod vendor

install:
	go install -mod=vendor -v -ldflags "-s -X $(PKGPATH).gitCommit=$(REVISION) -X $(PKGPATH).appVersion=$(TAG) -X $(PKGPATH).buildDate=$(DATE)" $(SOURCE)
