package main

import (
	"github.com/hiroki-kondo-git/interface/DI/light"
	lightBad "github.com/hiroki-kondo-git/interface/DI/lightNotDI"
)

func main() {
	//DIしてない抽象化もしてない
	myRoom := new(lightBad.Room)
	myRoom.SwitchOnOne() // 1番照明: フィラメントが光るよ!
	myRoom.SwitchOnTwo() // 2番照明: フィラメントが光るよ!

	//DI&抽象化

	lightOne := new(light.Incandescent) //注入するオブジェクト
	lightTwo := new(light.LedLight)     //注入するオブジェクト

	myNewRoom := light.NewRoom(lightOne, lightTwo) //DI

	myNewRoom.SwitchOnOne() // 1番照明: フィラメントが光るよ!
	myNewRoom.SwitchOnTwo() // 2番照明: LEDが光るよ!
}
