# syntax=docker/dockerfile:1

FROM golang:1.21.1 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./. ./

RUN CGO_ENABLED=0 GOOS=linux go build -o yablonka-backend ./cmd/app/main.go

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage ./app/yablonka-backend ./yablonka-backend
COPY --from=build-stage ./app/internal/config/.env ./internal/config/.env
COPY --from=build-stage ./app/internal/config/config.yml ./internal/config/config.yml

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/yablonka-backend"]