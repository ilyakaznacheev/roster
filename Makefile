generate/swagger:
# 	@rm -rf ./internal/api/models ./internal/api/restapi #cleanup
	@swagger generate server -f swagger.yml -t internal/api --exclude-main -A roster

generate/mocks:
	@mockery -name="StorageManager" -dir=$PWD -output="mocks"

generate: generate/swagger generate/mocks