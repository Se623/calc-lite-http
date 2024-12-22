# Калькулятор на go

Это калькулятор на golang. Он роасполагается на http сервере. Что-бы им воспользоваться, нужно послать POST запрос С JSON типа `'{"expression": "выражение"}'`, и он пришлёт ответ с JSON типа `{"result": "результат"}` и код 200, либо ошибку типа `{"error": "описание"}`. Калькулятор может выдать 2 кода ошибок:
* 422 - Если выражение не соответствуют требованиям приложения (Нелегальные символы, или выражение не решаемо.)
* 500 - Если в теле запроса есть ошибки (Запрос не оформлен по правилам JSON)

## Запуск

Нужно просто склонировать репозиторий (или перетащить файлы в директорию) и написать `go run "Путь до директории"`, и сервер запустится!\
Сервер распологается на порту 8080.

## Примеры
(Примеры протестированы в bash оболочке. Они могут не работать в powershell или cmd)

### Пример 1 (Обычное выражение)
Запрос: `curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*2"}'`\
Ответ: `{"result": "6"}`

### Пример 2 (Сложное выражение)
Запрос: `curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "(6+8.2)*5.12-(5.971-8.3335)/5"}'`\
Ответ: `{"result": "73.17649999999999"}`

### Пример 3 (Ошибка с выражением)
Запрос: `curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2++2*2"}'`\
Ответ: `{"error": "Expression is not valid"}`

### Пример 4 (Ошибка с json)
Запрос: `curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '<xml>Hold on, I am not a json, I am an XML! NOOO!!!</xml>'`\
Ответ: `{"error": "Internal server error"}`