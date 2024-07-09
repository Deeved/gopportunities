package router

import (
	"github.com/Deeved/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandle)
		v1.POST("/opening", handler.CreateOpeningHandle)
		v1.PUT("/opening", handler.UpdateOpeningHandle)
		v1.DELETE("/opening", handler.DeleteOpeningHandle)
		v1.GET("/openings", handler.ListOpeningsHandle)
	}
}
