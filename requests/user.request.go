package requests

import "go-auth-service/models"



type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string          `json:"username" binding:"required,min=3,max=100"`
	Password string          `json:"password" binding:"required,min=6"`
	Role     models.UserRole `json:"role" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
