package root

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": nil, "message": "OK"})
}
