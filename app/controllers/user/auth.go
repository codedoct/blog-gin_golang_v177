package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"

	"blog-gin_golang_v177/domain/user"
	"blog-gin_golang_v177/domain/user/model"
	"blog-gin_golang_v177/domain/user/repository"
	"blog-gin_golang_v177/lib/response"
)

type authController struct {
	AuthService user.AuthServiceInterface
}

func AuthController(db *gorm.DB) *authController {
	return &authController{
		AuthService: user.AuthService(repository.AuthRepository(db)),
	}
}

func (ac *authController) Register(context *gin.Context) {
	var reqBody model.ReqBodyRegister
	if err := context.ShouldBind(&reqBody); err != nil {
		response.Error(context, http.StatusBadRequest, err.Error())
		return
	}
	if errStatus, err := ac.AuthService.Register(&reqBody); err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}
	response.Json(context, http.StatusOK, nil)
}
