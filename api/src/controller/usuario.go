package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"petPorject/api/src/db"
	"petPorject/api/src/model"
	"petPorject/api/src/repository"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario model.Usuario
	if erro = json.Unmarshal(requestBody, &usuario); erro != nil {
		Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar(); erro != nil {
		Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		Error(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repository.NewUsersRepository(db)
	usuario.UsuarioId, erro = repository.Create(usuario)
	if erro != nil {
		Error(w, http.StatusInternalServerError, erro)
		return
	}
	JSON(w, http.StatusCreated, usuario)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

}
