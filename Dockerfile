# Wybieramy obraz bazowy
FROM golang:1.22 AS builder

# Ustawiamy katalog roboczy
WORKDIR /app

# Kopiujemy pliki go.mod i go.sum
COPY go.mod go.sum ./
RUN go mod download

# Kopiujemy pozostałe pliki
COPY . .

# Budujemy aplikację
RUN CGO_ENABLED=0 GOOS=linux go build -o my_app .

# Tworzymy nowy obraz
FROM debian:bullseye-slim

# Ustawiamy katalog roboczy
WORKDIR /root/

# Instalujemy wymagane biblioteki
RUN apt-get update && apt-get install -y ca-certificates && apt-get clean

# Kopiujemy plik binarny
COPY --from=builder /app/my_app .

# Otwieramy port 8080
EXPOSE 8080

# Uruchamiamy aplikację
CMD ["./my_app"]
