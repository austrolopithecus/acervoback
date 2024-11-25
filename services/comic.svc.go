package services

import (
	"acervoback/models"
	"acervoback/models/requests"
	"acervoback/repository"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// URL base da Open Library API
const openLibraryAPI = "https://openlibrary.org/api/books"

// ComicService estrutura para o serviço de quadrinhos
type ComicService struct {
	repo repository.ComicRepo
	c    *resty.Client
}

// Função para garantir que um valor seja uma string e evitar falhas
func getStringFromMap(m map[string]interface{}, key string) string {
	value, ok := m[key]
	if !ok {
		return ""
	}
	return value.(string)
}

// Função para criar um quadrinho
func (s *ComicService) Create(userID string, body requests.NewComicRequest) (*models.Comic, error) {
	// Validação do ISBN
	if len(body.ISBN) != 10 && len(body.ISBN) != 13 {
		log.Error().Str("isbn", body.ISBN).Msg("ISBN inválido antes de consultar API")
		return nil, errors.New("ISBN inválido ou malformado")
	}

	// Inicializando comic com dados padrão
	comic := &models.Comic{
		ID:     uuid.New().String(),
		Title:  "Título não especificado",
		Author: "Autor não especificado",
		UserID: userID,
	}

	// Construir a URL para consultar a Open Library API
	url := fmt.Sprintf("%s?bibkeys=ISBN:%s&format=json&jscmd=data", openLibraryAPI, body.ISBN)
	resp, err := s.c.R().Get(url)
	if err != nil {
		log.Warn().Msg("Erro ao consultar Open Library API, continuando com dados padrões")
	} else {
		var result map[string]interface{}
		if err := json.Unmarshal(resp.Body(), &result); err != nil {
			log.Err(err).Msg("Erro ao processar a resposta da API")
			return nil, err
		}

		key := fmt.Sprintf("ISBN:%s", body.ISBN)
		item, ok := result[key].(map[string]interface{})
		if !ok || item == nil {
			log.Warn().Str("isbn", body.ISBN).Msg("Dados não encontrados na resposta da API")
			return nil, errors.New("dados não encontrados na Open Library API")
		}

		comic.Title = getStringFromMap(item, "title")

		if authors, ok := item["authors"].([]interface{}); ok && len(authors) > 0 {
			if author, ok := authors[0].(map[string]interface{}); ok {
				comic.Author = getStringFromMap(author, "name")
			}
		}

		if publishers, ok := item["publishers"].([]interface{}); ok && len(publishers) > 0 {
			comic.Publisher = getStringFromMap(publishers[0].(map[string]interface{}), "name")
		}

		if pages, ok := item["number_of_pages"].(float64); ok {
			comic.Pages = int(pages)
		}

		comic.Description = getStringFromMap(item, "description")

		comic.Edition = getStringFromMap(item, "edition_name")

		if genres, ok := item["subjects"].([]interface{}); ok && len(genres) > 0 {
			comic.Genre = getStringFromMap(genres[0].(map[string]interface{}), "name")
		}

		if cover, ok := item["cover"].(map[string]interface{}); ok {
			comic.CoverURL = getStringFromMap(cover, "medium")
		}
	}

	if err := s.repo.Create(comic); err != nil {
		log.Err(err).Msg("Erro ao salvar comic no banco de dados")
		return nil, err
	}

	return comic, nil
}

func (s *ComicService) GetComics(userID string) ([]models.Comic, error) {
	comics, err := s.repo.FindByOwner(userID)
	if err != nil {
		log.Err(err).Str("userID", userID).Msg("Erro ao buscar comics do usuário")
		return nil, err
	}

	log.Debug().Str("userID", userID).Msgf("Comics encontrados: %d", len(comics))
	return comics, nil
}
