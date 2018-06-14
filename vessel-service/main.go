package main

import (
	"fmt"
	"os"

	pb "github.com/DesmondANIMUS/shipper/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
)

const (
	sname            = "go.micro.srv.vessel"
	version          = "latest"
	connectionString = "localhost:27017"
	dbName           = "shippy"
	vesselCollection = "vessels"
)

func main() {
	srv := micro.NewService(
		micro.Name(sname),
		micro.Version(version),
	)

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = connectionString
	}

	Session = CreateSession(host)
	defer Session.Close()

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
