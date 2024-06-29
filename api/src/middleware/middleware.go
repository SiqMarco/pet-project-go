package middleware

import (
	"fmt"
	"net/http"
)

// Logger - loga as requisições do usuário no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Authenticate(funcao func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return nil
}
