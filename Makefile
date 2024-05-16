IGNORE_DIR1=-path ./datastorage -prune -o
GOFMT_FILES?=$$(find . $(IGNORE_DIR1) -name '*.go' | grep -v vendor)
GOFMT := "goimports"

fmt: ## Run gofmt for all .go files
	@$(GOFMT) -w -d -e -l $(GOFMT_FILES)

DEPEND=\
	golang.org/x/tools/cmd/goimports \
	golang.org/x/tools/cmd/stringer \
	github.com/swaggo/swag/cmd/swag \
    github.com/githubnemo/CompileDaemon

depend: ## Install dependencies for dev
	@go get -v ./...
	@go get -v $(DEPEND)

lint: ## run linter
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.33.0 golangci-lint run -v

dev:
	scripts/run_local.sh

worker:
	scripts/run_worker_local.sh

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Need install migrate: https://www.geeksforgeeks.org/how-to-install-golang-migrate-on-ubuntu/
migrate-create:
	migrate create -ext sql -dir database/migrate -seq -digits 5 $(name) -verbose up

migrate:
	migrate -source file://database/migrate \
		-database "mysql://test:test@tcp(127.0.0.1:3306)/workflow" -verbose up

rollback:
	migrate -source file://database/migrate \
		-database "mysql://test:test@tcp(127.0.0.1:3306)/workflow" -verbose down

force:
	migrate -source file://database/migrate \
		-database "mysql://test:test@tcp(127.0.0.1:3306)/workflow" -verbose force $(version)

# Need install sqlc: https://docs.sqlc.dev/en/latest/overview/install.html
gen-sql:
	sqlc generate

gen-grpc:
	protoc --go_out=./pb/membership \
		--go_opt=paths=source_relative \
		--go-grpc_out=./pb/membership \
		--go-grpc_opt=paths=source_relative ./proto/membership.proto

swag:
	docker-compose up gen-docs
	sleep 3
	docker-compose up convert-docs

copy-docs:
	cat docs/swagger_v3.yaml | xclip -selection clipboard

coverage:
	go test -coverprofile=coverage.out ./internal/store/mysql/...

convert-coverage:
	go tool cover -html=coverage.out -o coverage.html

test-local:
	mkfifo /tmp/fifo-$$PPID
	grep -v 'no test files' </tmp/fifo-$$PPID & CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go test $$(go list ./... | grep -v /e2e/) >/tmp/fifo-$$PPID
	rm /tmp/fifo-$$PPID
