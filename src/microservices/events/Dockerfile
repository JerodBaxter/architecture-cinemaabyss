FROM golang:1.23-alpine AS builder

WORKDIR /app

# Установка необходимых пакетов для Alpine
RUN apk add --no-cache \
    gcc \
    musl-dev \
    libc-dev \
    linux-headers

# Копирование файлов конфигурации Go
COPY go.mod ./
COPY go.sum ./

# Установка зависимостей
RUN go mod download

# Копирование исходного кода
COPY . .

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o events-service .

# Использование меньшего базового образа
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копирование бинарного файла
COPY --from=builder /app/events-service .

# Экспозиция порта
EXPOSE 8082

# Команда запуска
CMD ["./events-service"]
