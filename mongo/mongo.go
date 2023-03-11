package mongo

import (
	"context"

	"GO_MSA/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBLayer struct {
	Session *mongo.Client
}

func NewMongoSession(envConfig config.Config) (*MongoDBLayer, error) {
	ctxMongo := context.Background()

	// fmt.Sprintf("mongodb+srv://%s:%s@msggo.wbwdsv8.mongodb.net/?retryWrites=true&w=majority", envConfig.MongoUserName, envConfig.MongoPassword)
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")

	mongoClient, err := mongo.Connect(context.Background(), mongoconn)
	if err != nil {
		return nil, err
	}

	err = mongoClient.Ping(ctxMongo, nil)

	if err != nil {
		return nil, err
	}

	return &MongoDBLayer{
		Session: mongoClient,
	}, err
}

func (mongoLayer MongoDBLayer) GetCollection(db, collection string) *mongo.Collection {
	return mongoLayer.Session.Database(db).Collection(collection)
}
