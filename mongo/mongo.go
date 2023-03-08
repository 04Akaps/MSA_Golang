package mongo

import (
	"gopkg.in/mgo.v2"
)

type MongoDBLayer struct {
	Session *mgo.Session
}

func NewMongoSession(path string) (*MongoDBLayer, error) {
	session, err := mgo.Dial(path)

	return &MongoDBLayer{
		Session: session,
	}, err
}

func (mongoLayer MongoDBLayer) GetFreshSession() *mgo.Session {
	return mongoLayer.Session.Copy()
}

func (mongoLayer MongoDBLayer) GetCollection(session *mgo.Session, db, collection string) *mgo.Collection {
	return session.DB(db).C(collection)
}
