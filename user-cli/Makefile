build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-cli .	
	docker build -t user-cli .

run:
	docker run --name=user_cli -e MICRO_REGISTRY=mdns user-cli