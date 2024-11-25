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

const googleBooksAPI = "https://www.googleapis.com/books/v1/volumes"

type ComicService struct {
	repo   repository.ComicRepo
	c      *resty.Client
	apiKey string
}

func NewComicService(repo repository.ComicRepo, apiKey string) *ComicService {
	return &ComicService{
		repo:   repo,
		c:      resty.New(),
		apiKey: apiKey,
	}
}

func (s *ComicService) Create(userID string, body requests.NewComicRequest) (*models.Comic, error) {
	// Validação do ISBN
	if len(body.ISBN) != 10 && len(body.ISBN) != 13 {
		log.Error().Str("isbn", body.ISBN).Msg("ISBN inválido antes de consultar API")
		return nil, errors.New("ISBN inválido ou malformado")
	}

	comic := &models.Comic{
		ID:     uuid.New().String(),
		Title:  "Título não especificado",
		Author: "Autor não especificado",
		UserID: userID,
	}

	// Tentar buscar dados na API
	url := fmt.Sprintf("%s?q=isbn:%s&key=%s", googleBooksAPI, body.ISBN, s.apiKey)
	resp, err := s.c.R().Get(url)
	if err == nil {
		// Processar resposta da API
		var result map[string]interface{}
		if json.Unmarshal(resp.Body(), &result) == nil {
			items, ok := result["items"].([]interface{})
			if ok && len(items) > 0 {
				volumeInfo := items[0].(map[string]interface{})["volumeInfo"].(map[string]interface{})
				if volumeInfo != nil {
					comic.Title = volumeInfo["title"].(string)
					comic.Author = volumeInfo["authors"].([]interface{})[0].(string)
					imageLinks := volumeInfo["imageLinks"].(map[string]string)
					comic.CoverURL = imageLinks["thumbnail"]
				}
			}
		}
	} else {
		log.Warn().Msg("Erro ao consultar API, continuando com dados padrões")
	}

	// Salvar comic no banco
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
