CREATE=migrate create -ext sql -dir resources/sqlite/migrations $(name)
create-migration:
	$(CREATE)

MIGRATE=migrate -source file://resources/sqlite/migrations -database sqlite3://resources/sqlite/sqlite.db
migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down


sqlc-generate:
	cd resources/sqlite && sqlc generate

seed:
	go run -tags="fts5" cmd/seed/main.go