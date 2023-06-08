package auth

type UserInput struct {
	FullName string `json:"fullname" binding:"required"`
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Type     string `json:"type" binding:"required"`
}

type SignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AddressInput struct {
	Address string `json:"address" binding:"required"`
	City    string `json:"city" binding:"required"`
}
