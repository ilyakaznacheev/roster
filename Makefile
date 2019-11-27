TEST_TIMEOUT := 30s

## generate/swagger: Generate swagger server
generate/swagger:
	@swagger generate server -f swagger.yml -t internal/api --exclude-main -A roster

## generate/mocks: Generate mocks
generate/mocks:
	@mockery -name="DatabaseService" -dir=internal/handlers -output=internal/handlers/mocks

## generate: Run full code generation
generate: generate/swagger generate/mocks

test:
	@go test ./... -covermode=count -timeout=$(TEST_TIMEOUT)

## help: Get makefile manual
help: Makefile
	@echo
	@echo Choose command to run:
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo