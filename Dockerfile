# Build stage
FROM golang:1.23.2-alpine3.19 AS builder
WORKDIR /app
COPY . ./
COPY .env ./
COPY ./config ./config
COPY .env /app/.env

RUN go build -o backend main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/backend .
COPY --from=builder /app/.env .
COPY --from=builder /app/doc ./doc
COPY --from=builder /app/config ./config
ENTRYPOINT [ "/app/backend" ]
