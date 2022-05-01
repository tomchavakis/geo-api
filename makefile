IMAGE_NAME=geo-api

help:
	cat makefile

fmt:
	go fmt ./...

lint:
	golangci-lint run --modules-download-mode=vendor --timeout=2m0s  -E gosec -E revive --exclude-use-default=false --build-tags integration

test-unit:
	go test ./...


## Build


build:
	CGO_ENABLED=0 GO111MODULE=on GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ./bin/geo-api ./cmd/geo-api/main.go

docker-build:
	DOCKER_BUILDKIT=1 docker build -f Dockerfile -t ${IMAGE_NAME} .


## Run


up:
	COMPOSE_PROJECT_NAME=geoapi docker-compose -f docker-compose.yml up

down:
	COMPOSE_PROJECT_NAME=geoapi docker-compose -f docker-compose.yml down

docker-run:
	COMPOSE_PROJECT_NAME=geoapi docker-compose -f docker-compose.yml up geo-api

run:
	GO111MODULE=on go run -mod=vendor cmd/geo-api/main.go
