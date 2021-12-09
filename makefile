fmt:
	go fmt ./...

lint:
	golangci-lint run --modules-download-mode=vendor --timeout=2m0s -E golint --exclude-use-default=false --build-tags integration

build:
	CGO_ENABLED=0 GO111MODULE=on GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ./bin/geo-api ./cmd/geo-api/main.go

docker-build:
	docker build -f infra/docker/dockerfile -t geo-api .

docker-run:
	docker run -d -p 3000:3000 -p 4000:4000 geo-api

test-unit:
	GO111MODULE=on go test -mod=vendor `go list -mod=vendor ./...` -race

run:
	GO111MODULE=on go run -mod=vendor cmd/geo-api/main.go