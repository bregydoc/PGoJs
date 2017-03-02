package processing

import "github.com/gopherjs/gopherjs/js"

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