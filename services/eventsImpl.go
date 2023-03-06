package services

import "context"

type EventServiceImpl struct {
	ctx context.Context
}

func NewEventService(ctx context.Context) Event {
	return &EventServiceImpl{
		ctx: ctx,
	}
}
