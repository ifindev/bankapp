postgres: 
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

start-postgres:
	docker start postgres12

stop-postgres:
	docker stop postgres12

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root bank_app 

dropdb: 
	docker exec -it postgres12 dropdb bank_app 

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_app?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_app?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: 
	postgres start-postgres createdb migrateup migratedown dropdb stop-postgres sqlc test
