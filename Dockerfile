# Użyj oficjalnego obrazu Golang jako bazowego
FROM golang:1.23.5-alpine AS builder

# Ustaw katalog roboczy
WORKDIR /app

# Skopiuj pliki modułów i zależności
COPY go.mod .
COPY go.sum .

# Pobierz zależności
RUN go mod download

# Skopiuj resztę kodu aplikacji
COPY . .

# Zbuduj narzędzie do migracji
RUN CGO_ENABLED=0 GOOS=linux go build -o migrate ./cmd/migrate/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Użyj minimalnego obrazu Alpine do uruchomienia aplikacji
FROM alpine:latest

WORKDIR /app

# Skopiuj plik wykonywalny z etapu budowania
COPY --from=builder /app/migrate .
COPY --from=builder /app/main .

# Skopiuj migracje
COPY cmd/migrate/migrations /app/migrations
COPY assets /app/assets

# Uruchom narzędzie do migracji
CMD ["sh", "-c", "./migrate up && exec ./main"]
