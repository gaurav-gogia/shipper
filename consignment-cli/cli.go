package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"

	pb "github.com/DesmondANIMUS/shipper/consignment-service/proto/consignment"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

const (
	sname           = "go.micro.srv.consignment"
	defaultFilename = "consignment.json"
)

func main() {
	cmd.Init()

	client := pb.NewShippingServiceClient(sname, microclient.DefaultClient)
	consignment, e := parseFile(defaultFilename)
	err(e, "COULD NOT PARSE FILE")

	r, e := client.CreateConsignment(context.Background(), consignment)
	err(e, "rpc call to consignment-service CreateConsignment FAILED")
	log.Printf("Created: %v", r.Consignment)

	r, e = client.GetConsignments(context.Background(), &pb.GetRequest{})
	err(e, "rpc call to consignment-service GetConsignments FAILED")
	log.Printf("Created: %v", len(r.Consignments))
}

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &consignment)
	return consignment, err
}

func err(err error, s string) {
	if err != nil {
		panic(err.Error() + s)
	}
}
