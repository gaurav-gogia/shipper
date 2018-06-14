build:
	protoc --go_out=plugins=micro:. ./proto/vessel/vessel.proto
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vessel-service .	
	docker build -t vessel-service .

run:
	docker run --name=vessel_service -d -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service