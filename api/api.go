package api

import (
	db "BakaFlash/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(r *gin.Engine) {
	//test/string/string
	r.GET("/test/:fistName/:lastName", func(c *gin.Context) {
		db.Run(c.Param("fistName"), c.Param("lastName"))
		c.JSON(http.StatusOK, gin.H{
			"message": "test passed",
		})
	})
	// /addUser reseve a post request with a json body
	// fistName : string
	// lastName : string
	r.POST("/addUser", func(c *gin.Context) {
		var json struct {
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
		}
		c.BindJSON(&json)
		db.Run(json.FirstName, json.LastName)
		c.JSON(http.StatusOK, gin.H{
			"message": "test passed",
		})
	})
	// /getUsers returns a json array of all users in the database
	r.GET("/getUsers", func(c *gin.Context) {
		//GetUsers returns a slice of person structs
		users := db.GetUsers()
		userJson := []gin.H{}
		for _, user := range users {
			userJson = append(userJson, gin.H{
				"id":        user.Id,
				"firstName": user.FirstName,
				"lastName":  user.LastName,
			})
		}
		c.JSON(http.StatusOK, userJson)

	})
}
