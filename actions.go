package main

import (
	"fmt"
	//"strings"
	//"strconv"
	//"github.com/go-vgo/robotgo"
)


//FIXME
func doAction(points map[string] string, mainBoxFilename string, color *Color) {
	//mp := strings.Split(points["mp"], ":")
	//mana, _ := strconv.ParseInt(mp[0], 10, 32)
	if points["hp"] == "Activitie" {
		//robotgo.KeyTap("f12")
		fmt.Println("Sitting")
	} else {
		attackTarget(mainBoxFilename, color)
	}
}

func attackTarget(filename string, color *Color) {
	coords := findTarget(filename, color)
	fmt.Println(coords)
}