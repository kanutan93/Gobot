package main
import (
	"time"
	"runtime"
	"fmt"
)

const (
	coordsConfig = "coords.json"
	mainImageFilename = "test.jpg"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)
	coordsConfig := readCoordsConf(coordsConfig)
	colorConfig := readColorConf(colorConf)
	points := getText(mainImageFilename, &coordsConfig)
	coords := findTarget(mainImageFilename, &colorConfig.Green)
	fmt.Println(coords, recognizePoints(points))
	time.Sleep(1 * time.Second)
}
