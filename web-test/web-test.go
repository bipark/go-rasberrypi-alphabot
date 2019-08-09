package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func apiHandler() {
	r := gin.Default()
	r.Static("/static", "./public")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "static")
	})

	r.GET("/api/go/foward", func(c *gin.Context) {
		c.String(200, "Foward")
	})
	r.GET("/api/go/backward", func(c *gin.Context) {
		c.String(200, "Backward")
	})
	r.GET("/api/go/left", func(c *gin.Context) {
		c.String(200, "Left")
	})
	r.GET("/api/go/right", func(c *gin.Context) {
		c.String(200, "Right")
	})
	r.Run(":8080")
}

func main() {

	apiHandler()

}
