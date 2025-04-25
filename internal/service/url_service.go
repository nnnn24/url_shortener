package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/nnnn24/url_shortener_service/internal/models"
	"github.com/nnnn24/url_shortener_service/internal/repository"
)

type URLService struct {
	urlRepo *repository.URLRepository
}

func NewURLService(urlRepo *repository.URLRepository) *URLService {
	return &URLService{
		urlRepo: urlRepo,
	}
}

func (s *URLService) CreateShortURL(ctx context.Context, req *models.CreateURLRequest) (*models.CreateURLResponse, error) {
	// Generate a unique short code
	shortCode, err := generateShortCode()
	if err != nil {
		return nil, err
	}

	isAlreadyExist := s.urlRepo.IsURLExist(ctx, req.Original_URL)

	if isAlreadyExist {
		return nil, errors.New("already exists")
	}

	url := &models.Url{
		URL:       req.Original_URL,
		ShortCode: shortCode,
	}

	if err := s.urlRepo.Create(ctx, url); err != nil {
		return nil, err
	}

	return &models.CreateURLResponse{
		ShortURL: shortCode,
	}, nil
}

func (s *URLService) FindURL(ctx context.Context, shortCode string) (*models.Url, error) {
	url, err := s.urlRepo.FindByShortCode(ctx, shortCode)

	if err != nil {
		return nil, err
	}

	if url.ExpiresAt != nil && time.Now().After(*url.ExpiresAt) {
		return nil, errors.New("expired")
	}

	if err := s.urlRepo.IncrementClicks(ctx, shortCode); err != nil {
		return nil, err
	}

	return url, nil
}

func (s *URLService) UpdateURL(ctx context.Context, req *models.UpdateURLRequest, shortCode string) (*models.Url, error) {
	if _, err := s.urlRepo.FindByShortCode(ctx, shortCode); err != nil {
		return nil, err
	}

	url, err := s.urlRepo.UpdateURL(ctx, req.URL, shortCode)

	if err != nil {
		return nil, err
	}

	return url, nil

}

func (s *URLService) DeleteURL(ctx context.Context, shortCode string) error {
	url, err := s.urlRepo.FindByShortCode(ctx, shortCode)

	if err != nil {
		return err
	}

	if err := s.urlRepo.DeleteURL(ctx, url); err != nil {
		return err
	}

	return nil
}

func generateShortCode() (string, error) {
	b := make([]byte, 6)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:6], nil
}
