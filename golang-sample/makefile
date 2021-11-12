
## Define targets

updatevendor:
	go vet ./...
	go mod tidy
	go mod vendor

test: updatevendor	
	go test -mod=vendor -v -race -cover -timeout 180s ./...

bench: updatevendor
	go test -mod=vendor -v -bench=. -benchmem -run=none ./...

runmain: updatevendor
	go run -mod=vendor main.go
