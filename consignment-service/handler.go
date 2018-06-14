package main

import (
	"golang.org/x/net/context"

	pb "github.com/DesmondANIMUS/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/DesmondANIMUS/shipper/vessel-service/proto/vessel"
)

type handler struct {
	vesselClient vesselProto.VesselServiceClient
}

func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	var spec vesselProto.Specification
	spec.Capacity = int32(len(req.Containers))
	spec.MaxWeight = req.Weight

	vesselRes, err := s.vesselClient.FindAvailable(context.Background(), &spec)
	if err != nil {
		return err
	}
	req.VesselId = vesselRes.Vessel.Id

	session := Session.Clone()
	defer session.Close()

	if err := Create(req, session); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req

	return nil
}

func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	session := Session.Clone()
	defer session.Close()

	cons, err := GetAll(session)
	if err != nil {
		return err
	}

	res.Consignments = cons

	return nil
}
