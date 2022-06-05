package lightBad

import "fmt"

//DIしてない抽象化もしてない

type Incandescent struct{}

func (i *Incandescent) LightUp() string {
	return "フィラメントが光る"
}

type Room struct {
	LightOne Incandescent
	LightTwo Incandescent
}

func (r *Room) SwitchOnOne() {
	fmt.Println("1番証明", r.LightOne.LightUp())
}
func (r *Room) SwitchOnTwo() {
	fmt.Println("2番証明", r.LightTwo.LightUp())
}
