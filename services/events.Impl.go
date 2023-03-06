package services

import (
	"context"

	"GO_MSA/models"
)

type EventServiceImpl struct {
	ctx context.Context
}

func NewEventService(ctx context.Context) Event {
	return &EventServiceImpl{
		ctx: ctx,
	}
}

func (ei *EventServiceImpl) AddEvent(event *models.EventModel) ([]byte, error) {
	return nil, nil
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
