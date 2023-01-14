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

protoc:
	rm -f pkg/pb/*.go
	protoc --proto_path=proto --go_out=pkg/pb --go_opt=paths=source_relative \
	--go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl

.PHONY: postgres, createdb, dropdb, migrateup, migratedown, evans