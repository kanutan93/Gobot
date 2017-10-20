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
	color := readColorConf(colorConf)
	for {
		createScreenshot(mainScreenshot)
		points := getPoints(mainScreenshot, &coords)
		doAction(recognizePoints(points), "test2.jpg", &color) //FIXME поменять на mainScreenshot
		time.Sleep(1 * time.Second)
	}
}
