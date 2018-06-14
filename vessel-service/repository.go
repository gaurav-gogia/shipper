package main

import (
	pb "github.com/DesmondANIMUS/shipper/vessel-service/proto/vessel"
	mgo "gopkg.in/mgo.v2"
)

func FindVessel(session *mgo.Session) (*pb.Vessel, error) {
	var vessel *pb.Vessel
	c := session.DB(dbName).C(vesselCollection)
	c.Find(nil).One(&vessel)
	return vessel, nil
}

func CreateVessel(vessel *pb.Vessel, session *mgo.Session) error {
	c := session.DB(dbName).C(vesselCollection)
	return c.Insert(vessel)
}
