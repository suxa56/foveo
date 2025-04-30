# Stage 1 — билдим бинарник
FROM golang:1.24.2-alpine AS builder

# Рабочая директория в контейнере
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Качаем зависимости
RUN go mod download

# Копируем остальной код
COPY . .

# Собираем бинарник
RUN go build -o app ./cmd/app

# Открываем порт
EXPOSE 8090

# Команда запуска
CMD ["./app"]
