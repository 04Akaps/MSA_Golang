package grpcclient

import (
	"context"
	"fmt"
	"log"

	"GO_MSA/personpb"

	"google.golang.org/grpc"
)

func ErrHandling(funcName string, err error) {
	if err != nil {
		log.Fatal(funcName, " : Error is ocured : ", err)
	}
}

func main() {
	fmt.Println("Person gRPC Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)

	ErrHandling("grpc.Dial", err)

	defer cc.Close()

	c := personpb.NewPersonServiceClient(cc)

	fmt.Println(" ---------- Creating the New Person  ---------- ")

	newPhoneNumber := &personpb.PhoneNumber{
		Number: "010-1234-5038",
	}
	newPerson := &personpb.Person{
		Name: "my gRPC Test Name",
		Age:  27,
	}
	newPerson.PhoneNumbers = append(newPerson.PhoneNumbers, newPhoneNumber)

	createPersonResponse, err := c.CreatePerson(context.Background(), &personpb.CreatePersonRequest{Person: newPerson})

	ErrHandling("CreatePerson", err)

	fmt.Printf("Person is Created!! : &v", createPersonResponse)
}
