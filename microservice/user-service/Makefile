all:
	protoc -I. --go_out=plugins=micro:. \
		proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build -o user-service