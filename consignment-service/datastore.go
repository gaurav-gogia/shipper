package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

var Session *mgo.Session

func CreateSession(host string) *mgo.Session {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	log.Println("MONGODB SUCCESSFULLY CONNECTED on consignment-service")

	return session
}
