postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1805 -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root drr_bank

dropdb:
	docker exec -it postgres16 dropdb drr_bank

migrateup:
	migrate -path db/migration -database "postgres://root:1805@localhost:5432/drr_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:1805@localhost:5432/drr_bank?sslmode=disable" -verbose down

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
