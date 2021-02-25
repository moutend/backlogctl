install:
	@echo "Install required tools"
	@go get -u -t github.com/volatiletech/sqlboiler/v4

migrate:
	@echo "Migrate models"
	@go run tools/migrate.go
	@sqlboiler sqlite3 --no-tests --output ./internal/models
	@rm temporary.db3

prepare:
	@echo "Create empty .db3 file"
	@cat migrations/*.up.sql | sqlite3 frequency.db3
