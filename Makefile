SOURCES = $(sort $(dir $(wildcard ./api/protofile/*/)))

.ONESHELL:

.PHONY: build-go
build-go:
	go build -v ./cmd/app

.PHONY: generate-pb
generate-pb:
	for SOURCE in $(SOURCES); do \
		echo $$SOURCE; \
 		protoc -I api/protofile --go_out=module=github.com/midaef/emmet-server/internal/api,plugins=grpc:internal/api $$SOURCE*.proto; \
	done; \

.PHONY: evans
evans:
	evans api/protofile/$(name).proto -p $(port)