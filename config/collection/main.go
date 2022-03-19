package collection

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"blog-gin_golang_v177/app/controllers/user"
)

func MainRouter(db *gorm.DB, main *gin.RouterGroup) {
	userCtrl := user.RoleController(db)
	user := main.Group("role")
	{
		user.GET("/", userCtrl.GetRole)
	}
}
