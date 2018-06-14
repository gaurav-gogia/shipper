package main

import (
	pb "github.com/DesmondANIMUS/shipper/vessel-service/proto/vessel"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func FindVessel(spec *pb.Specification, session *mgo.Session) (*pb.Vessel, error) {
	var vessel *pb.Vessel
	c := session.DB(dbName).C(vesselCollection)

	query := bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}

	c.Find(query).One(&vessel)
	return vessel, nil
}

func CreateVessel(vessel *pb.Vessel, session *mgo.Session) error {
	c := session.DB(dbName).C(vesselCollection)
	return c.Insert(vessel)
}
