package vehicle

import (
	"coordinates-seeder/internal/pkg/errorHelper"
	"net/http"

	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type TruckApp struct {
	topic      string
	repository IVehicleRepository
	publisher  *kafka.Publisher
}

func NewVehicleApp(topicName string, db *sqlx.DB, publisher *kafka.Publisher) *TruckApp {
	return &TruckApp{
		topic:      topicName,
		repository: NewVehicleRepository(db),
		publisher:  publisher,
	}
}

func (t *TruckApp) RegisterVehicle(ctx *fiber.Ctx) error {
  var req RegisterVehicleRequest
  if err := ctx.BodyParser(&req); err != nil {
    return ctx.Status(http.StatusBadRequest).
      JSON(errorHelper.SimpleErrorResponse(errorHelper.ErrBadRequest))
  }

  errs := req.ValidateRegisterRequest()
  if len(errs) != 0 {
    return ctx.Status(http.StatusBadRequest).
      JSON(errorHelper.ApplicationError(errs))
  }

  err := t.repository.RegisterVehicleData(ctx.UserContext(), &req)
  if err != nil {
    return ctx.Status(http.StatusInternalServerError).
      JSON(errorHelper.SimpleErrorResponse(errorHelper.ErrInternalServer))
  }

	return ctx.SendStatus(http.StatusCreated)
}
