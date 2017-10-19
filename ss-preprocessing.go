package main

import (
	"github.com/vova616/screenshot"
	"github.com/lazywei/go-opencv/opencv"
	"image"
	"os"
	"image/png"
	"time"
	"fmt"
	"github.com/otiai10/gosseract"
	"strings"
	"io/ioutil"
	"encoding/json"
)

type BoxesCoords struct {
	HpBox BoxCoords
	MpBox BoxCoords
}

type BoxCoords struct {
	X int
	Y int
	Width int
	Height int
}

func readCoordsConf(filename string) BoxesCoords{
	coordsConfig, err := ioutil.ReadFile(filename)
	check(err)
	var coords BoxesCoords
	json.Unmarshal([]byte(coordsConfig), &coords)
	return coords
}

func createScreenshot(filename string) {
	img, err := screenshot.CaptureScreen()
	check(err)
	myImg := image.Image(img)
	file, err := os.Create(filename)
	check(err)
	png.Encode(file, myImg)
	fmt.Println("Screenshot created at: " + time.Now().String())
}

func getPoints(filename string, coords *BoxesCoords) map[string]string {
	image := opencv.LoadImage(filename)
	defer image.Release()
	hpBox := getBoxFilename("hpBox.png", image, &coords.HpBox)
	mpBox := getBoxFilename("mpBox.png", image, &coords.MpBox)
	boxes:= map[string]string{"hp": hpBox, "mp": mpBox}
	return boxes
}

func getBoxFilename(croppedFilename string, image *opencv.IplImage, boxCoords *BoxCoords) string {
	box := opencv.Crop(image, boxCoords.X, boxCoords.Y, boxCoords.Width, boxCoords.Height)
	opencv.SaveImage(croppedFilename, box, nil)
	box.Release()
	return croppedFilename
}

func recognizeText(boxes map[string]string) map[string]string {
	for key, value := range boxes {
		boxes[key] = strings.TrimSpace(gosseract.Must(gosseract.Params{Src: value, Languages: "eng"}))
	}
	return boxes
}
