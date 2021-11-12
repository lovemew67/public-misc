## Define variables

# exported variables
export GO111MODULE=on


## Define targets

updatevendor:
	go mod tidy
	go mod vendor

runmain:
	go run main.go

test:
	go vet ./...
	go test -v -race -cover -timeout 180s ./...

bench:
	go vet ./...
	go test -v -bench=. -benchmem -run=none ./...

modrunmain: updatevendor
	go run -mod=vendor main.go

modtest: updatevendor
	go vet ./...
	go test -mod=vendor -v -race -cover -timeout 180s ./...

modbench: updatevendor
	go test -mod=vendor -v -bench=. -benchmem -run=none ./...
