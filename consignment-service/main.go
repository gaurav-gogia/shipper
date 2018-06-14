package main

import (
	"fmt"

	"os"

	pb "github.com/DesmondANIMUS/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/DesmondANIMUS/shipper/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
)

const (
	sname                 = "go.micro.srv.consignment"
	vesselService         = "go.micro.srv.vessel"
	version               = "latest"
	connectionString      = "localhost:27017"
	dbName                = "shippy"
	consignmentCollection = "consignments"
)

func main() {
	srv := micro.NewService(
		micro.Name(sname),
		micro.Version(version),
	)
	vesselClient := vesselProto.NewVesselServiceClient(vesselService, srv.Client())
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = connectionString
	}

	Session = CreateSession(host)
	defer Session.Close()

	srv.Init()
	pb.RegisterShippingServiceHandler(srv.Server(), &handler{vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
