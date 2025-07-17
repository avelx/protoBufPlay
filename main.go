package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	pb "protoBufPlay/github.com/protocolbuffers/protobuf/examples/go/tutorialpb"
)

func main() {

	person := &pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_MOBILE},
		},
	}

	fmt.Printf("Person, %s!\n", person)

	out, _ := proto.Marshal(person)

	if err := ioutil.WriteFile("/Users/pavel/devcore/GoLang/protoBufPlay/data/person.dat", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

}
