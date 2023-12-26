package events

import "context"

type KafkaConfig struct {
	Brokers []string
}

type personDeletedEvent struct {
	ID int32 `json:"person_id"`
}

type PersonsEventsMQ interface {
	PersonDeleted(ctx context.Context, id int32) error
}
