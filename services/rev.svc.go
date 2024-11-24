package services

import (
    "acervoback/models"
    "acervoback/repository"
    "github.com/google/uuid"
)

type ReviewService struct {
    reviewRepo repository.ReviewRepo
}

func NewReviewService(reviewRepo repository.ReviewRepo) *ReviewService {
    return &ReviewService{reviewRepo: reviewRepo}
}

func (s *ReviewService) AddReview(userID, comicID string, rating int, comment string) (*models.Review, error) {
    review := &models.Review{
        ID:      uuid.New().String(),
        ComicID: comicID,
        UserID:  userID,
        Rating:  rating,
        Comment: comment,
    }

    err := s.reviewRepo.Create(review)
    if err != nil {
        return nil, err
    }

    return review, nil
}

func (s *ReviewService) GetReviewsByComic(comicID string) ([]models.Review, error) {
    return s.reviewRepo.FindByComicID(comicID)
}

