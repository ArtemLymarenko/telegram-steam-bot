#Example: make generate-protos proto_path=games/games.proto
generate-protos:
	cd protos && protoc -I proto ./proto/$(proto_path) \
		--go_out=./gen/go/ \
		--go_opt=paths=source_relative \
		--go-grpc_out=./gen/go/ \
		--go-grpc_opt=paths=source_relative