FROM golang:1.23.5-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o migrate ./cmd/migrate/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/migrate .
COPY --from=builder /app/main .

COPY cmd/migrate/migrations /app/migrations
COPY assets /app/assets
COPY .env /app

CMD ["sh", "-c", "./migrate up && exec ./main"]