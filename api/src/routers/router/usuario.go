package router

import (
	"net/http"
	"petPorject/api/src/controller"
)

var rotasUsers = []Rota{
	{
		Uri:                "/usuario",
		Metodo:             http.MethodPost,
		Funcao:             controller.CreateUser,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuario",
		Metodo:             http.MethodGet,
		Funcao:             controller.GetUsers,
		RequerAutenticacao: false,
	},
}
