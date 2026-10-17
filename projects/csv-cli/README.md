\# csv-cli



CLI-утилита на Go для чтения CSV-файлов.



\## Что делает



Принимает путь к CSV-файлу, парсит его и печатает содержимое в stdout.

## Требования



\- Go 1.22 или выше


\## Установка



git clone https://github.com/ilya1andr/avito-prep.git

cd avito-prep/projects/csv-cli

make build



\## Использование



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

