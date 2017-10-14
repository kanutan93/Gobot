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
	for {
		CreateScreenshot("ss.png")
		GetPointsBox("ss.png", "points-box.png")
		time.Sleep(1 * time.Second)
	}
}