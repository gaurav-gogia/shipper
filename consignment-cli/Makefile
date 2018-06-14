build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o consignment-cli .
	docker build -t consignment-cli .

run:
	docker run --name=consignment_cli -e MICRO_REGISTRY=mdns consignment-cli