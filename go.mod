module github.com/steam-tg-bot

go 1.23.3

replace github.com/steam-tg-bot => ./steam-tg-bot

require (
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/golang-migrate/migrate/v4 v4.18.1
	github.com/joho/godotenv v1.5.1
	github.com/mattn/go-sqlite3 v1.14.24
)

require (
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)

