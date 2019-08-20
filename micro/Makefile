build:
	protoc -I consignment/ --go_out=plugins=grpc:consignment consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build -o consignment-service
	docker build -t shippy-service-consignment .

run:
	docker run -d -p 50051:50051 shippy-service-consignment
