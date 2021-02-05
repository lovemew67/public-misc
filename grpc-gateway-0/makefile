## Define variables

# static variables
APP=grpc-gateway-0
SOURCE=./...
CONF=local.toml
PKGPATH=github.com/lovemew67/cornerstone

# derived variables

# variables from shell
GOPATH=$(shell env | grep GOPATH | cut -d'=' -f 2)
REVISION=$(shell git rev-list -1 HEAD)
TAG=$(shell git tag --sort=-version:refname --points-at HEAD | head -n1)
BR=$(shell git rev-parse --abbrev-ref HEAD)
DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
THISDIR=$(dir $(abspath $(firstword $(MAKEFILE_LIST))))

# if-else blocks
ifeq ($(TAG),)
TAG=$(REVISION)
endif

# exported variables


## Define targets
vendor:
	- rm -f go.sum
	GO111MODULE=on go build -mod=mod -v $(SOURCE) 
	GO111MODULE=on go mod tidy

server: vendor
	go install -mod=mod -i -v -ldflags "-s -X $(PKGPATH).appName=$(APP) -X $(PKGPATH).gitCommit=$(REVISION) -X $(PKGPATH).appVersion=$(TAG) -X $(PKGPATH).buildDate=$(DATE)" $(SOURCE) 
	$(GOPATH)/bin/$(APP) server --config $(CONF) 

genproto:
	protoc -I ./proto \
	--go_out ./gen/proto --go_opt paths=source_relative \
	--go-grpc_out ./gen/proto --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./gen/proto --grpc-gateway_opt paths=source_relative \
	--openapiv2_out ./gen/openapiv2 --openapiv2_opt logtostderr=true \
	./proto/*.proto

swaggerui:
	docker run --rm --name swaggerui -p 8080:8080 \
	-v $(THISDIR)gen/openapiv2/echo.swagger.json:/app/swagger.json \
	swaggerapi/swagger-ui:v3.37.2
