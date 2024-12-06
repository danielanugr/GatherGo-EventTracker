package services

import "github.com/danielanugr/GatherGo-EventTracker/models"

type EventServiceImpl struct {
	// TODO: Add database connection or repository here
}

func NewEventService() EventService {
	return &EventServiceImpl{}
}

func (s *EventServiceImpl) CreateEvent(event *models.Event) error {
	// TODO: Implement create event logic
	return nil
}

func (s *EventServiceImpl) GetEventById(id *string) (*models.Event, error) {
	// TODO: Implement get event by id logic
	return nil, nil
}

func (s *EventServiceImpl) GetAll() (*[]models.Event, error) {
	// TODO: Implement get all events logic
	return nil, nil
}

func (s *EventServiceImpl) UpdateEvent(event *models.Event) error {
	// TODO: Implement update event logic
	return nil
}

func (s *EventServiceImpl) DeleteEvent(id *string) error {
	// TODO: Implement delete event logic
	return nil
}
