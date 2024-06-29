package model

import (
	"errors"
	"fmt"
	"time"
)

type Usuario struct {
	UsuarioId    uint64    `json:"usuario_id"`
	Nome         string    `json:"nome"`
	Cpf          string    `json:"cpf"`
	Email        string    `json:"email"`
	Senha        string    `json:"senha"`
	DataCadastro time.Time `json:"data_cadastro"`
	UsuarioTipo  string    `json:"usuario_tipo,omitempty"`
}

func (usuario *Usuario) Preparar() error {
	var message = "O campo %s é obrigatório e não pode estar em branco"
	if usuario.Nome == "" {
		return errors.New(fmt.Sprintf(message, "nome"))
	}
	if usuario.Cpf == "" {
		return errors.New(fmt.Sprintf(message, "cpf"))
	}
	if usuario.Email == "" {
		return errors.New(fmt.Sprintf(message, "email"))
	}
	if usuario.Senha == "" {
		return errors.New(fmt.Sprintf(message, "senha"))
	}
	return nil
}
