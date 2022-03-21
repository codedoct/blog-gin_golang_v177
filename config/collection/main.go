package collection

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"blog-gin_golang_v177/app/controllers/user"
)

func MainRouter(db *gorm.DB, main *gin.RouterGroup) {
	userRoleCtrl := user.RoleController(db)
	userRoute := main.Group("role")
	{
		userRoute.GET("/", userRoleCtrl.GetRole)
	}

	userAuthCtrl := user.AuthController(db)
	authRoute := main.Group("auth")
	{
		authRoute.POST("/register", userAuthCtrl.Register)
		authRoute.POST("/login", userAuthCtrl.Login)
	}
}
