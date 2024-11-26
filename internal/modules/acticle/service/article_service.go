package service

import articleModel "github.com/phn00dev/go-blog-temp/internal/modules/acticle/models"

type ArticleService interface {
	GetStories(limit int) []articleModel.Article
	GetFeatured(limit int) []articleModel.Article
}
