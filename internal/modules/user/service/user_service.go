package service

import (
	userDTO "github.com/phn00dev/go-blog-temp/internal/modules/user/dto"
	"github.com/phn00dev/go-blog-temp/internal/modules/user/dto/authDTO"
)

type UserService interface {
	Create(registerRequest authDTO.RegisterRequest) (userDTO.User, error)
	CheckUserExists(email string) bool
	HandleUserLogin(loginRequest authDTO.LoginRequest) (userDTO.User, error)
}
