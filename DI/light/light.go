package light

import "fmt"

//DI&抽象化

// 2種類の電球(Incandescent,LED)をsocketで抽象化
type LightSocket interface {
	LightUp() string
}

type Incandescent struct{}

func (*Incandescent) LightUp() string {
	return "フィラメントが光る"
}

// 追加のLED
type LedLight struct{}

func (*LedLight) LightUp() string {
	return "LEDが光る"
}

type Room struct {
	LightOne LightSocket
	LightTwo LightSocket
}

func (r *Room) SwitchOnOne() {
	fmt.Println("1番照明:", r.LightOne.LightUp())
}

func (r *Room) SwitchOnTwo() {
	fmt.Println("2番照明:", r.LightTwo.LightUp())
}

func NewRoom(lightOne, lightTwo LightSocket) *Room {
	room := &Room{
		LightOne: lightOne,
		LightTwo: lightTwo,
	}
	return room
}
