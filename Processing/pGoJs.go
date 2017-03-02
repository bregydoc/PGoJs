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

func CreateCanvas(width, height int) {
	pG.Call("createCanvas", width, height)
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
Missing some important methods. From Limit, please view the p5js reference
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
