package test

import (
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type TestStruct struct {
	Id   bson.ObjectId `json:"id" bson:"id"`
	Name string        `json:"name" bson:"name"`
	Age  int64         `json:"age" bson:"age"`
}

func Test(t *testing.T) {
	testCtx := context.Background()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.NewClient(clientOptions) // 가상의 메모리 DB생성
	if err != nil {
		t.Fatal(err)
	}

	// defer func() {
	// 	if err = client.Disconnect(testCtx); err != nil {
	// 		t.Fatal(err)
	// 	}
	// }()

	testCollection := client.Database("test").Collection("test_collection")

	newTestData := &TestStruct{
		Id:   bson.NewObjectId(),
		Name: RandomName(),
		Age:  RandomAge(),
	}

	_, err = testCollection.InsertOne(testCtx, newTestData)

	if err != nil {
		t.Fatal(err)
	}

	var findTestData TestStruct

	err = testCollection.FindOne(testCtx, bson.M{"_id": newTestData.Id}).Decode(findTestData)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Insert Data's Id : %v, Name : %s, Age : %v \n", newTestData.Id, newTestData.Name, newTestData.Age)

	fmt.Printf("Finded Data's Id : %v, Name : %s, Age : %v \n", findTestData.Id, findTestData.Name, findTestData.Age)
}
