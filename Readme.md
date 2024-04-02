# Homeworker
## Запуск
Чтобы запустить nats сервер нужно выполнить это в консоли, находясь в корне проекта:
```bash
docker-compose up -d
```
Перед запуском убедитесь, что создан файл конфигов.
```yaml
env: "local"
db_string: "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
http_server:
  address: "localhost:3002"
  timeout: "4s"
  idle_timeout: "60s"
database:
  host: "localhost"
  user: "postgres"
  password: "postgres"
  dbname: "gorm"
  port: 5432
nats:
  nats_address: "nats:4222"
```
Заполните его своими данными.