package api

import (
	db "BakaFlash/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		db.Run()
		c.JSON(http.StatusOK, gin.H{
			"message": "test passed",
		})
	})
}
