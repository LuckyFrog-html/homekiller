# Homeworker
## Запуск
Чтобы запустить nats сервер нужно выполнить это в консоли, находясь в корне проекта:
```bash
docker-compose up -d
```
Перед запуском убедитесь, что создан файл конфигов.
```yaml
env: "local"
http_server:
  address: "localhost:3002"
  timeout: "4s"
  idle_timeout: "60s"
database:
  host: "localhost"
  user: "postgres"
  password: "postgres"
  dbname: "homekiller"
  port: 8000
```
Заполните его своими данными.