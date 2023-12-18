package vehicle

import (
	"time"

	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

type Vehicle struct {
	Name      string
	Type      string
	Brand     string
	BuildDate string

	LastLatitude  float64
	LastLongitude float64
}

func (v *Vehicle) Publish(
	publisher *kafka.Publisher,
	msg *message.Message,
	topic string,
) error {
	count := 3
	for {
		if err := publisher.Publish(topic, msg); err != nil {
			count--
			if count == 0 {
				return err
			}
			time.Sleep(500 * time.Millisecond)
		} else {
			break
		}
	}

	return nil
}
