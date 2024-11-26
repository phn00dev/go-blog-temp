package migration

import (
	"fmt"
	articleModel "github.com/phn00dev/go-blog-temp/internal/modules/acticle/models"
	userModel "github.com/phn00dev/go-blog-temp/internal/modules/user/models"
	"github.com/phn00dev/go-blog-temp/pkg/database"
	"log"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&userModel.User{}, articleModel.Article{})
	if err != nil {
		log.Println("database migration failed error :", err)
		return
	}
	fmt.Println("Migration Done...")
}
