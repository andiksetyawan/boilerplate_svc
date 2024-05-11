APP_NAME=boilerplate_svc
#IMAGE_REGISTRY=docker.io/andiksetyawan
#IMAGE_NAME=$(IMAGE_REGISTRY)/$(APP_NAME)
IMAGE_NAME=$(APP_NAME)
IMAGE_TAG=$(shell git rev-parse --short HEAD)

.PHONY:

build:
	go build -o $(APP_NAME) cmd/rest/main.go

test:
	go test -coverprofile=coverage.out ./... && \
    go tool cover -func=coverage.out | \
    awk '!/mocks\//'

mock-generator:
	./mock_gen.sh

container-start:
	docker compose up

container-stop:
	docker compose down

container-purge:
	docker compose -f docker-compose.yml down --volumes

container-clean:
	docker compose down --volumes