FROM golang:1.21-alpine as builder
RUN apk add --no-cache --update gcc musl-dev g++ make git gnutls gnutls-dev gnutls-c++ bash git

WORKDIR /src

ADD ./go.mod ./go.sum ./
RUN go mod download

COPY cmd cmd
COPY config config
COPY database database
COPY internal internal
COPY pkg pkg
COPY pb pb
COPY i18n i18n

COPY .env.example .env.example

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags musl -o /dist/server cmd/server/*.go
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags musl -o /dist/worker cmd/worker/*.go
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags musl -o /dist/migrate cmd/migrate/*.go

FROM alpine:latest

RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/*

COPY --from=builder /dist/server /app/bin/server
COPY --from=builder /dist/worker /app/bin/worker
COPY --from=builder /dist/migrate /app/bin/migrate
COPY --from=builder /src/i18n /app/bin/i18n
COPY --from=builder /src/.env.example /app/bin/.env
COPY --from=builder /src/database /app/bin/database

WORKDIR /app/bin
EXPOSE 9000


