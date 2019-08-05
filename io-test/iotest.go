package main

import (
	"github.com/gin-gonic/gin"
	"github.com/warthog618/gpio"
)

var IN1 := gpio.NewPin(12)
var IN2 := gpio.NewPin(13)
var IN3 := gpio.NewPin(20)
var IN4 := gpio.NewPin(21)
var ENA := gpio.NewPin(6)
var ENB := gpio.NewPin(26)

func setupPin() {
	IN1.Output()
	IN2.Output()
	IN3.Output()
	IN4.Output()
	ENA.Output()
	ENB.Output()
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

	// err := gpio.Open()
	// fmt.Println(err)
	apiHandler()

}
