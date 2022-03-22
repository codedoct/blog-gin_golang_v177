package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"

	auth "blog-gin_golang_v177/domain/user"
	"blog-gin_golang_v177/domain/user/repository"
	authLib "blog-gin_golang_v177/lib/auth"
	"blog-gin_golang_v177/lib/constant"
	"blog-gin_golang_v177/lib/response"
)

func Auth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authService := auth.AuthService(repository.AuthRepository(db))
		user, err := authService.CheckAuth(c.Request.Header.Get("Authorization"))
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		if user.Role.Name != "Super Admin" {
			response.Error(c, http.StatusUnauthorized, constant.NotAuthorize)
			c.Abort()
			return
		}

		userStr, err := json.Marshal(&authLib.AuthData{
			ID:        user.ID,
			Email:     user.Email,
			RoleID:    user.RoleID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Set("auth", string(userStr))
	}
}
