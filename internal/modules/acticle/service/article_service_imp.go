package service

import (
	"errors"
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
	articles := a.articleRepo.List(limit)
	articlesResponses := articleDTO.ToArticles(articles)
	return articlesResponses
}

func (a *ArticleServiceImp) GetFeatured(limit int) articleDTO.Articles {
	articles := a.articleRepo.List(limit)
	articlesResponses := articleDTO.ToArticles(articles)
	return articlesResponses
}

func (a *ArticleServiceImp) Find(id int) (*articleDTO.Article, error) {
	article, err := a.articleRepo.Find(id)
	if err != nil {
		return nil, err
	}
	if article.ID == 0 {
		return nil, errors.New("article not found")
	}
	articleResponse := articleDTO.ToArticle(*article)
	return &articleResponse, nil
}
