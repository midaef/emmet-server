FROM golang:alpine as builder

ADD . /go/src/emmet-server
WORKDIR /go/src/emmet-server
RUN go mod download

COPY . ./

RUN chmod +x wait-for-postgres.sh

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /emmet-server ./cmd/app/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /emmet-server ./emmet-server
RUN mkdir ./configs
COPY ./configs/default-config.yaml ./configs

EXPOSE 65000

ENTRYPOINT ["./emmet-server"]
