package models

import "gopkg.in/mgo.v2/bson"

type EventModel struct {
	Id        bson.ObjectId `json:"id" bson:"id"`
	Name      string        `json:"name" bson:"name" binding:"required" `
	Duration  int64         `json:"duration" binding:"required" bson:"duration"`
	StartDate int64         `json:"start_date" binding:"required" bson:"start_date"`
	EndDate   int64         `json:"end_date" binding:"required" bson:"end_date"`
	Location  Loccation     `json:"location" binding:"required" bson:"location"`
}

type Loccation struct {
	Id        bson.ObjectId `json:"id" binding:"required" bson:"id"`
	Name      string        `json:"name" binding:"required" bson:"name"`
	Address   string        `json:"address" binding:"required" bson:"address"`
	Countrry  string        `json:"country" binding:"required" bson:"country"`
	OpenTime  int           `json:"open_time" binding:"required" bson:"open_time"`
	CloseTime int           `json:"close_time" binding:"required" bson:"close_time"`
	Halls     []Hall        `json:"halls" binding:"required" bson:"halls"`
}

type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Capacity int    `json:"capacity"`
}
