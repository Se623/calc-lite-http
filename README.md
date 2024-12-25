# Калькулятор на go

Это калькулятор на golang. Он роасполагается на http сервере. Что-бы им воспользоваться, нужно послать POST запрос С JSON типа `'{"expression": "выражение"}'`, и он пришлёт ответ с JSON типа `{"result": "результат"}` и код 200, либо ошибку типа `{"error": "описание"}`. Калькулятор может выдать 2 кода ошибок:
* 422 - Если выражение не соответствуют требованиям приложения (Нелегальные символы, или выражение не решаемо.)
* 500 - Если в теле запроса есть ошибки (Запрос не оформлен по правилам JSON)

## Запуск

1. Cклонировать репозиторий (Нужна программа git)
```bash
git clone https://github.com/Se623/calc-lite-http
```
2. Перейти в директорию программы
```bash
cd ./calc-lite-http
```
3. Запустить калькулятор
```bash
go run ./cmd
```

Сервер распологается на порту 8080.

## Примеры

### Пример 1 (Обычное выражение)
Запрос:\
Bash(Linux): `curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*2"}'`\
Cmd: `curl --location "localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{\"expression\": \"2+2*2\"}"`\

Ответ: `{"result": "6"}`

### Пример 2 (Сложное выражение)
Запрос:\
Bash(Linux): `curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "(6+8.2)*5.12-(5.971-8.3335)/5"}'`\
Cmd: `curl --location "localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{\"expression\": \"(6+8.2)*5.12-(5.971-8.3335)/5\"}"`

Ответ: `{"result": "73.17649999999999"}`

### Пример 3 (Ошибка с выражением)
Запрос:\
Bash(Linux): `curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2++2*2"}'`\
Cmd: `curl --location "localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{\"expression\": \"2++2*2\"}"`

Ответ: `{"error": "Expression is not valid"}`

### Пример 4 (Ошибка с json)
Запрос:\
Bash(Linux): `curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '<xml>Hold up, I am not a json, I am an XML! NOOO!!!</xml>'`\
Cmd: `curl --location "localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "Thankfully, i am not XML"`

Ответ: `{"error": "Internal server error"}`

## Тесты

Тесты находятся в файле `./cmd/main_test.go`, чтобы из запустить, нужно ввести команду (в папке calc-lite-http):
```bash
go test ./cmd
```