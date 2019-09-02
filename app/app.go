package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hybridgroup/mjpeg"
	"gocv.io/x/gocv"
)

var (
	webcam *gocv.VideoCapture
	stream *mjpeg.Stream
	err    error
	// in1 *gpio.Pin
	// in2 *gpio.Pin
	// in3 *gpio.Pin
	// in4 *gpio.Pin
	// ena *gpio.Pin
	// enb *gpio.Pin
)

// func setupPin() {

// 	err := gpio.Open()
// 	if err != nil {
// 		fmt.Println(err)
// 		panic(err)
// 	}
// 	defer gpio.Close()

// 	in1 = gpio.NewPin(gpio.GPIO12) // 12
// 	in2 = gpio.NewPin(gpio.GPIO13) // 13
// 	in3 = gpio.NewPin(gpio.GPIO20) // 20
// 	in4 = gpio.NewPin(gpio.GPIO21) // 21
// 	ena = gpio.NewPin(gpio.GPIO7)  // 6
// 	enb = gpio.NewPin(gpio.J8p37)  // 26

// 	in1.Output()
// 	in2.Output()
// 	in3.Output()
// 	in4.Output()
// 	ena.Output()
// 	enb.Output()

// 	in1.Low()
// 	in2.Low()
// 	in3.Low()
// 	in4.Low()
// 	ena.Low()
// 	enb.Low()

// }

// func gostop() {
// 	in1.Low()
// 	in2.Low()
// 	in3.Low()
// 	in4.Low()
// }

// func goforward() {
// 	fmt.Printf("%v\n\n\n", in1)
// 	in1.High()
// 	in2.Low()
// 	in3.Low()
// 	in4.High()
// 	fmt.Printf("%v\n\n\n", in1)
// }

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

func main() {

	// open webcam
	webcam, err = gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Printf("Error opening capture device: %v\n", 0)
		return
	}
	defer webcam.Close()

	//setupPin()

	// Static Files
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// Stream API
	stream = mjpeg.NewStream()
	go mjpegCapture()

	http.Handle("/stream", stream)

	// http run
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))

}
