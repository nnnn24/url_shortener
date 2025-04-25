package models

import (
	"time"

	"gorm.io/gorm"
)

type Url struct {
	gorm.Model

	URL       string     `gorm:"not null;index" json:"url" binding:"required,url"`
	ShortCode string     `gorm:"uniqueIndex;not null" json:"short_code"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	Clicks    int64      `gorm:"default:0" json:"clicks"`
}

type CreateURLRequest struct {
	Original_URL string `json:"original_url" binding:"required,url"`
}

type CreateURLResponse struct {
	ShortURL string `json:"short_url"`
}

type UpdateURLRequest struct {
	URL string `json:"url" binding:"required,url"`
}
