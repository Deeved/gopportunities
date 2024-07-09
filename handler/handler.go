package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "SHOW opening",
	})
}

func CreateOpeningHandle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "CREATE opening",
	})
}

func UpdateOpeningHandle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "UPDATE opening",
	})
}

func DeleteOpeningHandle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE opening",
	})
}

func ListOpeningsHandle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "LIST openings",
	})
}
