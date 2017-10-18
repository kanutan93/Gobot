package main

import (
	"time"
	"fmt"
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
		hpBox, mpBox := getPointsBoxes("ss.png", &coords)
		boxes := map[string]string{"hp": hpBox, "mp": mpBox}
		fmt.Println(recognizeText(boxes))
		time.Sleep(1 * time.Second)
	}
}
