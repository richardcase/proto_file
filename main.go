package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/richardcase/proto_file/pkg/history/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	item := &pb.History{
		Id:         "id1234567",
		Provider:   "eks",
		Identity:   "saml",
		Alias:      "dev-cluster",
		Configfile: "~/.kube/config",
		Created:    timestamppb.Now(),
		Flags: []*pb.History_Flag{
			{
				Name:  "username",
				Value: "test@somewhere.com",
			},
			{
				Name:  "idp-endpoint",
				Value: "http://somehwre",
			},
		},
	}

	out, err := proto.Marshal(item)
	if err != nil {
		panic(err)
	}

	outJSON, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}

	filename := "./history.bin"
	filenameJSON := "./history.json"

	if err := ioutil.WriteFile(filename, out, os.ModePerm); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(filenameJSON, outJSON, os.ModePerm); err != nil {
		panic(err)
	}

	in, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	inItem := &pb.History{}
	if err := proto.Unmarshal(in, inItem); err != nil {
		panic(err)
	}

	if inItem.Id != item.Id {
		panic("no equal")
	}

}
