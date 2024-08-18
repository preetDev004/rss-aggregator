FROM golang:1.22.6-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

ARG DB_URL

RUN goose -dir sql/schema  postgres "${DB_URL}" up

RUN sqlc generate

RUN go build -o rssagg

# multi stage build
FROM alpine

WORKDIR /app

COPY --from=builder /app/rssagg .

EXPOSE 8080

CMD ["./rssagg"]