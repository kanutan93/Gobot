package main

import (
	"github.com/lazywei/go-opencv/opencv"
	"github.com/otiai10/gosseract"
	"strings"
	"io/ioutil"
	"encoding/json"
)

const (
	textBoxFilename = "textBox.png"
	targetsImageFilename = "targets.png"
	colorConf = "color.json"
)

type BoxCoords struct {
	X      int
	Y      int
	Width  int
	Height int
}

type ColorConfig struct {
	Green Color
	White Color
}

type Color struct {
	Lower []float64
	Upper []float64
	ColorIndent float64
}

func readCoordsConf(filename string) BoxCoords {
	coordsJson, err := ioutil.ReadFile(filename)
	check(err)
	var coords BoxCoords
	json.Unmarshal([]byte(coordsJson), &coords)
	return coords
}

func readColorConf(filename string) ColorConfig {
	colorJson, err := ioutil.ReadFile(filename)
	check(err)
	var colorConfig ColorConfig
	json.Unmarshal([]byte(colorJson), &colorConfig)
	return colorConfig
}

func getText(filename string, coords *BoxCoords) map[string]string {
	img := opencv.LoadImage(filename)
	defer img.Release()
	textImage := getBoxFilename(textBoxFilename, img, coords)
	text := map[string]string{"text": textImage}
	return text
}

func getBoxFilename(pointsImageFilename string, image *opencv.IplImage, boxCoords *BoxCoords) string {
	box := opencv.Crop(image, boxCoords.X, boxCoords.Y, boxCoords.Width, boxCoords.Height)
	opencv.SaveImage(pointsImageFilename, box, nil)
	box.Release()
	return pointsImageFilename
}

func recognizePoints(points map[string]string) map[string]string {
	for key, value := range points {
		points[key] = strings.TrimSpace(gosseract.Must(gosseract.Params{Src: value, Languages: "eng"}))
	}
	return points
}

func drawTargetBox(img *opencv.IplImage, x int, y int) {
	opencv.Rectangle(img,
		opencv.Point{x - 30, y},
		opencv.Point{x, y + 30},
		opencv.NewScalar(170, 255, 102, 255),
		3, 4, 0)
}

func findTarget(filename string, color *Color) map[string]int{
	img := opencv.LoadImage(filename)
	defer img.Release()

	coords := map[string]int{"x": 0, "y": 0}
	colorUpper := [4]float64{color.Upper[0], color.Upper[1], color.Upper[2], 0} //верхняя граница зеленого цвета BGR
	colorLower := [4]float64{color.Lower[0], color.Lower[1], color.Lower[2], 0} //нижняя граница зеленого цвета BGR

	for x := 0; x < img.Width(); x++ {
		for y := 0; y < img.Height(); y++ {
			scalar := img.Get2D(x, y)
			scalarArray := scalar.Val() //Массив пикселей изображения
			if((scalarArray[0] < scalarArray[1] - 70) && (scalarArray[1] > scalarArray[2] - 70)){ // Если зеленого больше
				if (scalarArray[0] >= colorLower[0]) && (scalarArray[0] <= colorUpper[0]) { //, то ищем пиксель в заданном выше диапозоне
					if (scalarArray[1] >= colorLower[1]) && (scalarArray[1] <= colorUpper[1]) {
						if (scalarArray[2] >= colorLower[2]) && (scalarArray[2] <= colorUpper[2]) {
							coords["x"] = x
							coords["y"] = y
							break
						}
					}
				}
			}
		}
	}
	drawTargetBox(img, coords["x"], coords["y"])
	opencv.SaveImage(targetsImageFilename, img, nil)
	return coords
}