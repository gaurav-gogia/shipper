build:		
	protoc --go_out=plugins=micro:. ./proto/user/user.proto
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-service .	
	docker build -t user-service .

run:
	docker run --name=user_service -d -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns user-service