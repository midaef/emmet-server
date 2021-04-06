.ONESHELL:

.PHONY: build-go

init-swagger:
	swag init -g cmd/app/main.go

build-go:
	go build -v ./cmd/app
