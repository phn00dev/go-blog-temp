package repository

import "github.com/phn00dev/go-blog-temp/internal/modules/acticle/models"

type ArticleRepository interface {
	List(limit int) []models.Article
}
