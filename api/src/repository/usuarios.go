package repository

import (
	"database/sql"
	"fmt"
	"log"
	"petPorject/api/src/model"
)

type Usuario struct {
	db *sql.DB
}

// NewUsersRepository - cria um novo repositório de usuários
func NewUsersRepository(db *sql.DB) *Usuario {
	return &Usuario{db}
}

func (repository Usuario) Create(usuario model.Usuario) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into usuarios (cpf, nome, email, senha, usuario_tipo) values (?, ?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resul, erro := statement.Exec(
		usuario.Cpf,
		usuario.Nome,
		usuario.Email,
		usuario.Senha,
		usuario.UsuarioTipo,
	)
	if erro != nil {
		return 0, erro
	}
	ultimoID, erro := resul.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoID), nil
}

func (repository Usuario) Search(cpf string) ([]model.Usuario, error) {
	cpf = fmt.Sprintf("%s%%%%", cpf)
	log.Println(cpf)
	lines, erro := repository.db.Query(
		"select id, cpf, nome, email, data_cadastro, usuario_tipo from usuarios where cpf LIKE ? ",
		cpf,
	)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var usuarios []model.Usuario

	for lines.Next() {
		var usuario model.Usuario
		if erro = lines.Scan(
			&usuario.UsuarioId,
			&usuario.Cpf,
			&usuario.Nome,
			&usuario.Email,
			&usuario.DataCadastro,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}
