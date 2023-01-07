package postgresql

import (
	"context"

	"github.com/durmusrasit/pano-api/internal/backend"
	"github.com/durmusrasit/pano-api/internal/models"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type PostgreSQLBackend struct {
	DB *gorm.DB
}

func NewPostgreSQLBackend(db *gorm.DB) backend.Backender {
	return &PostgreSQLBackend{
		DB: db,
	}
}

func (b *PostgreSQLBackend) GetPost(ctx context.Context, id string) (*models.Post, error) {
	post := models.Post{}

	result := b.DB.First(&post, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil
}

func (b *PostgreSQLBackend) CreatePost(ctx context.Context, title string, url string, content string, userId string) (*models.Post, error) {
	slug := slug.Make(title)

	post := models.Post{
		Title:   title,
		Url:     url,
		Content: content,
		Slug:    slug,
		UserID:  userId,
	}

	result := b.DB.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil
}