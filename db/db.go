package db

import (
	"acervoback/models"
	"errors"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDB() error {
	// Pega URL do banco de dados das variáveis de ambiente
	url := os.Getenv("DATABASE_URL")

	// Verifica se a URL está vazia
	if url == "" {
		log.Error().Msg("DATABASE_URL está vazio")
		return errors.New("DATABASE_URL está vazio")
	}

	// Inicializa o banco de dados
	var err error
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("Erro ao conectar ao banco de dados")
		return err
	}

	// Cria tabelas necessárias
	err = DB.AutoMigrate(
		&models.User{},
		&models.Comic{},
		&models.Exchange{}, // Adicionando o modelo Exchange
	)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao criar tabelas")
		return err
	}

	log.Info().Msg("Banco de dados conectado e tabelas criadas com sucesso")
	return nil
}
