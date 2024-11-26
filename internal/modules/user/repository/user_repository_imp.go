package repository

import (
	userModel "github.com/phn00dev/go-blog-temp/internal/modules/user/models"
	"github.com/phn00dev/go-blog-temp/pkg/database"
	"gorm.io/gorm"
)

type UserRepositoryImp struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepositoryImp {
	return &UserRepositoryImp{
		DB: database.Connection(),
	}
}

func (userRepo *UserRepositoryImp) Create(user userModel.User) userModel.User {
	var newUser userModel.User
	userRepo.DB.Create(&user).Scan(&newUser)
	return newUser
}

func (userRepo *UserRepositoryImp) FindByEmail(email string) userModel.User {
	var user userModel.User
	userRepo.DB.First(&user, "email = ?", email)
	return user
}

func (userRepo *UserRepositoryImp) FindById(id int) userModel.User {
	var user userModel.User
	userRepo.DB.First(&user, "id = ?", id)
	return user
}
