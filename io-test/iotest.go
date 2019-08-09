package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func apiHandler() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/foward", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "go foward",
		})
	})

	router.Run()
}
func main() {

	// err := gpio.Open()
	// fmt.Println(err)
	apiHandler()

}
