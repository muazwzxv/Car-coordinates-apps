include .env

run:
	go run ./cmd

ensure-mod:
	go mod tidy
	go mod vendor

migrations.new:
	migrate create -ext sql -dir ./db/migrations -seq $(name)

migrations.up:
	migrate -path ./db/migrations -database "${DB_SOURCE}" -verbose up $(N_MIGRATION)

migrations.down:
	migrate -path ./db/migrations -database "${DB_SOURCE}" -verbose down $(N_MIGRATION)
