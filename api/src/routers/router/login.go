package router

import "net/http"

var rotaLogin = Rota{
	Uri:    "/login",
	Metodo: "POST",
	Funcao: func(w http.ResponseWriter, r *http.Request) {

	},
	RequerAutenticacao: false,
}
