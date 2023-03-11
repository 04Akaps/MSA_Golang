package mongo

import (
	"context"
	"time"

	"GO_MSA/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type MongoDBLayer struct {
	Session *mongo.Client
}

func NewMongoSession(ctxMongo context.Context, envConfig config.Config) (*MongoDBLayer, error) {
	// fmt.Sprintf("mongodb+srv://%s:%s@msggo.wbwdsv8.mongodb.net/?retryWrites=true&w=majority", envConfig.MongoUserName, envConfig.MongoPassword)

	// 해당 옵션들에 대한 설명은
	// https://medium.com/@sdl182975/mongodb-mongo-driver-d920a438d626

	jmajority := writeconcern.New(writeconcern.J(true))                           // 로그 기록을 사용 여부
	wmajority := writeconcern.New(writeconcern.W(1))                              // 쓰기 확인을 위한 노드 수
	tmajority := writeconcern.New(writeconcern.WTimeout(1000 * time.Microsecond)) // 쓰기 제한 시간

	readConcert := readconcern.New(readconcern.Level("majority"))

	readProf, err := readpref.New(readpref.PrimaryPreferredMode)
	if err != nil {
		return nil, err
	}

	mongoconn := options.Client().
		ApplyURI("mongodb://localhost:27017").
		SetConnectTimeout(10 * time.Second).
		SetMaxPoolSize(50).SetMinPoolSize(5).
		SetWriteConcern(jmajority).
		SetWriteConcern(wmajority).
		SetWriteConcern(tmajority).
		SetReadConcern(readConcert).SetReadPreference(readProf)

	mongoClient, err := mongo.Connect(ctxMongo, mongoconn)
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
