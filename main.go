package main

import (
	"google.golang.org/protobuf/proto"
	"log"
	"math/rand"
	"os"
	pb "protoBufPlay/github.com/protocolbuffers/protobuf/examples/go/tutorialpb"
)

func main() {

	// Create an iterator to produce 10M person records
	personItemProducer := func() <-chan pb.Person {
		const maxPersonsToBeCreated int = 10000000
		const bufferSize int = 1024
		results := make(chan pb.Person, bufferSize)
		// Some processing function
		go func() {
			defer close(results)
			// TODO: add random data
			person := pb.Person{
				Id:    rand.Int31n(10000000),
				Name:  "John Doe",
				Email: "jdoe@example.com",
				Phones: []*pb.Person_PhoneNumber{
					{Number: "555-4321", Type: pb.Person_MOBILE},
				},
			}
			for i := 0; i < maxPersonsToBeCreated; i++ {
				results <- person
			}
		}()
		return results
	}

	// consume persons records from a readonly channel: write all to a fille
	// TODO: implement batch writing
	personConsumer := func(persons <-chan pb.Person) {
		const fileName string = "/Users/pavel/devcore/GoLang/protoBufPlay/data/person.dat"
		file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		for nextPerson := range persons {
			personInBytes, _ := proto.Marshal(&nextPerson)
			_, err := file.Write(personInBytes)
			//err := os.WriteFile(fileName, personInBytes, 0644)
			if err != nil {
				log.Fatalln("Failed to write address book:", err)
			}
		}
	}

	// MAIN LOGIC
	persons := personItemProducer()
	personConsumer(persons)
}
