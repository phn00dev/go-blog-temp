package dto

import (
	articleModel "github.com/phn00dev/go-blog-temp/internal/modules/acticle/models"
	userDTO "github.com/phn00dev/go-blog-temp/internal/modules/user/dto"
)

type Article struct {
	ID        uint
	Image     string
	Title     string
	Content   string
	CreatedAt string
	User      userDTO.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article articleModel.Article) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Image:     "/assets/image.png",
		CreatedAt: article.CreatedAt.Format("02-01-2006 15:04:05"),
		User:      userDTO.ToUser(article.User),
	}
}

func ToArticles(articles []articleModel.Article) Articles {
	var responses Articles
	for _, article := range articles {
		responses.Data = append(responses.Data, ToArticle(article))
	}
	return responses
}
