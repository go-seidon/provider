
default: help

.PHONY: help
help:
	@echo 'goseidon provider'
	@echo 'usage: make [target] ...'

.PHONY: install-tool
install-tool:
	go get -u github.com/golang/mock/gomock
	go get -u github.com/golang/mock/mockgen

.PHONY: install-dependency
install-dependency:
	go mod tidy
	go mod verify
	go mod vendor

.PHONY: clean-dependency
clean-dependency:
	rm -f go.sum
	rm -rf vendor
	go clean -modcache

.PHONY: test
test:
	go test ./.../ -p 1 -race -coverprofile coverage.out
	go tool cover -func coverage.out | grep ^total:

.PHONY: test-coverage
test-coverage:
	ginkgo -r -v -p -race --progress --randomize-all --randomize-suites -cover -coverprofile="coverage.out"

.PHONY: test-unit
test-unit:
	ginkgo -r -v -p -race --label-filter="unit" -cover -coverprofile="coverage.out"

.PHONY: test-integration
test-integration:
	ginkgo -r -v -p -race --label-filter="integration" -cover -coverprofile="coverage.out"

.PHONY: test-watch-unit
test-watch-unit:
	ginkgo watch -r -v -p -race --trace --label-filter="unit"

.PHONY: test-watch-integration
test-watch-integration:
	ginkgo watch -r -v -p -race --trace --label-filter="integration"

.PHONY: generate-mock
generate-mock:
	mockgen -package=mock_config -source config/config.go -destination=config/mock/config_mock.go
	mockgen -package=mock_context -source context/context.go -destination=context/mock/context_mock.go
	mockgen -package=mock_datetime -source datetime/clock.go -destination=datetime/mock/clock_mock.go
	mockgen -package=mock_dbmongo -source db-mongo/client.go -destination=db-mongo/mock/client_mock.go
	mockgen -package=mock_encoding -source encoding/encoder.go -destination=encoding/mock/encoder_mock.go
	mockgen -package=mock_hashing -source hashing/hasher.go -destination=hashing/mock/hasher_mock.go
	mockgen -package=mock_identifier -source identifier/id.go -destination=identifier/mock/id_mock.go
	mockgen -package=mock_io -source io/io.go -destination=io/mock/io_mock.go
	mockgen -package=mock_logging -source logging/log.go -destination=logging/mock/log_mock.go
	mockgen -package=mock_mysql -source mysql/client.go -destination=mysql/mock/client_mock.go
	mockgen -package=mock_serialization -source serialization/serializer.go -destination=serialization/mock/serializer_mock.go
	mockgen -package=mock_validation -source validation/validator.go -destination=validation/mock/validator_mock.go
