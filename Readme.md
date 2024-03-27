# Homeworker
## Запуск
Чтобы запустить nats сервер нужно выполнить это в консоли:
```bash
docker pull nats
docker run -p 4222:4222 -d -ti nats
```
Далее для запуска необходимо в каком-то порядке выполнить следующее
```bash
go run server/microservices/gatewayMicroservice/cmd/gateway/main.go
go run server/microservices/auth_microservice/cmd/main.go
```
Далее можно слать Post запросы на localhost:8080