package models

type DatabaseModel interface {
	AddEvent(EventModel) ([]byte, error)
	FindEvent([]byte) (EventModel, error)
	FindEventByName(string) (EventModel, error)
	FindAllAvaliableEvents() ([]EventModel, error)
}
