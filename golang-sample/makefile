## Define variables

# exported variables
export GO111MODULE=on


## Define targets

run:
	go run main.go

test:
	go vet ./...
	go test -v -race -cover -timeout 180s ./...

bench:
	go vet ./...
	go test -v -bench=. -benchmem -run=none ./...

modrun:
	go run -mod=vendor main.go

modvendor:
	go mod tidy
	go mod vendor

modtest:
	go vet ./...
	go test -mod=vendor -v -race -cover -timeout 180s ./...

modbench:
	go test -mod=vendor -v -bench=. -benchmem -run=none ./...
