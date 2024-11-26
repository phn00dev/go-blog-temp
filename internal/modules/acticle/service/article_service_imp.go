package service

import (
	articleModel "github.com/phn00dev/go-blog-temp/internal/modules/acticle/models"
	articleRepository "github.com/phn00dev/go-blog-temp/internal/modules/acticle/repository"
)

type ArticleServiceImp struct {
	articleRepo articleRepository.ArticleRepository
}

func NewArticleServiceImp() *ArticleServiceImp {
	return &ArticleServiceImp{
		articleRepo: articleRepository.NewArticleRepo(),
	}
}

func (a *ArticleServiceImp) GetStories(limit int) []articleModel.Article {
	articles := a.articleRepo.List(6)
	return articles
}

func (a *ArticleServiceImp) GetFeatured(limit int) []articleModel.Article {
	articles := a.articleRepo.List(4)
	return articles
}
