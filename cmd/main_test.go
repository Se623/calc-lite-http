package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
		status   int
	}{
		"Обычное выражение":   {input: `{"expression": "2+2*2"}`, expected: `{"result": "6"}`, status: http.StatusOK},
		"Сложное выражение":   {input: `{"expression": "(6+8.2)*5.12-(5.971-8.3335)/5"}`, expected: `{"result": "73.17649999999999"}`, status: http.StatusOK},
		"Ошибка с выражением": {input: `{"expression": "2++2*2"}`, expected: `{"error": "Expression is not valid"}`, status: http.StatusUnprocessableEntity},
		"Ошибка с json":       {input: `<xml>Hold on, I am not a json, I am an XML! NOOO!!!</xml>`, expected: `{"error": "Internal server error"}`, status: http.StatusInternalServerError},
		"Большие числа":       {input: `{"expression": "999999*999999999/5*88978884754+41246525624"}`, expected: `{"result": "1.779575913722733e+25"}`, status: http.StatusOK},
		"++Плюсы++":           {input: `{"expression": "1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1"}`, expected: `{"result": "23"}`, status: http.StatusOK},
		"(Скобки)":            {input: `{"expression": "((5/77-(4+8))+(6+(5*(5/7)-26/(23+5)))+(9*9/8))"}`, expected: `{"result": "6.832792207792208"}`, status: http.StatusOK},
		"Неправльные скобки":  {input: `{"expression": "((5+5-(8*6)(5-6)(58/7)(23*9))"}`, expected: `{"error": "Expression is not valid"}`, status: http.StatusUnprocessableEntity},
		"Неправильный ключ":   {input: `{"what": "2+2"}`, expected: `{"error": "Expression is not valid"}`, status: http.StatusUnprocessableEntity},
		"Без запроса":         {input: ``, expected: `{"error": "Internal server error"}`, status: http.StatusInternalServerError},
	}
	for name, test := range tests {
		req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", strings.NewReader(test.input))
		w := httptest.NewRecorder()

		rec := Receiver(answerHandler)
		rec.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)

		if res.StatusCode != test.status {
			t.Errorf("Error in test '%s': expected status code %d, got %dh", name, test.status, res.StatusCode)
		}
		if string(data) != test.expected {
			t.Errorf("Error in test '%s': expected output %s, got %sh", name, test.expected, string(data))
		}
		if err != nil {
			t.Errorf("Error in test '%s': %s", name, err)
		}
	}
}
