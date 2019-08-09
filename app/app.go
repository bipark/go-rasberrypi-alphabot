package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hybridgroup/mjpeg"
	"gocv.io/x/gocv"
)

var (
	err    error
	webcam *gocv.VideoCapture
	stream *mjpeg.Stream
)

func mjpegCapture() {
	img := gocv.NewMat()
	defer img.Close()

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", 0)
			return
		}
		if img.Empty() {
			continue
		}

		buf, _ := gocv.IMEncode(".jpg", img)
		stream.UpdateJPEG(buf)
	}
}

// 카메라 스트림
func cameraHandler(c *gin.Context) {
	c.Stream(func(w io.Writer) bool {
		c.SSEvent("message", stream)
		return true
	})
}

func apiHandler() {
	r := gin.Default()
	r.Static("/static", "./public")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "static")
	})

	r.GET("/stream", func(c *gin.Context) {
		cameraHandler(c)
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

	webcam, err = gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Printf("Error opening capture device: %v\n", 0)
		return
	}
	defer webcam.Close()
	stream = mjpeg.NewStream()
	go mjpegCapture()

	apiHandler()

}
