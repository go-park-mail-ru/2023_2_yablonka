# syntax=docker/dockerfile:1

FROM golang:1.21.1 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./. ./

RUN CGO_ENABLED=0 GOOS=linux go build -o yablonka-backend ./cmd/app/main.go
RUN go install github.com/jackc/tern/v2@latest

FROM alpine:latest AS build-release-stage

RUN apk --no-cache add  \
        gcompat         \
        libstdc++       \
        bash

RUN addgroup --system nonroot
RUN adduser --system nonroot --ingroup nonroot

COPY --from=build-stage ./app/db ./db
COPY --from=build-stage ./go/bin/tern ./tern
COPY --from=build-stage ./app/yablonka-backend ./yablonka-backend
COPY --from=build-stage ./app/internal/config/.env ./internal/config/.env
COPY --from=build-stage ./app/internal/config/config.yml ./internal/config/config.yml

EXPOSE 8080
EXPOSE 5432

USER nonroot:nonroot