package main

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/go-vgo/robotgo"
)


//FIXME
func doAction(points map[string] string) {
	mp := strings.Split(points["mp"], ":")
	mana, _ := strconv.ParseInt(mp[0], 10, 32)
	if points["hp"] == "Activitie" || mana < 13 {
		robotgo.KeyTap("f12")
		fmt.Println("Sitting")
	} else {
		robotgo.KeyTap("f1")
	}

}