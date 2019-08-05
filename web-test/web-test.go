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
	r.Run(":8080")
}

func main() {

	apiHandler()

}
