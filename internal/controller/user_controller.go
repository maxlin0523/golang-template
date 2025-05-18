package controller

import (
	"go-template/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// @Summary      GetUser
// @Tags         users
// @Produce      json
// @Success      200  {object} UserResponse
// @Router       /public/api/v1/user [get]
func (uc *UserController) GetUser(ctx *gin.Context) {
	user := uc.userService.GetUser()
	ctx.JSON(http.StatusOK, UserResponse{
		User: user,
	})
}

type UserResponse struct {
	User string `json:"user"`
}
