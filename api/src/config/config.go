package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	// StingDbConnection - string de conexão com o banco de dados
	StingDbConnection = ""

	// Port - porta onde a API vai estar rodando
	Port = 0

	// SecretKey - chave para assinar o token
	SecretKey []byte
)

// Config - struct de configuração
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		fmt.Println("Erro ao carregar variáveis de ambiente")
		log.Fatal(erro)
	}

	// strconv.Atoi converte uma string para um inteiro
	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 9000
	}

	StingDbConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
