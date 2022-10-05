run:
	go run main.go

docker-compose/up:
	docker compose up

docker-compose/down:
	docker compose down

.PHONY: psql/client
psql/client:
	docker compose exec db psql -U root -d example
