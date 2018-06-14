package main

import (
	"log"

	pb "github.com/DesmondANIMUS/shipper/user-service/proto/user"
	micro "github.com/micro/go-micro"
)

const (
	sname   = "go.micro.srv.user"
	version = "latest"
)

func main() {
	var repo Repository
	srv := micro.NewService(
		micro.Name(sname),
		micro.Version(version),
	)

	srv.Init()
	pb.RegisterUserServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
