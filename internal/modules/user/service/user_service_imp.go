package service

import (
	"errors"
	userDTO "github.com/phn00dev/go-blog-temp/internal/modules/user/dto"
	"github.com/phn00dev/go-blog-temp/internal/modules/user/dto/authDTO"
	userModel "github.com/phn00dev/go-blog-temp/internal/modules/user/models"
	userRepository "github.com/phn00dev/go-blog-temp/internal/modules/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImp struct {
	userRepo userRepository.UserRepository
}

func NewUserService() *UserServiceImp {
	return &UserServiceImp{
		userRepo: userRepository.NewUserRepository(),
	}
}

func (userService *UserServiceImp) Create(registerRequest authDTO.RegisterRequest) (userDTO.User, error) {
	var response userDTO.User
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return response, errors.New("error hashing password")
	}
	user := userModel.User{
		Name:         registerRequest.Username,
		Email:        registerRequest.Email,
		PasswordHash: string(hashPassword),
	}
	newUser := userService.userRepo.Create(user)
	return userDTO.ToUser(newUser), nil
}
