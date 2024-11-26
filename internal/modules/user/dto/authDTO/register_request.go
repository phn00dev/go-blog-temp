package authDTO

type RegisterRequest struct {
	Username string `form:"username" binding:"required,min=3,max=50"`
	Email    string `form:"email" binding:"required,email,min=3,max=100"`
	Password string `form:"password" binding:"required"`
}
