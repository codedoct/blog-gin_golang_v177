package config

import (
	"blog-gin_golang_v177/app/controllers/root"
	"github.com/gin-gonic/gin"
)

var Routers = gin.Default()

func init() {
	corsConfig(Routers)

	Routers.GET("/", root.Index)
}
