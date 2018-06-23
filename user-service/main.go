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
	db, err := CreateConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&pb.User{})

	repo := UserRepository{db}
	tokenService := TokenService{&repo}

	srv := micro.NewService(
		micro.Name(sname),
		micro.Version(version),
	)

	srv.Init()
	pb.RegisterUserServiceHandler(srv.Server(), &service{&repo, &tokenService})

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
