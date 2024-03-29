package mappers

type LoginForm struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type RegisterForm struct {
	UserName string  `form:"username" json:"username" binding:"required,max=30,min=5"`
	Password string  `form:"password" json:"password" binding:"required"`
	Email    *string `form:"email" json:"email" binding:"omitempty,email"`
}

type Token struct {
	RefreshToken string `form:"refresh_token" json:"refresh_token" binding:"required"`
}
