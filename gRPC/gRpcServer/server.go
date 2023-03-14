package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"GO_MSA/config"
	m "GO_MSA/mongo"
	"GO_MSA/personpb"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var collection *mongo.Collection

type server struct {
	personpb.PersonServiceServer
}

func ErrHandling(funcName string, err error) {
	if err != nil {
		log.Fatal(funcName, " : Error is ocured : ", err)
	}
}

type Person struct {
	Name         string                  `bson:"name"`
	Age          int32                   `bson:"age"`
	PhoneNumbers []*personpb.PhoneNumber `bson:"phone_numbers"`
	LastUpdated  *timestamppb.Timestamp  `bson:"last_updated"`
}

func (*server) CreatePerson(ctx context.Context, req *personpb.CreatePersonRequest) (*personpb.CreatePersonResponse, error) {
	fmt.Println("Create New Person")

	person := req.GetPerson()

	data := Person{
		Name:         person.Name,
		Age:          person.Age,
		PhoneNumbers: person.PhoneNumbers,
		LastUpdated:  person.LastUpdated,
	}

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	// oid, ok := res.InsertedID.(primitive.ObjectID)

	// if !ok {
	// 	return nil, status.Errorf(
	// 		codes.Internal,
	// 		fmt.Sprintf("Cannot convert to OID"),
	// 	)
	// }

	return &personpb.CreatePersonResponse{
		Person: &personpb.Person{
			Name:         person.Name,
			Age:          person.Age,
			PhoneNumbers: person.PhoneNumbers,
			LastUpdated:  person.LastUpdated,
		},
	}, nil
}

// gRPC 서버 까지 root main.go에 작성하고 싶지는 않아서 따로 분리,
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println(" ---------- Connect TO MongoDB ---------- ")

	envConfig := config.LoadConfig("../../")

	mongoCtx := context.Background()

	mongoSession, err := m.NewMongoSession(mongoCtx, envConfig)
	ErrHandling("NewMongoSession", err)

	defer mongoSession.Session.Disconnect(mongoCtx)

	collection = mongoSession.Session.Database("gRPC").Collection("Person")
	fmt.Println("---------- gRPC Start ----------")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	ErrHandling("net.Listen", err)

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)

	personpb.RegisterPersonServiceServer(s, &server{})
	reflection.Register(s)

	go func() {
		fmt.Println("Starting gRPC Server")
		if err := s.Serve(lis); err != nil {
			ErrHandling("Starting gRPC Server", err)
		}
	}()
}
