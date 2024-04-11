FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux go build -a -o main ./cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

# Копируйте исполняемый файл из билдера
COPY --from=builder /app/main .
COPY /config config

# Запустите приложение
CMD ["./main"]
