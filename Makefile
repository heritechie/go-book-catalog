postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root book_catalog

dropdb:
	docker exec -it postgres15 dropdb --username=root book_catalog

migrateup:
	migrate -path internal/db/migrations -database "postgresql://root:secret@localhost:5432/book_catalog?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migrations -database "postgresql://root:secret@localhost:5432/book_catalog?sslmode=disable" -verbose down

.PHONY: postgres, createdb, dropdb, migrateup, migratedown