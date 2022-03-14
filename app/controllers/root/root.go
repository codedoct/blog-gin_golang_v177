package root

import (
	"blog-gin_golang_v177/lib/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	response.Json(context, http.StatusOK, nil)
}
