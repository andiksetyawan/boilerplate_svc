APP_NAME=boilerplate_svc
#IMAGE_REGISTRY=docker.io/andiksetyawan

IMAGE_NAME=$(APP_NAME)
#IMAGE_NAME=$(IMAGE_REGISTRY)/$(APP_NAME)
IMAGE_TAG=latest
#IMAGE_TAG=$(shell git rev-parse --short HEAD)

.PHONY: build test mock-generator docker-build docker-compose-start docker-compose-stop docker-compose-purge docker-compose-clean

build:
	go build -o $(APP_NAME) cmd/rest/main.go

test:
	go test -coverprofile=coverage.out ./... && \
    go tool cover -func=coverage.out | \
    awk '!/mocks\//'

mock-generator:
	./mock_gen.sh

docker-build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .
	#docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(IMAGE_NAME):latest

docker-compose-start:
	docker compose up

docker-compose-stop:
	docker compose down

docker-compose-purge:
	docker compose -f docker-compose.yml down --volumes

docker-compose-clean:
	docker compose down --volumes