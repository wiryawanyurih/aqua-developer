include .env
export

build:
	CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./main.go && env

run: build
	docker-compose up redis db

.DEFAULT_GOAL := run
