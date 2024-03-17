FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./microservices/gateway_microservice/cmd/gateway/main.go

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