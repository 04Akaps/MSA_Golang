package services

import (
	"context"
	"log"

	"GO_MSA/models"
	"GO_MSA/mongo"

	"gopkg.in/mgo.v2/bson"
)

type EventServiceImpl struct {
	session *mongo.MongoDBLayer
	ctx     context.Context
}

func NewEventService(ctx context.Context, session *mongo.MongoDBLayer) Event {
	return &EventServiceImpl{
		ctx:     ctx,
		session: session,
	}
}

const (
	DB    = "myevents"
	EVENT = "events"
)

func (ei *EventServiceImpl) AddEvent(event *models.EventModel) (bson.ObjectId, error) {
	if !event.Id.Valid() {
		// id가 문제가 없는지 확인 -> 반드시 12 bytes를 소유하고 있을 떄 Valid가 동작
		event.Id = bson.NewObjectId()
	}

	if !event.Location.Id.Valid() {
		event.Location.Id = bson.NewObjectId()
	}

	_, err := ei.session.GetCollection(DB, EVENT).InsertOne(ei.ctx, event)

	return (event.Id), err
	// DB는 들어가는 인자에 맞는 database를 준다.
	// C는 컬렉션을 반환
	// Insert를 사용하여 데이터를 추가
}

func (ei *EventServiceImpl) FindEvent(id string) (*models.EventModel, error) {
	e := &models.EventModel{}
	err := ei.session.GetCollection(DB, EVENT).FindOne(ei.ctx, bson.M{"_id": id}).Decode(e)
	// 특정 EVENT를 찾는 코드 또한 이와 같다.
	return e, err
}

func (ei *EventServiceImpl) FindEventByName(name string) (*models.EventModel, error) {
	e := &models.EventModel{}
	err := ei.session.GetCollection(DB, EVENT).FindOne(ei.ctx, bson.M{"name": name}).Decode(e)
	return e, err
}

func (ei *EventServiceImpl) FindAllAvaliableEvents() ([]models.EventModel, error) {
	e := []models.EventModel{}

	result, err := ei.session.GetCollection(DB, EVENT).Find(ei.ctx, bson.M{})

	defer result.Close(ei.ctx)

	for result.Next(ei.ctx) {
		var element models.EventModel

		err := result.Decode(element)
		if err != nil {
			log.Fatal("Get All Events Error", err)
		}

		e = append(e, element)
	}

	return e, err
}
