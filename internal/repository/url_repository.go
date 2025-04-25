package repository

import (
	"context"
	"errors"

	"github.com/nnnn24/url_shortener_service/internal/models"
	"gorm.io/gorm"
)

type URLRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) *URLRepository {
	return &URLRepository{
		db: db,
	}
}

func (r *URLRepository) Create(ctx context.Context, url *models.Url) error {
	if err := r.db.WithContext(ctx).Create(url).Error; err != nil {
		return err
	}

	return nil
}

func (r *URLRepository) FindByShortCode(ctx context.Context, shortCode string) (*models.Url, error) {
	var url *models.Url

	err := r.db.WithContext(ctx).Where(&models.Url{ShortCode: shortCode}).First(&url).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	return url, err
}

func (r *URLRepository) IsURLExist(ctx context.Context, url string) bool {
	var model *models.Url
	if err := r.db.WithContext(ctx).
		Where("url=?", url).
		First(&model).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

func (r *URLRepository) IncrementClicks(ctx context.Context, shortCode string) error {
	if err := r.db.WithContext(ctx).
		Model(&models.Url{}).
		Where(&models.Url{ShortCode: shortCode}).
		Update("clicks", gorm.Expr("clicks + ?", 1)).Error; err != nil {
		return err
	}

	return nil
}

func (r *URLRepository) UpdateURL(ctx context.Context, newUrl string, shortCode string) (*models.Url, error) {
	url := &models.Url{
		ShortCode: shortCode,
	}

	if err := r.db.WithContext(ctx).Model(&url).Where(&models.Url{ShortCode: shortCode}).Update("url", newUrl).Error; err != nil {
		return nil, err
	}

	return url, nil

}

func (r *URLRepository) DeleteURL(ctx context.Context, url *models.Url) error {
	if err := r.db.WithContext(ctx).Delete(&url).Error; err != nil {
		return err
	}
	return nil
}
