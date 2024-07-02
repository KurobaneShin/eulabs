dsn = "root:root@tcp(localhost:3306)/db" 
migrationPath = "db/migrations" 

run: build
	@./bin/eulabs

build:
	@go build -o bin/eulabs .

db-status:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(dsn) goose -dir=$(migrationPath) status

db-up:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(dsn) goose -dir=$(migrationPath) up

db-reset:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(dsn) goose -dir=$(migrationPath) reset

migration-new:
	@GOOSE_MIGRATION_DIR=$(migrationPath) goose create $(filter-out $@,$(MAKECMDGOALS))

db-seed:
	go run ./cmd/seed/main.go
