## Define variables

# static variables
SOURCE=./...

# derived variables

# variables from shell

# if-else blocks

# exported variables
export GO111MODULE=on


## Define targets

modvendor:
	go mod tidy
	go mod vendor

test: modvendor
	go vet $(SOURCE)
	# https://github.com/golang/go/issues/41572
	# https://github.com/golang/go/issues/49138
	MallocNanoZone=0 go test -mod=vendor -race -cover -failfast -timeout 180s $(SOURCE)
