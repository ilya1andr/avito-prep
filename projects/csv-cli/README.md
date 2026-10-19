\# csv-cli



CLI-утилита на Go для чтения CSV-файлов.



\## Что делает



Принимает путь к CSV-файлу, парсит его и печатает содержимое в stdout.

## Требования

- Go 1.22 или выше
- Docker (для запуска PostgreSQL)


\## Установка



git clone https://github.com/ilya1andr/avito-prep.git

cd avito-prep/projects/csv-cli

make build



## Использование

1. Запустите базу данных в Docker:
   ```bash
   docker run --name pg -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:16

./csv-cli data.csv



Вывод:

name | age | city

Alice | 30 | Moscow



\## Тесты



go test -v ./…



\## Команды Make



\- `make build` — собрать бинарник

\- `make run` — запустить на data.csv

\- `make test` — тесты

\- `make clean` — удалить бинарник


## Переменные окружения

Программа использует следующие переменные окружения для подключения к базе данных (если они не заданы, используются значения по умолчанию):

- `DB_HOST` — хост базы данных (по умолчанию `localhost`)
- `DB_USER` — пользователь (по умолчанию `postgres`)
- `DB_PASSWORD` — пароль (по умолчанию `secret`)
- `DB_NAME` — имя базы данных (по умолчанию `postgres`)
