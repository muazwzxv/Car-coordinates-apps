package kafka

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
)

type KafkaPublisher struct {
	Publisher *kafka.Publisher
}

func New() (*KafkaPublisher, error) {
	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers: []string{
				"localhost:9092",
			},
			OTELEnabled: true,
			Marshaler:   kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(true, true),
	)

  if err != nil {
    return nil, err
  }

	return &KafkaPublisher{
    Publisher: publisher,
  }, nil
}
