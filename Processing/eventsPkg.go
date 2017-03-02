package processing

////////////////////////////////////////////////////////////
//     					EVENTS

// MOUSE:
var MouseMoved func()
var MouseDragged func()
var MousePressed func()
var MouseReleased func()
var MouseClicked func()
var MouseWheel func(event Event)

type Event struct {
	data map[string]interface{}
}

func (event *Event) Delta() float64 { //Delta is a method in Go
	return event.data["delta"].(float64)
}