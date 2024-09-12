package db

import (
	"acervoback/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDB() {
	// Pega url do banco de dados das variaveis de ambiente
	url := os.Getenv("DATABASE_URL")
	// verifica se a url esta vazia
	if url == "" {
		log.Fatal().Msg("DATABASE_URL esta vazio")
	}
	// Cria variavel de erro
	var err error
	// Codigo para inicializar o banco de dados
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Erro ao conectar ao banco de dados")
	}
	// Cria tabela de usuario
	err = DB.AutoMigrate(&models.User{}, &models.Comic{})
	if err != nil {
		log.Fatal().Err(err).Msg("Erro ao criar tabela de usuario")
		return
	}
	log.Info().Msg("Banco de dados conectado")
}
