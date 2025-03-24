DEFAULT_TARGET: build
DB_URL= postgres://postgres:log123@localhost:9696/clockbakers

.PHONY: fmt vet build run

fmt:
	@go fmt ./...

vet: fmt
	@go vet ./...

build: vet
	@go build -o ./bin/server.exe

run:
	@./bin/server


# up migration
up:
	@cd sql/schema && goose postgres $(DB_URL) up

# down migration
down:
	@cd sql/schema && goose postgres $(DB_URL) down


# sqlc queries gen
gen:
	@sqlc generate


# swagger
swagger:
	@swagger generate spec -o ./handlers/swagger.yaml --scan-models
