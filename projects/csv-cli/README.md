# csv-cli

CLI-утилита на Go для чтения CSV-файлов.

## Что делает
1. Читает CSV-файл и сохраняет данные в базу данных PostgreSQL.
2. Умеет делать SQL-запрос с агрегацией (`GROUP BY`) и выводить статистику (количество людей и средний возраст по городам).

## Требования

- Go 1.22 или выше
- Docker (для запуска PostgreSQL)


## Установка

```bash
git clone https://github.com/ilya1andr/avito-prep.git
cd avito-prep/projects/csv-cli
make build
```

## Использование

1. Запустите базу данных в Docker:
   ```bash
   docker run --name pg -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:16
   ```

2. Создайте таблицу (зайдите в psql через docker exec -it pg psql -U postgres):
   ```sql
   CREATE TABLE people (id SERIAL PRIMARY KEY, name TEXT, age INT, city TEXT);
   ```

3. Загрузка данных:
   ```bash
   go run . data.csv
   ```

4. Просмотр статистики
   ```bash
   go run . --stats
   ```


## Тесты
```bash
go test -v ./...
```

## Команды Make

- `make build` — собрать бинарник
- `make run` — запустить на `data.csv`
- `make test` — тесты
- `make clean` — удалить бинарник
- `make stats` — запустить вывод статистики

## Переменные окружения

Программа использует следующие переменные окружения для подключения к базе данных (если они не заданы, используются значения по умолчанию):

- `DB_HOST` — хост базы данных (по умолчанию `localhost`)
- `DB_USER` — пользователь (по умолчанию `postgres`)
- `DB_PASSWORD` — пароль (по умолчанию `secret`)
- `DB_NAME` — имя базы данных (по умолчанию `postgres`)
