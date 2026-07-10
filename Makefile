go ?= 1.24

.PHONY: up build run test coverage

help:
	@echo "Available options:"
	@echo "    make up"
	@echo "    make up go=$(go)"
	@echo "    make test"
	@echo "    make coverage"

up:
	docker rm -f manifesting > /dev/null 2>&1
	docker build . -t manifesting-image --build-arg GO_VERSION=$(go)
	docker run --interactive --detach --volume $(shell pwd):/app --name manifesting manifesting-image

build:
	docker exec -ti manifesting go build -v -o bin/manifesting .

test:
	docker exec -ti manifesting go test ./...

coverage:
	docker exec -ti manifesting go test ./... -coverprofile=tests/coverage.txt
	docker exec -ti manifesting go tool cover -html=tests/coverage.txt -o=tests/coverage.html
	sudo chown $$USER tests/coverage.html
	firefox tests/coverage.html
