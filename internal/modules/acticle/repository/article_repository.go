package repository

import articleModel "github.com/phn00dev/go-blog-temp/internal/modules/acticle/models"

type ArticleRepository interface {
	List(limit int) []articleModel.Article
	Find(id int) (*articleModel.Article, error)
}
