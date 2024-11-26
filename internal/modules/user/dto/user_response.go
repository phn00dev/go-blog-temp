package dto

import (
	"fmt"
	userModel "github.com/phn00dev/go-blog-temp/internal/modules/user/models"
)

type User struct {
	ID    uint
	Image string
	Name  string
	Email string
}

type Users struct {
	Data []User
}

func ToUser(user userModel.User) User {
	return User{
		ID:    user.ID,
		Image: fmt.Sprintf("https://www.ui-avatars.com/api/?name=%s", user.Name),
		Name:  user.Name,
		Email: user.Email,
	}
}
