FROM golang:1.13.7-alpine as builder
WORKDIR /go/src/

# Copy vendor first and build for cache layers
ENV GO111MODULE=on
COPY vendor vendor
COPY go.mod .
COPY go.sum .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -mod=vendor -ldflags="-w -s" -v ./vendor/...

# Copy source codes and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -mod=vendor -ldflags="-w -s" -v ./...

# Copy ran binary
FROM m3ng9i/ran:latest as ran

# For minimal image
FROM alpine:3.10
RUN apk --no-cache add ca-certificates 
COPY --from=builder /go/bin/ran-wrapper .
COPY --from=ran /ran .

# Start up server
COPY web web
COPY local.toml .
COPY cert.pem .
COPY key.pem .
CMD ./ran-wrapper
