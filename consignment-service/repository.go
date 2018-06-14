package main

import (
	pb "github.com/DesmondANIMUS/shipper/consignment-service/proto/consignment"
	mgo "gopkg.in/mgo.v2"
)

func Create(consignment *pb.Consignment, session *mgo.Session) error {
	c := session.DB(dbName).C(consignmentCollection)
	return c.Insert(consignment)
}

func GetAll(session *mgo.Session) ([]*pb.Consignment, error) {
	var consignments []*pb.Consignment
	c := session.DB(dbName).C(consignmentCollection)
	err := c.Find(nil).All(&consignments)
	return consignments, err
}
