## Define variables

# static variables
APP=golang-sample
SOURCE=./...
CONF=local.toml
PKGPATH=github.com/lovemew67/public-misc/cornerstone

# derived variables

# variables from shell
GOPATH=$(shell env | grep GOPATH | cut -d'=' -f 2)
REVISION=$(shell git rev-list -1 HEAD)
TAG=$(shell git tag --sort=-version:refname --points-at HEAD | head -n1)
DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# if-else blocks
ifeq ($(TAG),)
TAG=$(REVISION)
endif

# exported variables


## Define targets

genproto:
	protoc -I ./proto \
	--go_out ./gen/proto --go_opt paths=source_relative \
	--go-grpc_out ./gen/proto --go_opt paths=source_relative \
	--grpc-gateway_out ./gen/proto --grpc-gateway_opt paths=source_relative \
	--openapiv2_out ./gen/openapiv2 --openapiv2_opt logtostderr=true \
	./proto/*.proto

injecttag:
	protoc-go-inject-tag -input=./gen/proto/staff.pb.go

genall: genproto injecttag

mock:
	./mock-script

swaggerui:
	docker run --rm --name swaggerui -p 8080:8080 \
	-v $(THISDIR)gen/openapiv2/staff_service.swagger.json:/app/swagger.json \
	swaggerapi/swagger-ui:v3.37.2

updatevendor:
	go mod tidy
	go mod vendor
	go vet ./...

test: updatevendor	
	go test -mod=vendor -v -race -cover -timeout 180s ./...

bench: updatevendor
	go test -mod=vendor -v -bench=. -benchmem -run=none ./...

runmain: updatevendor
	go run -mod=vendor main.go

runserver: updatevendor
	go run -mod=vendor main.go server

server: updatevendor
	go install -mod=mod -v -ldflags "-s -X $(PKGPATH).appName=$(APP) -X $(PKGPATH).gitCommit=$(REVISION) -X $(PKGPATH).appVersion=$(TAG) -X $(PKGPATH).buildDate=$(DATE)" $(SOURCE) 
	$(GOPATH)/bin/$(APP) server --config $(CONF)
