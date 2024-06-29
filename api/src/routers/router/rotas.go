package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"petPorject/api/src/middleware"
)

// Rota - representa todas as rotas da API
type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

// ConfigRotas - configura todas as rotas da API
func ConfigRotas(r *mux.Router) *mux.Router {
	rotas := rotasUsers
	rotas = append(rotas, rotaLogin)
	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			r.HandleFunc(rota.Uri, middleware.Logger(
				middleware.Authenticate(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.Uri, middleware.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	return r
}
