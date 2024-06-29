package routers

import (
	"github.com/gorilla/mux"
	"petPorject/api/src/routers/router"
)

// GetRouter - retorna um novo router com as rotas configuradas.
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	return router.ConfigRotas(r)
}
