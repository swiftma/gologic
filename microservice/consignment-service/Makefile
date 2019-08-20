all:
	protoc -I consignment/ --go_out=plugins=micro:consignment consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build -o consignment-service