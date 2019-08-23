
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc -I. --go_out=plugins=micro:. proto/test/test.proto

.PHONY: build
build: proto

	go build -o test-srv main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t test-srv:latest
