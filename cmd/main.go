package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	calc "github.com/Se623/calc-lite-http/internal"
)

type Exp struct {
	Expression string `json:"expression"`
}

var answer float64

func Receiver(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			decoder := json.NewDecoder(r.Body)
			var resp Exp
			err := decoder.Decode(&resp)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, `{"error": "Internal server error"}`)
				return
			}
			answer, err = calc.Calc(resp.Expression)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				fmt.Fprint(w, `{"error": "Expression is not valid"}`)
				return
			}
			next.ServeHTTP(w, r)
		}()
	})
}

func answerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"result": "%g"}`, answer)
}

func main() {
	http.HandleFunc("/api/v1/calculate", Receiver(answerHandler))
	http.ListenAndServe(":8080", nil)
}
