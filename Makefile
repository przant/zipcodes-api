PROJECTDIR=$$PWD

.PHONY: help
help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Compiles the cmd/zipcodes api server locally
	CGO_ENABLED=0 go build -o apisrv ./cmd/zipcodes

compile: ## Compile the api server for macos, linux, and windows in different archs
	GOOS=darwin  GOARCH=arm64 go build -o ./bin/darwin/apisrv-darwin-arm64  ./cmd/zipcodes
	GOOS=darwin  GOARCH=amd64 go build -o ./bin/darwin/apisrv-darwin-amd64  ./cmd/zipcodes
	GOOS=linux   GOARCH=arm   go build -o ./bin/linux/apisrv-linux-arm      ./cmd/zipcodes
	GOOS=linux   GOARCH=arm64 go build -o ./bin/linux/apisrv-linux-arm64    ./cmd/zipcodes
	GOOS=linux   GOARCH=386   go build -o ./bin/linux/apisrv-linux-386      ./cmd/zipcodes
	GOOS=linux   GOARCH=amd64 go build -o ./bin/linux/apisrv-linux-amd64    ./cmd/zipcodes
	GOOS=windows GOARCH=arm   go build -o ./bin/windows/apisrv-darwin-arm   ./cmd/zipcodes
	GOOS=windows GOARCH=arm64 go build -o ./bin/windows/apisrv-darwin-arm64 ./cmd/zipcodes
	GOOS=windows GOARCH=386   go build -o ./bin/windows/apisrv-darwin-386   ./cmd/zipcodes
	GOOS=windows GOARCH=amd64 go build -o ./bin/windows/apisrv-darwin-amd64 ./cmd/zipcodes


compose-localdb: ## Build the app with Docker Compose
	@docker compose -f docker-compose.local.yaml build

compose-mysqldb: ## Build the app with Docker Compose to use MySQL
	@docker compose -f docker-compose.mysql.yaml build

start-localdb: ## Starts the api server with local database in Docker
	@docker compose -f docker-compose.local.yaml up

start-mysqldb: ## Starts the api server with a MySQl database in Docker
	@docker compose -f docker-compose.mysql.yaml up

run: ## Runs a local compiled cmd/zipcodes api server
	./apisrv

clean: ## Removes the binaries, and services created
	@rm -f apisrv* 2>/dev/null
	@cd ${PROJECTDIR}/bin/darwin/ && rm -f * 2>/dev/null
	@cd ${PROJECTDIR}/bin/linux/  && rm -f * 2>/dev/null
	@cd ${PROJECTDIR}/bin/windows/ && rm -f * 2>/dev/null
	@docker compose -f docker-compose.local.yaml down
	@docker compose -f docker-compose.mysql.yaml down