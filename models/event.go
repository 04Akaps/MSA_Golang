package models

import "gopkg.in/mgo.v2/bson"

type EventModel struct {
	Id        bson.ObjectId `bson:"id"`
	Name      string
	Duration  int64
	StartDate int64
	EndDate   int64
	Location  Loccation
}

type Loccation struct {
	Id        bson.ObjectId `bson:"id"`
	Name      string
	Address   string
	Countrry  string
	OpenTime  int
	CloseTime int
	Halls     []Hall
}

type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location,omitempty"`
	// omitempty은 해당 필드가 nil, 0, false같은 zero value라면 해당 필드를 무시하고 저장하지 않는다는 의미
	Capacity int `json:"capacity"`
}
