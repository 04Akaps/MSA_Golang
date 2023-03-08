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

const (
	DB    = "myevents"
	EVENT = "events"
)

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

	return []byte(event.Id), ei.session.GetCollection(s, DB, EVENT).Insert(event)
	// DB는 들어가는 인자에 맞는 database를 준다.
	// C는 컬렉션을 반환
	// Insert를 사용하여 데이터를 추가
}

func (ei *EventServiceImpl) FindEvent(id []byte) (*models.EventModel, error) {
	// []byte타입을 받는 이유는 단순합니다.
	// 우리가 만약 mysql을 사용한다면, bson.ObjectId라는 값으로 데이터를 찾는 것은 불가능 합니다.
	// -> string(id) 이런식으로 특정 index를 찾아야 겠죠??
	// 그러기 떄문에 후에 혹시 마이그레이션을 해야 하는 경우를 대비하여 []byte값으로 받고 있습니다.
	s := ei.session.GetFreshSession()

	defer s.Clone()

	e := &models.EventModel{}

	err := ei.session.GetCollection(s, DB, EVENT).FindId(bson.ObjectId(id)).One(e)
	// 특정 EVENT를 찾는 코드 또한 이와 같다.
	return e, err
}

func (ei *EventServiceImpl) FindEventByName(name string) (*models.EventModel, error) {
	s := ei.session.GetFreshSession()

	defer s.Close()

	e := &models.EventModel{}

	err := ei.session.GetCollection(s, DB, EVENT).Find(bson.M{"name": name}).One(e)
	return e, err
}

func (ei *EventServiceImpl) FindAllAvaliableEvents() (*[]models.EventModel, error) {
	s := ei.session.GetFreshSession()

	defer s.Close()

	e := &[]models.EventModel{}

	err := ei.session.GetCollection(s, DB, EVENT).Find(nil).All(e)

	return e, err
}
