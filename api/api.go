package api

import (
	db "BakaFlash/database"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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
			"message": "User added",
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

	// /uploadImage recives a post request with a ParseMultipartForm
	r.POST("/uploadImage", func(c *gin.Context) {
		err := c.Request.ParseMultipartForm(10 << 20) // 10 MB max for the entire request
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to parse multipart form",
			})
			return
		}

		file, _, err := c.Request.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to retrieve the 'image' file from the form",
			})
			return
		}

		// Process the file
		imageType := c.Request.FormValue("imageType")
		imageName := c.Request.FormValue("imageName")

		// You can save the file to disk or handle it as needed
		// For example, save it to a local file:
		// err = c.SaveUploadedFile(file, "uploads/"+imageName)

		// Or, save it to the database using your UploadImage function
		imageBytes, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to read the file",
			})
			return
		}

		// Call your UploadImage function to save the image to the database
		db.UploadImage(imageName, imageType, imageBytes)

		c.JSON(http.StatusOK, gin.H{
			"message": "Image uploaded",
		})
	})

	// /getImage by name posts the image with the given name to the client
	// /getImage by name retrieves the image with the given name and serves it to the client
	r.POST("/getImage", func(c *gin.Context) {
		imageName := c.Request.FormValue("imageName")
		image := db.GetImage(imageName)
		if image == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Image not found",
			})
			return
		}
		c.Data(http.StatusOK, "image/jpeg", image)
	})

	r.GET("/listImageNames", func(c *gin.Context) {
		imageNames := db.ListImageNames()
		c.JSON(http.StatusOK, imageNames)
	})
}
