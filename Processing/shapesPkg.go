package processing

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
