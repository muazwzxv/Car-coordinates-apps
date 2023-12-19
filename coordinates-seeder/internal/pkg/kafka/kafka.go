package kafka

import (
	"coordinates-seeder/internal/pkg/config"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
)

type KafkaPublisher struct {
	Publisher *kafka.Publisher
}

func NewAppPublisher(cfg config.Config) (*KafkaPublisher, error) {
	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers: []string{
				cfg.Broker,
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
