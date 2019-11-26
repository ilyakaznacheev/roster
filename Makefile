generate/swagger:
	@swagger generate server -f swagger.yml -t internal --exclude-main -A roster

generate/mocks:
	@mockery -name="StorageManager" -dir=$PWD -output="mocks"

generate: generate/swagger generate/mocks