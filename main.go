package main

import (
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	coords := readCoordsConf("coords.json")
	for {
		createScreenshot("ss.png")
		boxes := getPoints("ss.png", &coords)
		//fmt.Println(recognizeText(boxes))
		doAction(recognizeText(boxes))
		time.Sleep(1 * time.Second)
	}
}
