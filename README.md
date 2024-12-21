## Калькулятор на go

Это калькулятор на golang.

## Запуск

Нужно просто склонировать репозиторий (или перетащить файлы в директорию) и написать `go run "Путь до директории"`, и сервер запустится!

## Примеры
(Примеры протестированы в bash оболочке. Они могут не работать в powershell или cmd)

# Пример 1 (Обычное выражение)
Запрос: `curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*2"}'`
Ответ: `{"result": "6"}`

# Пример 2 (Сложное выражение)
Запрос: `curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "(6+8,2)*5,12-(5,971-8,3335)/5"}'`
Ответ: `{"result": "73,1765"}`

# Пример 3 (Ошибка с выражением)
Запрос: `curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2++2*2"}'`
Ответ: `{"error": "Expression is not valid"}`

# Пример 4 (Ошибка с json)
Запрос: `curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '<xml>Hold on, I am not a json, I am an XML! NOOO!!!</xml>'`
Ответ: `{"error": "Internal server error"}`