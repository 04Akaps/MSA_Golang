package services

import (
	"GO_MSA/models"

	"gopkg.in/mgo.v2/bson"
)

type Event interface {
	AddEvent(*models.EventModel) (bson.ObjectId, error)
	FindEvent(string) (*models.EventModel, error)
	FindEventByName(string) (*models.EventModel, error)
	FindAllAvaliableEvents() (*[]models.EventModel, error)
}
