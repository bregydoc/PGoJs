//constAndVarHeader.go
package processing

import "github.com/gopherjs/gopherjs/js"

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
