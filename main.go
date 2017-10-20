package main

import (
	"time"
	"runtime"
)

const (
	coordsConfig = "coords.json"
	mainScreenshot = "ss.png"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	coords := readCoordsConf(coordsConfig)
	for {
		createScreenshot(mainScreenshot)
		points := getPoints(mainScreenshot, &coords)
		doAction(recognizePoints(points), "123.png") //FIXME поменять на mainScreenshot
		time.Sleep(1 * time.Second)
	}
}
