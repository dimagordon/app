
build-app:
	go build -tags=jsoniter -o ./.bin/app ./cmd/auth_service/auth_service.go

run: build-app
	docker-compose -f develop/docker-compose.yml up api

down:
	docker-compose -f develop/docker-compose.yml down

build:
	docker-compose -f develop/docker-compose.yml build api

migrate:
	docker-compose -f develop/docker-compose.yml up migrate

migrate-down:
	docker-compose -f develop/docker-compose.yml up migrate-down
