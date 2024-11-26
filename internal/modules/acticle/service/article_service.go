package service

import (
	articleDTO "github.com/phn00dev/go-blog-temp/internal/modules/acticle/dto"
	userDTO "github.com/phn00dev/go-blog-temp/internal/modules/user/dto"
)

type ArticleService interface {
	GetStories(limit int) articleDTO.Articles
	GetFeatured(limit int) articleDTO.Articles
	Find(id int) (*articleDTO.Article, error)
	Store(createRequest articleDTO.CreateArticleRequest, user userDTO.User) (articleDTO.Article, error)
}
