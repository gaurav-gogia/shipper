build:
	protoc --go_out=plugins=micro:. ./proto/consignment/consignment.proto
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o consignment-service .	
	docker build -t consignment-service .

run:	
	docker run --name=consignment_service -d -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-service