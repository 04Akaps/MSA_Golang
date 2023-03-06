package services

import "GO_MSA/models"

type Event interface {
	AddEvent(*models.EventModel) ([]byte, error)
	FindEvent([]byte) (*models.EventModel, error)
	FindEventByName(string) (*models.EventModel, error)
	FindAllAvaliableEvents() ([]*models.EventModel, error)
}
