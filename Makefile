SOURCES = $(sort $(dir $(wildcard ./api/*/)))

.ONESHELL:

.PHONY: build-go
build-go:
	go build -v ./cmd/app

.ONESHELL:
.PHONY: generate-pb
generate-pb:
	for SOURCE in $(SOURCES); do \
  		  echo $$SOURCE; \
  		  protoc -I=./api --go_out=plugins=grpc:./extra $$SOURCE*.proto; \
  	done;

.PHONY: evans
evans:
	evans api/$(name).proto -p $(port)