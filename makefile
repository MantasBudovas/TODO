export

WEB_PORT ?= 8080

.PHONY: run init all dep build clean fmt lint test qa fix-git help docs wire wire-test

run: ## Run the web app
	go run cmd/web/main.go --port=${WEB_PORT}

init: fix-git dep run ## Initialize, build and run the web app

dep: ## Download app dependencies
	go mod tidy
	go mod vendor

build:
	docker-compose up --build

generate-mocks:
	go tool mockery

coverage: test ## Show test coverage info in the browser
	go tool cover -html .testCoverage.txt

