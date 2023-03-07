package services

import (
	"context"

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

func (ei *EventServiceImpl) AddEvent(event *models.EventModel) ([]byte, error) {
	s := ei.session.GetFreshSession()

	if !event.Id.Valid() {
		// id가 문제가 없는지 확인 -> 반드시 12 bytes를 소유하고 있을 떄 Valid가 동작
		event.Id = bson.NewObjectId()
	}

	if !event.Location.Id.Valid() {
		event.Location.Id = bson.NewObjectId()
	}

	defer s.Close()

	return []byte(event.Id), s.DB("myevents").C("events").Insert(event)
	// DB는 들어가는 인자에 맞는 database를 준다.
	// C는 컬렉션을 반환
	// Insert를 사용하여 데이터를 추가
}

func (ei *EventServiceImpl) FindEvent(bt []byte) (*models.EventModel, error) {
	return nil, nil
}

func (ei *EventServiceImpl) FindEventByName(name string) (*models.EventModel, error) {
	return nil, nil
}

func (ei *EventServiceImpl) FindAllAvaliableEvents() ([]*models.EventModel, error) {
	return nil, nil
}
