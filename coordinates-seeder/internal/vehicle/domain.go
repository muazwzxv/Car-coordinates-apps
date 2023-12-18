package vehicle

import "github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"

type Vehicle struct {
	Name      string
	Type      string
	Brand     string
	BuildDate string

  LastLatitude float64
  LastLongitude float64
}

func (v *Vehicle) Publish(publisher *kafka.Publisher) error {
  return nil
}
