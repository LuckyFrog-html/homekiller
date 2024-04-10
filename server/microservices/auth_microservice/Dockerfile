FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux go build -a -o main ./microservices/auth_microservice/cmd/main.go

FROM alpine:latest

#RUN apk update && apk --no-cache add ca-certificates
#
#RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируйте исполняемый файл из билдера
COPY --from=builder /app/main .
COPY /config config

# Запустите приложение
CMD ["./main"]