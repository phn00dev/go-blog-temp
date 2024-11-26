package helpers

import (
	"github.com/gin-gonic/gin"
	userDTO "github.com/phn00dev/go-blog-temp/internal/modules/user/dto"
	userRepository "github.com/phn00dev/go-blog-temp/internal/modules/user/repository"
	"github.com/phn00dev/go-blog-temp/pkg/sessions"
	"strconv"
)

func Auth(ctx *gin.Context) userDTO.User {
	var response userDTO.User
	authID := sessions.Get(ctx, "auth")
	userID, _ := strconv.Atoi(authID)
	var userRepo = userRepository.NewUserRepository()
	user := userRepo.FindById(userID)
	if user.ID == 0 {
		return response
	}
	return userDTO.ToUser(user)
}
