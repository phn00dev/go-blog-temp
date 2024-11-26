package repository

import userModel "github.com/phn00dev/go-blog-temp/internal/modules/user/models"

type UserRepository interface {
	Create(user userModel.User) userModel.User
	FindByEmail(email string) userModel.User
}
