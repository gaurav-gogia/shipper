package main

import (
	"context"

	pb "github.com/DesmondANIMUS/shipper/vessel-service/proto/vessel"
)

type service struct{}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	session := Session.Clone()
	defer session.Close()

	vessel, err := FindVessel(req, session)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	session := Session.Clone()
	defer session.Close()

	if err := CreateVessel(req, session); err != nil {
		return err
	}

	res.Vessel = req
	res.Created = true
	return nil
}
