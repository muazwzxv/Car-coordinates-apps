package vehicle

import (
	"coordinates-seeder/internal/pkg/errorHelper"
	"time"

	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

type RegisterVehicleRequest struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Brand     string `json:"brand"`
	BuildDate string `json:"build_date"`
}

type VehicleDomain struct {
	Name      string
	Type      string
	Brand     string
	BuildDate string

	LastLatitude  float64
	LastLongitude float64
}

var VehicleType = map[string]bool{
	"CAR":       true,
	"LORRY":     true,
	"MOTORBIKE": true,
	"VAN":       true,
}

func (v *RegisterVehicleRequest) ValidateRegisterRequest() []errorHelper.ErrorDetail {
	var errs []errorHelper.ErrorDetail
	if v.Name == "" {
		errs = append(errs, errorHelper.ErrMissingName)
	}

	if v.Brand == "" {
		errs = append(errs, errorHelper.ErrMissingBrand)
	}

	if v.Type == "" {
		errs = append(errs, errorHelper.ErrMissingType)
	}

	if _, ok := VehicleType[v.Type]; !ok {
		errs = append(errs, errorHelper.ErrInvalidType)
	}

	if v.BuildDate == "" {
		errs = append(errs, errorHelper.ErrMissingBuildDate)
	}

	return errs
}

func (v *VehicleDomain) Publish(
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
