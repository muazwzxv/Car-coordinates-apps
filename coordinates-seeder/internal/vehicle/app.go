package vehicle

import (
	"context"
	"coordinates-seeder/internal/pkg/errorHelper"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Move struct {
	LatitudeIncrease bool
	LatitudeDecrease bool

	LongitudeIncrease bool
	LongitudeDecrease bool
}

type MoveStrategy map[int]Move

var (

	// Latitude moves up, longitude stays
	FirstMove = Move{
		LatitudeIncrease: true,
		LatitudeDecrease: false,

		LongitudeIncrease: false,
		LongitudeDecrease: false,
	}

	// Longtidue goes up, latitude stays
	SecondMove = Move{
		LatitudeIncrease: true,
		LatitudeDecrease: false,

		LongitudeIncrease: false,
		LongitudeDecrease: false,
	}

	// Latitude goes up, latitude goes down
	ThirdMove = Move{
		LatitudeIncrease: false,
		LatitudeDecrease: true,

		LongitudeIncrease: true,
		LongitudeDecrease: false,
	}

	// Latitude goes down, Longitude goes up
	FourthMove = Move{
		LatitudeIncrease: false,
		LatitudeDecrease: true,

		LongitudeIncrease: true,
		LongitudeDecrease: false,
	}

	Strategies = MoveStrategy{
		0: FirstMove,
		1: SecondMove,
		2: ThirdMove,
		3: FourthMove,
	}
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
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.SendStatus(http.StatusCreated)
}

func (t *TruckApp) StartSeeding(ctx *fiber.Ctx) error {
	// query 10 vehicle
	vehicles, err := t.repository.GetAllVehicle(ctx.UserContext())
	if err != nil {
		return err
	}

	go t.startBackgroundSeed(ctx.UserContext(), vehicles)

	return ctx.SendStatus(201)
}

func (t *TruckApp) startBackgroundSeed(ctx context.Context, vehicles []*VehicleDomain) {
	var wg sync.WaitGroup

	for _, vehicle := range vehicles {

		strategyCode := rand.Intn(4-0) + 0
		move, ok := Strategies[strategyCode]
		if !ok {
			log.Println("strategy does not exist")
		}

    wg.Add(1)
		go t.publishLocationEvents(move, vehicle)
	}

  wg.Wait()
}

// nolint:unused
func (t *TruckApp) publishLocationEvents(move Move, vehicle *VehicleDomain) {
	totalMove := 100

	for i := 0; i == totalMove; i++ {
		// Move based on strategy
		log.Println("Lmaoo")

		time.Sleep(500 * time.Millisecond)
	}
}
