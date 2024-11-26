package seeder

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	articleModel "github.com/phn00dev/go-blog-temp/internal/modules/acticle/models"
	userModel "github.com/phn00dev/go-blog-temp/internal/modules/user/models"
	"github.com/phn00dev/go-blog-temp/pkg/database"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

func Seed() {
	db := database.Connection()
	var users []userModel.User
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("12345678"), 12)
	for i := 1; i <= 10; i++ {
		user := userModel.User{
			Name:         "user" + strconv.Itoa(i),
			Email:        "user" + strconv.Itoa(i) + "@gmail.com",
			PasswordHash: string(passwordHash),
		}
		users = append(users, user)
	}
	if err := db.Create(&users).Error; err != nil {
		log.Fatal("user creation error", err)
		return
	}

	fmt.Println("users seeder successfully")

	var articles []articleModel.Article

	for i := 1; i <= 10; i++ {

		article := articleModel.Article{
			Title:   gofakeit.Sentence(1),
			Content: gofakeit.Sentence(20),
			UserID:  uint(i),
		}
		articles = append(articles, article)
	}
	if err := db.Create(&articles).Error; err != nil {
		log.Fatal("article creation error", err)
		return
	}
	fmt.Println("article seeder successfully")
}
