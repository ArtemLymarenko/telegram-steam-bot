CREATE=migrate create -ext sql -dir resources/migrations -seq $(name)
create-migration:
	$(CREATE)


DB_PATH=sqlite3://resources/sqlite/sqlite.db?cache=shared

MIGRATE=migrate -source file://resources/migrations -database $(DB_PATH)
migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

