postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root radium

dropdb:
	docker exec -it postgres12 dropdb radium

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/radium?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/radium?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v /home/zhs2si/projects/radium/jun-radium-go-gin:/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server

