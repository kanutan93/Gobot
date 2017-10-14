package main

import (
	"github.com/vova616/screenshot"
	"github.com/lazywei/go-opencv/opencv"
	"image"
	"os"
	"image/png"
	"time"
	"fmt"
)

func CreateScreenshot(filename string) {
	img, err := screenshot.CaptureScreen()
	check(err)
	myImg := image.Image(img)
	file, err := os.Create(filename)
	check(err)
	png.Encode(file, myImg)
	fmt.Println("Screenshot created at" + time.Now().String())
}

func GetPointsBox(filename string, croppedFilename string) {
	image := opencv.LoadImage(filename)
	defer image.Release()
	crop := opencv.Crop(image, 0, 0, 50, 50)
	opencv.SaveImage(croppedFilename, crop, nil)
	crop.Release()
}

