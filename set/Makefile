
.PHONY: all help tidy clean build run deploy quality test test/cover test/bench upgradeable
all: help

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /' 

# ==================================================================================== #
# Build/Run/Deploy
# ==================================================================================== #

## tidy: tidy modfiles and format .go files
tidy:
	@go mod tidy -v
	@go fmt ./...
## clean: clean build directory
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

## build: build binary
build:
	@echo "Building..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/$(BINARY_NAME)/main.go

## run: runs the binary
run:
	@./$(BUILD_DIR)/$(BINARY_NAME)

## deploy: builds and runs the binary
deploy: tidy clean build run

## docker: builds and runs the binary in a docker container
docker:
	docker build -t $(DOCKER) .
	docker run -p 3000:3000 $(DOCKER)

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: run quality control checks
audit: test test/cover test/bench upgradeable
	@go mod tidy -diff
	@go mod verify
	@test -z "$(shell gofmt -l .)" 
	@go vet ./...

## test: run all tests
test:
	@echo "Running unit tests..."
	@CGO_ENABLED=0 go test -v -buildvcs ./...


## test/cover: run all tests and display coverage
test/cover:
	@echo "Running unit tests with coverage..."
	@CGO_ENABLED=0 go test -v -buildvcs -coverprofile=/tmp/coverage.out ./...
	@go tool cover -html=/tmp/coverage.out

## test/bench: run all benchmarks
test/bench:
	@echo "Running benchmarks..."
	@CGO_ENABLED=0 go test ./... -v -run=^$$ -bench=. -benchtime=5s

## upgradeable: list direct dependencies that have upgrades available
upgradeable:
	@echo "Checking for upgrades..."
	@go list -u -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' -m all
