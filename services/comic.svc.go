package services

import (
	"acervoback/models"
	"acervoback/models/requests"
	"acervoback/repository"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

type ComicService struct {
	repo repository.ComicRepo
	c    *resty.Client
}

func NewComicService(repo repository.ComicRepo) *ComicService {
	return &ComicService{repo: repo, c: resty.New()}
}
func (s *ComicService) Create(userID string, body requests.NewComicRequest) (*models.Comic, error) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/isbn/v1/%s", body.ISBN)
	resp, err := s.c.R().Get(url)
	if err != nil {
		log.Err(err).Msg("Erro ao buscar dados do ISBN")
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("ISBN invalido")
	}
	comic := s.jsonToComic(resp.String())
	comic.UserID = userID
	err = s.repo.Create(comic)
	if err != nil {
		return nil, err
	}
	return comic, nil
}

func (s *ComicService) jsonToComic(json string) *models.Comic {
	title := gjson.Get(json, "title").String()
	publisher := gjson.Get(json, "publisher").String()
	cover := gjson.Get(json, "cover_url").String()
	year := gjson.Get(json, "year").Int()
	authors := ""
	for _, author := range gjson.Get(json, "authors").Array() {
		authors += author.String() + ", "
	}

	return &models.Comic{
		ID:        uuid.New().String(),
		Title:     title,
		Author:    authors,
		Publisher: publisher,
		CoverURL:  cover,
		Year:      int(year),
		Owner:     models.User{},
		UserID:    "",
	}
}

func (s *ComicService) GetComics(userID string) ([]models.Comic, error) {
	comics, err := s.repo.FindByOwner(userID)
	if err != nil {
		return nil, err
	}
	return comics, nil
}
