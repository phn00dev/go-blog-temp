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

func (userService *UserServiceImp) CheckUserExists(email string) bool {
	user := userService.userRepo.FindByEmail(email)
	if user.ID != 0 {
		return true
	}
	return false
}

func (userService *UserServiceImp) HandleUserLogin(loginRequest authDTO.LoginRequest) (userDTO.User, error) {
	var response userDTO.User
	existingUser := userService.userRepo.FindByEmail(loginRequest.Email)
	if existingUser.ID == 0 {
		return response, errors.New("invalid credentials")
	}
	err := bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(loginRequest.Password))
	if err != nil {
		return response, errors.New("invalid credentials")
	}
	return userDTO.ToUser(existingUser), nil
}
