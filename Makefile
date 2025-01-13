#Example: make generate-protos proto_path=games/games.proto
generate-protos:
	cd protos && protoc -I proto ./proto/$(proto_path) \
		--go_out=./gen/go/ \
		--go_opt=paths=source_relative \
		--go-grpc_out=./gen/go/ \
		--go-grpc_opt=paths=source_relative


parser-create-migration:
	cd services/parser && make create-migration $(name)

parser-m-up:
	cd services/parser && make migrate-up

parser-m-down:
	cd services/parser && make migrate-down

parser-sqlc-generate:
	cd services/parser && make sqlc-generate