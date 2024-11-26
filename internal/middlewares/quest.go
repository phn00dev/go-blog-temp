package middlewares

import (
	"github.com/gin-gonic/gin"
	userRepository "github.com/phn00dev/go-blog-temp/internal/modules/user/repository"
	"github.com/phn00dev/go-blog-temp/pkg/sessions"
	"net/http"
	"strconv"
)

func IsGuest() gin.HandlerFunc {
	var userRepo = userRepository.NewUserRepository()
	return func(ctx *gin.Context) {

		authID := sessions.Get(ctx, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindById(userID)
		if user.ID != 0 {
			ctx.Redirect(http.StatusFound, "/")
			return
		}
		ctx.Next()
	}
}
