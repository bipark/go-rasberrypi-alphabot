package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stianeikeland/go-rpio"
)

func setupPin() {

	err := rpio.Open()
	if err != nil {
		fmt.Printf("i/o Open error")
		return
	}

}

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

	apiHandler()

}
