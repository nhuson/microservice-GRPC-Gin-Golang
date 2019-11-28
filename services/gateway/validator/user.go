package validator

type CreateUser struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}
