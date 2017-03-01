/*
Port from p5Js to golang using gopherJs
In this version, the framework implement native types from p5Js (using gopherJs.Object struct),
in the future, I will create types and structures entirely write in golang for improve the performance

WORK IN PROGRESS (10% completed)

Author: Bregy Malpartida Ramos
Date: 03/01/2017

*/

package processing

import (
	"github.com/gopherjs/gopherjs/js"
	"reflect"
)

var pG *js.Object
var Setup func()
var Draw func()

// Constants in p5Js: (In go are only variables)
var HALF_PI float64
var PI float64
var QUARTER_PI float64
var TAU float64
var TWO_PI float64

//Modes for any functions:
var OPEN string //Is a String
var CHORD string
var PIE string

var RGB string
var HSB string

var LEFT string
var RIGHT string
var CENTER string

// Cast of different variables
var MouseX int
var MouseY int
var PmouseX int
var PmouseY int
var WinMouseX int
var WinMouseY int
var PwinMouseX int
var PwinMouseY int
var MouseButton string
var MouseIsPressed bool

var Width int
var Height int

func CreateCanvas(width, height int) {
	pG.Call("createCanvas", width, height)
}

////////////////////////////////////////////////////////////
//     					COLOR

//Creating and reading:

func Color(params ...interface{}) *js.Object {
	switch len(params) {
	case 1:
		return pG.Call("color", params[0])
	case 2:
		return pG.Call("color", params[0], params[1])
	case 3:
		return pG.Call("color", params[0], params[1], params[2])
	case 4:
		return pG.Call("color", params[0], params[1], params[2], params[3])
	default:
		println("Error in Color function (1)")
		return nil
	}
}

func Alpha(color *js.Object) int {
	return pG.Call("alpha", color).Int()
}

func Blue(color *js.Object) int {
	return pG.Call("blue", color).Int()
}

func Brightness(color *js.Object) int {
	return pG.Call("brightness", color).Int()
}

func Green(color *js.Object) int {
	return pG.Call("green", color).Int()
}

func Hue(color *js.Object) int {
	return pG.Call("hue", color).Int()
}

func LerpColor(from interface{}, to interface{}, amt float64) *js.Object {
	return pG.Call("lerpColor", from, to, amt)
}

func Lightness(color *js.Object) int {
	return pG.Call("lightness", color).Int()
}

func Red(color *js.Object) int {
	return pG.Call("red", color).Int()
}

func Saturation(color *js.Object) int {
	return pG.Call("saturation", color).Int()

}

//Setting:

func Background(values ...interface{}) {

	switch len(values) {
	case 1:
		pG.Call("background", values[0])
		break

	case 2:

		pG.Call("background", values[0].(int), values[1].(int))
		break

	case 3:
		pG.Call("background", values[0].(int), values[1].(int), values[2].(int))
		break

	default:
		println("Error in background function (2)...")
		return
	}
}

func Clear() {
	pG.Call("clear")
}

func ColorMode(mode string, maxValues ...int) {
	switch len(maxValues) {
	case 0:
		pG.Call("colorMode", mode)
		break
	case 1:
		pG.Call("colorMode", mode, maxValues[0])
		break
	case 2:
		pG.Call("colorMode", mode, maxValues[0], maxValues[1])
		break
	case 3:
		pG.Call("colorMode", mode, maxValues[0], maxValues[1], maxValues[2])
		break
	case 4:
		pG.Call("colorMode", mode, maxValues[0], maxValues[1], maxValues[2], maxValues[3])
		break
	default:
		println("Error in colorMode (1)")
		return
	}

}

func Fill(firstValue interface{}, extraValues ...float64) {
	switch len(extraValues) {
	case 0:
		/* Darwin!
		typeOfFirstValue := reflect.TypeOf(firstValue).Name()
		if typeOfFirstValue == reflect.Int.String() || typeOfFirstValue == reflect.Float64.String(){
			pG.Call("fill", firstValue.(float64))
		}
		*/
		pG.Call("fill", firstValue)
		break
	case 1:
		pG.Call("fill", firstValue, extraValues[0])
		break
	case 2:
		pG.Call("fill", firstValue, extraValues[0], extraValues[1])
		break
	case 3:
		pG.Call("fill", firstValue, extraValues[0], extraValues[1], extraValues[2])
		break
	}
}

func NoFill() {
	pG.Call("noFill")

}

func NoStroke() {
	pG.Call("noStroke")

}

func Stroke(firstValue interface{}, extraValues ...int) {
	switch len(extraValues) {
	case 0:
		pG.Call("stroke", firstValue)
		break
	case 1:
		pG.Call("stroke", firstValue, extraValues[0])
		break
	case 2:
		pG.Call("stroke", firstValue, extraValues[0], extraValues[1])
		break
	case 3:
		pG.Call("stroke", firstValue, extraValues[0], extraValues[1], extraValues[2])
		break
	}
}

func StrokeWeight(weight int) {
	pG.Call("strokeWeight", weight)
}

////////////////////////////////////////////////////////////
//     					SHAPES

// 2D PRIMITIVES:
func Arc(a, b, c, d, start, stop interface{}, mode string) {
	pG.Call("arc", a, b, c, d, start, stop, mode)
}

func Ellipse(x, y, w interface{}, h ...interface{}) {
	switch len(h) {
	case 0:
		pG.Call("ellipse", x, y, w)
		break
	case 1:
		pG.Call("ellipse", x, y, w, h[0])
		break
	default:
		println("Error in ellipse funciton (1)")
		return
	}
}

func Line(x1, y1, x2, y2 interface{}) {
	pG.Call("line", x1, y1, x2, y2)
}

func Point(x, y interface{}) {
	pG.Call("point", x, y)
}

func Quad(x1, y1, x2, y2, x3, y3, x4, y4 interface{}) {
	pG.Call("quad", x1, y1, x2, y2, x3, y3, x4, y4)
}

func Rect(x, y, w, h interface{}, extraParameters ...interface{}) {
	switch len(extraParameters) {
	case 0:
		pG.Call("rect", x, y, w, h)
		break
	case 1:
		pG.Call("rect", x, y, w, h, extraParameters[0])
		break
	case 2:
		pG.Call("rect", x, y, w, h, extraParameters[0], extraParameters[1])
		break
	case 3:
		pG.Call("rect", x, y, w, h, extraParameters[0], extraParameters[1], extraParameters[2])
		break
	case 4:
		pG.Call("rect", x, y, w, h, extraParameters[0], extraParameters[1], extraParameters[2], extraParameters[3])
		break

	default:
		println("Error in Rect funciton (1)")
		return
	}
}

func Triangle(x1, y1, x2, y2, x3, y3 interface{}) {
	pG.Call("triangle", x1, y1, x2, y2, x3, y3)
}

////////////////////////////////////////////////////////////
//     					EVENTS

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

////////////////////////////////////////////////////////////

///////////////////// VECTOR CLASS //////////////////////////

type PVector struct {
	data *js.Object
	X    float64
	Y    float64
	Z    float64
}

func CreateVector(cord ...float64) *PVector {
	var r *js.Object

	switch len(cord) {
	case 0:
		r = pG.Call("createVector")
		break
	case 1:
		r = pG.Call("createVector", cord[0])
		break
	case 2:
		r = pG.Call("createVector", cord[0], cord[1])
		break
	case 3:
		r = pG.Call("createVector", cord[0], cord[1], cord[2])
		break
	default:
		r = nil
		println("Error in createVector function (1)")
		break
	}
	return &PVector{data: r, X: r.Get("x").Float(), Y: r.Get("y").Float(), Z: r.Get("z").Float()}
}

func (v *PVector) Add(x interface{}, other ...interface{}) *PVector {
	var r *js.Object

	switch len(other) {
	case 0:
		if reflect.TypeOf(x).Kind() == reflect.Float64 || reflect.TypeOf(x).Kind() == reflect.Int {
			if reflect.TypeOf(x).Kind() == reflect.Float64 {
				r = v.data.Call("add", x.(float64))
			} else if reflect.TypeOf(x).Kind() == reflect.Int {
				r = v.data.Call("add", x.(int))
			}
		} else {
			r = v.data.Call("add", x.(*PVector).data)
		}
		break
	case 1:
		r = v.data.Call("add", x.(float64), other[0].(float64))
		break

	case 2:
		r = v.data.Call("add", x.(float64), other[0].(float64), other[1].(float64))
		break
	default:
		r = nil
		println("Error in add vector method (1)")
		break
	}
	v.X = r.Get("x").Float()
	v.Y = r.Get("y").Float()
	v.Z = r.Get("z").Float()

	return &PVector{data: r, X: r.Get("x").Float(), Y: r.Get("y").Float(), Z: r.Get("z").Float()}

}

func (v *PVector) ToString() string {
	r := v.data.Call("toString")
	return r.String()
}

func (v *PVector) Set(cord ...float64) *PVector {
	var r *js.Object

	switch len(cord) {
	case 1:
		r = v.data.Call("set", cord[0])
		break
	case 2:
		r = v.data.Call("set", cord[0], cord[1])
		break
	case 3:
		r = v.data.Call("set", cord[0], cord[1], cord[2])
		break
	default:
		r = nil
		println("Error in set (PVector) function (1)")
		break
	}

	v.X = r.Get("x").Float()
	v.Y = r.Get("y").Float()
	v.Z = r.Get("z").Float()

	return &PVector{data: r, X: r.Get("x").Float(), Y: r.Get("y").Float(), Z: r.Get("z").Float()}
}

func (v *PVector) Copy() *PVector {
	r := v.data.Call("copy")
	return &PVector{data: r, X: r.Get("x").Float(), Y: r.Get("y").Float(), Z: r.Get("z").Float()}
}

func (v *PVector) Sub(x interface{}, other ...interface{}) *PVector {
	var r *js.Object

	switch len(other) {
	case 0:
		if reflect.TypeOf(x).Kind() == reflect.Float64 || reflect.TypeOf(x).Kind() == reflect.Int {
			if reflect.TypeOf(x).Kind() == reflect.Float64 {
				r = v.data.Call("sub", x.(float64))
			} else if reflect.TypeOf(x).Kind() == reflect.Int {
				r = v.data.Call("sub", x.(int))
			}
		} else {
			r = v.data.Call("sub", x.(*PVector).data)
		}
		break
	case 1:
		r = v.data.Call("sub", x.(float64), other[0].(float64))
		break

	case 2:
		r = v.data.Call("sub", x.(float64), other[0].(float64), other[1].(float64))
		break
	default:
		r = nil
		println("Error in sub vector method (1)")
		break
	}
	v.X = r.Get("x").Float()
	v.Y = r.Get("y").Float()
	v.Z = r.Get("z").Float()

	return &PVector{data: r, X: r.Get("x").Float(), Y: r.Get("y").Float(), Z: r.Get("z").Float()}

}

func (v *PVector) Mult(val float64) *PVector {
	r := v.data.Call("mult", val)
	v.X = r.Get("x").Float()
	v.Y = r.Get("y").Float()
	v.Z = r.Get("z").Float()
	return &PVector{data: r, X: r.Get("x").Float(), Y: r.Get("y").Float(), Z: r.Get("z").Float()}

}

func (v *PVector) Div(val float64) *PVector {
	r := v.data.Call("div", val)
	v.X = r.Get("x").Float()
	v.Y = r.Get("y").Float()
	v.Z = r.Get("z").Float()
	return &PVector{data: r, X: r.Get("x").Float(), Y: r.Get("y").Float(), Z: r.Get("z").Float()}

}

func (v *PVector) Mag() float64 {
	r := v.data.Call("mag")
	return r.Float()
}

func (v *PVector) MagSq() float64 {
	r := v.data.Call("magSq")
	return r.Float()
}

func (v *PVector) Dot(v2 *PVector) float64 {
	r := v.data.Call("dot", v2.data)
	return r.Float()
}

func (v *PVector) Cross(v2 *PVector) *PVector {
	r := v.data.Call("cross", v2.data)
	return &PVector{data: r, X: r.Get("x").Float(), Y: r.Get("y").Float(), Z: r.Get("z").Float()}
}

func (v *PVector) Dist(v2 *PVector) float64 {
	r := v.data.Call("dist", v2.data)
	return r.Float()
}

func (v *PVector) Normalize() *PVector {
	r := v.data.Call("normalize")
	v.X = r.Get("x").Float()
	v.Y = r.Get("y").Float()
	v.Z = r.Get("z").Float()
	return &PVector{data: r, X: r.Get("x").Float(), Y: r.Get("y").Float(), Z: r.Get("z").Float()}
}

/*
Missing some important methods. From Limit, view the p5js reference
*/

func Random2D() *PVector {
	r := js.Global.Get("p5").Get("Vector").Call("random2D")
	return &PVector{data: r, X: r.Get("x").Float(), Y: r.Get("y").Float(), Z: r.Get("z").Float()}
}

func Random3D() *PVector {
	r := js.Global.Get("p5").Get("Vector").Call("random3D")
	return &PVector{data: r, X: r.Get("x").Float(), Y: r.Get("y").Float(), Z: r.Get("z").Float()}
}

/////////////////////////////////////////////////////////////

func LaunchApp() {

	var sketch = func(p *js.Object) {
		pG = p

		////////////////// CONSTANTS ///////////////
		HALF_PI = pG.Get("HALF_PI").Float()
		PI = pG.Get("PI").Float()
		QUARTER_PI = pG.Get("QUARTER_PI").Float()
		TAU = pG.Get("TAU").Float()
		TWO_PI = pG.Get("TWO_PI").Float()

		OPEN = pG.Get("OPEN").String()
		CHORD = pG.Get("CHORD").String()
		PIE = pG.Get("PIE").String()

		RGB = pG.Get("RGB").String()
		HSB = pG.Get("HSB").String()

		LEFT = pG.Get("LEFT").String()
		CENTER = pG.Get("CENTER").String()
		RIGHT = pG.Get("RIGHT").String()
		///////////////////////////////////////////

		p.Set("setup", func() {
			Setup()
			//p.Call("createCanvas", 640, 480)
		})

		p.Set("draw", func() {
			////////// VARIABLES ///////////
			MouseX = p.Get("mouseX").Int()
			MouseY = p.Get("mouseY").Int()
			PmouseX = p.Get("pmouseX").Int()
			PmouseY = p.Get("pmouseY").Int()
			WinMouseX = p.Get("winMouseX").Int()
			WinMouseY = p.Get("winMouseY").Int()
			PwinMouseX = p.Get("pwinMouseX").Int()
			PwinMouseY = p.Get("pwinMouseY").Int()
			MouseButton = p.Get("mouseButton").String()
			MouseIsPressed = p.Get("mouseIsPressed").Bool()

			Width = p.Get("width").Int()
			Height = p.Get("height").Int()

			////////////////////////////////

			Draw()

			//p.Call("background", 0)
		})

		p.Set("mouseMoved", func() {
			if MouseMoved != nil {
				MouseMoved()
			}

		})

		p.Set("mouseDragged", func() {
			if MouseDragged != nil {
				MouseDragged()
			}

		})

		p.Set("mousePressed", func() {
			if MousePressed != nil {
				MousePressed()
			}
		})

		p.Set("mouseReleased", func() {
			if MouseReleased != nil {
				MouseReleased()
			}
		})

		p.Set("mouseClicked", func() {
			if MouseClicked != nil {
				MouseClicked()
			}
		})
		p.Set("mouseWheel", func(event interface{}) {
			if MouseWheel != nil {
				MouseWheel(Event{data: event.(map[string]interface{})})
			}
		})

	}

	js.Global.Get("p5").New(sketch)

}
