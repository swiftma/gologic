all:
	protoc -I. --go_out=plugins=micro:. \
		proto/vessel/vessel.proto
	GOOS=linux GOARCH=amd64 go build -o vessel-service