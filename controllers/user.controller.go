package controllers

import (
	"go-auth-service/databases"
	"go-auth-service/requests"
	"go-auth-service/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(resource *databases.Resource) *UserController {
	userService := services.NewUserService(resource)
	return &UserController{userService: userService}
}

// @Summary Login
// @Tags Users
// @version 1.0
// @Description Login with the input payload
// @Accept  application/json
// @Produce  application/json
// @Param user body requests.LoginRequest true "User for login"
// @Success 200 {object} string "OK"
// @Router /login [post]
func (controller *UserController) Login(c *gin.Context) {
	var req requests.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// responses.Error(c, "Invalid input")
		return
	}

	controller.userService.Login(c, req)
}
