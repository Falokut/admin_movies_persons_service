package events

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type personsEvents struct {
	eventsWriter *kafka.Writer
	logger       *logrus.Logger
}

func NewPersonsEvents(cfg KafkaConfig, logger *logrus.Logger) *personsEvents {
	w := &kafka.Writer{
		Addr:   kafka.TCP(cfg.Brokers...),
		Logger: logger,
	}
	w.AllowAutoTopicCreation = true
	return &personsEvents{eventsWriter: w, logger: logger}
}

const (
	personDeletedTopic = "person_deleted"
)

func (e *personsEvents) Shutdown() error {
	return e.eventsWriter.Close()
}

func (e *personsEvents) PersonDeleted(ctx context.Context, id int32) error {
	body, err := json.Marshal(personDeletedEvent{ID: id})
	if err != nil {
		e.logger.Fatal(err)
	}
	return e.eventsWriter.WriteMessages(ctx, kafka.Message{
		Topic: personDeletedTopic,
		Key:   []byte(fmt.Sprint("person_", id)),
		Value: body,
	})
}
