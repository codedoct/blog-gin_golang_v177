package config

import (
	"github.com/gin-gonic/gin"

	"blog-gin_golang_v177/app/controllers/root"
	"blog-gin_golang_v177/config/collection"
	"blog-gin_golang_v177/db"
)

var Routers = gin.Default()

func init() {
	corsConfig(Routers)

	Routers.GET("/", root.Index)
	main := Routers.Group("v1")
	collection.MainRouter(db.DB, main)
}
