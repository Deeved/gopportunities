package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "CREATE opening",
	})
}
