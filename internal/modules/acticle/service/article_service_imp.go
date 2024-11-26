package service

import (
	articleDTO "github.com/phn00dev/go-blog-temp/internal/modules/acticle/dto"
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

func (a *ArticleServiceImp) GetStories(limit int) articleDTO.Articles {
	articles := a.articleRepo.List(6)
	articlesResponses := articleDTO.ToArticles(articles)
	return articlesResponses
}

func (a *ArticleServiceImp) GetFeatured(limit int) articleDTO.Articles {
	articles := a.articleRepo.List(4)
	articlesResponses := articleDTO.ToArticles(articles)
	return articlesResponses
}
