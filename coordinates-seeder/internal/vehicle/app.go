package vehicle

import (
	"context"
	"coordinates-seeder/internal/pkg/errorHelper"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
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

  newCtx := context.Background()
	go t.startBackgroundSeed(newCtx, vehicles)

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
		log.Printf("Seeding for vehicle id: %d", vehicle.ID)

		wg.Add(1)
		go t.publishLocationEvents(&wg, move, vehicle)
	}

  // blocking
	wg.Wait()

	log.Println("Finish moving coordinates")

  log.Println("Storing the current Lat lon state into database")
  for _, vehicle := range vehicles {
    err := t.repository.UpdateLatLonState(ctx, vehicle)
    log.Println(err)
  }

  log.Println("Background job finished")

}

// nolint:unused
func (t *TruckApp) publishLocationEvents(wg *sync.WaitGroup, move Move, vehicle *VehicleDomain) {
	defer wg.Done()
	totalMove := 100

	latitudeNoChange := !move.LatitudeIncrease && !move.LatitudeDecrease
	longitudeNoChange := !move.LongitudeIncrease && !move.LongitudeDecrease

	for i := 0; i < totalMove; i++ {
		log.Printf("seeding attempt id: %d, move: %d", vehicle.ID, i)
		if !latitudeNoChange {
			if move.LatitudeDecrease {
				vehicle.LastLatitude -= 10.55
			}

			if move.LatitudeIncrease {
				vehicle.LastLatitude += 20.65
			}
		}

		if !longitudeNoChange {
			if move.LongitudeDecrease {
				vehicle.LastLongitude -= 16.43004
			}

			if move.LongitudeIncrease {
				vehicle.LastLongitude += 24.2223
			}
		}

		payloadBytes, _ := json.Marshal(vehicle)
		msg := message.NewMessage(watermill.NewUUID(), payloadBytes)

		for j := 0; j < 2; j++ {
			err := vehicle.Publish(t.publisher, msg, t.topic)
			if err != nil {
				continue
			}
			break
		}

		time.Sleep(200 * time.Millisecond)
	}
}
