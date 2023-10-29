package api

import (
	"BakaFlash/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		database.run()
		c.JSON(http.StatusOK, gin.H{
			"message": "test passed",
		})
	})
}
