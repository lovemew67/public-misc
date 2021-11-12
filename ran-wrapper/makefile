## Define variables

# exported variables
export GO111MODULE=on


## Define targets

run:
	go run main.go

updatevendor:
	go mod tidy
	go mod vendor

dockerbuild: updatevendor
	docker build -t lovemew67/ran-wrapper:latest .

dockerpush: dockerbuild
	docker push lovemew67/ran-wrapper:latest
