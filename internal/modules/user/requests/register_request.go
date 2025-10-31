package requests

type RegisterRequest struct {
	Name     string `form:"name" binding:"required,min=3,max=20"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=8,max=20"`
}
