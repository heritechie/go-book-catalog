protoc:
	rm -f pkg/pb/*.go
	protoc --proto_path=proto --go_out=pkg/pb --go_opt=paths=source_relative \
	--go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

evans:
	evans --host localhost --port 9092 -r repl

.PHONY: protoc, evans