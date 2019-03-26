
init:
	docker-compose up -d

run:
	go run main.go

migrate:
	go run migration/migrate.go

test:
	go test ./...

genmock:
	for repo in `ls domain/repository | grep .go`; do \
		mockgen -source domain/repository/$${repo} -destination adapter/mock/repository/$${repo}; \
	done
	for usecase in `ls usecase | grep .go`; do \
		mockgen -source usecase/$${usecase} -destination adapter/mock/usecase/$${usecase}; \
	done
