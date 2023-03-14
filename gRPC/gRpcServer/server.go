package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"GO_MSA/config"
	m "GO_MSA/mongo"
	"GO_MSA/personpb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/mgo.v2/bson"
)

var collection *mongo.Collection

type server struct {
	personpb.PersonServiceServer
}

// CreatePerson(context.Context, *CreatePersonRequest) (*CreatePersonResponse, error)
// ReadPerson(context.Context, *ReadPersonRequest) (*ReadPersonResponse, error)
// ListPerson(context.Context, *ListPersonRequest) (*ListPersonResponse, error)

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
	// 기존에 Mongo를 사용할 예정이였기 떄문에 ID필드도 정의내려야 했지만, 일단 ID필드는 따로 없기 떄문에 따로 만들지는 않음
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

func (*server) ListPerson(_ *personpb.ListPersonRequest, stream personpb.PersonService_ListPersonServer) error {
	fmt.Println("List Persons")

	listPersonCtx := context.Background()

	cur, err := collection.Find(listPersonCtx, primitive.D{}) // 모두 찾기
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	defer cur.Close(listPersonCtx)

	personList := []*Person{}

	for cur.Next(listPersonCtx) {
		data := &Person{}

		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)
		}

		// stream grpc.ServerStream
		// client에서 Msg를 전송해 주어야 한다.
		stream.SendMsg(&personpb.ListPersonResponse{Person: &personpb.Person{
			Name:         data.Name,
			Age:          data.Age,
			PhoneNumbers: data.PhoneNumbers,
			LastUpdated:  data.LastUpdated,
		}})

		personList = append(personList, data)
	}

	for _, person := range personList {
		fmt.Printf("Name: %s, Age: %d, Phone Numbers: %v, Last Updated: %s\n", person.Name, person.Age, person.PhoneNumbers, person.LastUpdated)
	}

	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	return nil
}

func (*server) ReadPerson(ctx context.Context, req *personpb.ReadPersonRequest) (*personpb.ReadPersonResponse, error) {
	fmt.Println("Read One Person")

	personName := req.GetPersonName()

	data := &Person{}
	filter := bson.M{"name": personName}

	response := collection.FindOne(ctx, filter)

	if err := response.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", err),
		)
	}

	return &personpb.ReadPersonResponse{
		Person: &personpb.Person{
			Name:         data.Name,
			Age:          data.Age,
			PhoneNumbers: data.PhoneNumbers,
			LastUpdated:  data.LastUpdated,
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
