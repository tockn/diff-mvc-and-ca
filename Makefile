
init:
	docker-compose up -d

run:
	go run main.go

migrate:
	go run migration/migrate.go

test:
	go test ./...

