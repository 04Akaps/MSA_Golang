package mongo

import (
	"context"
	"fmt"

	"GO_MSA/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBLayer struct {
	Session *mongo.Client
}

func NewMongoSession(ctxMongo context.Context, envConfig config.Config) (*MongoDBLayer, error) {
	mongoconn := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.w5vs9re.mongodb.net/?retryWrites=true&w=majority", envConfig.MongoUserName, envConfig.MongoPassword))
	mongoClient, err := mongo.Connect(ctxMongo, mongoconn)

	return &MongoDBLayer{
		Session: mongoClient,
	}, err
}

func (mongoLayer MongoDBLayer) GetCollection(db, collection string) *mongo.Collection {
	return mongoLayer.Session.Database(db).Collection(collection)
}
