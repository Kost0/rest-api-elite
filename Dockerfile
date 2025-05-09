FROM golang:1.23.2-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/main.go --parseInternal --parseDependency

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/rest_api ./cmd/main.go

FROM alpine:3.18

WORKDIR /app

RUN apk add --no-cache tzdata

COPY --from=builder /app/rest_api /app/
COPY --from=builder /app/pkg/templates ./pkg/templates/
COPY --from=builder /app/docs ./docs/

ENV DB_HOST="localhost"
ENV DB_PORT="5432"
ENV DB_USER="postgres"
ENV DB_PASSWORD=""
ENV DB_NAME="postgres"

EXPOSE 8080

ENTRYPOINT ["/app/rest_api"]