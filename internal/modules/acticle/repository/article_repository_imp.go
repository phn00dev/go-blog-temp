package repository

import (
	articleModel "github.com/phn00dev/go-blog-temp/internal/modules/acticle/models"
	"github.com/phn00dev/go-blog-temp/pkg/database"
	"gorm.io/gorm"
	"log"
)

type ArticleRepositoryImp struct {
	DB *gorm.DB
}

func NewArticleRepo() *ArticleRepositoryImp {
	return &ArticleRepositoryImp{
		DB: database.Connection(),
	}
}

func (a *ArticleRepositoryImp) List(limit int) []articleModel.Article {
	var articles []articleModel.Article
	err := a.DB.Limit(limit).Joins("User").Find(&articles).Error
	if err != nil {
		log.Fatal("article repository error :", err)
		return articles
	}
	return articles
}

func (a *ArticleRepositoryImp) Find(id int) (*articleModel.Article, error) {
	var article articleModel.Article
	err := a.DB.Joins("User").First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (a *ArticleRepositoryImp) Create(article articleModel.Article) articleModel.Article {
	var newArticle articleModel.Article
	a.DB.Create(&article).Scan(&newArticle)
	return newArticle
}
