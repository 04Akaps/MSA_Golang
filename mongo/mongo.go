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

func NewMongoSession(ctxMongo context.Context, envConfig config.Config) (*MongoDBLayer, error) {
	mongoconn := options.Client().ApplyURI("mongodb+srv://hojin:12345@cluster0.w5vs9re.mongodb.net/?retryWrites=true&w=majority")

	mongoClient, err := mongo.Connect(ctxMongo, mongoconn)

	return &MongoDBLayer{
		Session: mongoClient,
	}, err
}

func (mongoLayer MongoDBLayer) GetCollection(db, collection string) *mongo.Collection {
	return mongoLayer.Session.Database(db).Collection(collection)
}
