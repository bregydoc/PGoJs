# PGoJs


> Processing with Go, using p5Js as core


PGoJs is a binding/port from p5Js using gopherJs, the idea is create sketchs in web using Golang but easy and fast like processing framework.

## How to use

PGoJs need to [gopherJs](https://github.com/gopherjs/gopherjs) builder for work, you make sure you have it.

#### Create a index.html file
You need a HTML file where you will embedded the js generated by PGo Js using gopherjs, your html index file has look like this:
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>PGoJs</title>

</head>
<body>
    <script  src="https://cdnjs.cloudflare.com/ajax/libs/p5.js/0.5.7/p5.js"></script>
    <script src="sketch.js"></script>
</body>
</html>
```

#### Write your sketch in Go
You need import the PGoJs library, I recommended use this with different namespace (like "p")
```go
package main

import (
	p "./Processing"
)

func setup() {
	p.CreateCanvas(600, 600)
	p.Background(230)
}

func draw() {
	p.StrokeWeight(4)
	if !p.MouseIsPressed {
		p.NoStroke()
	} else {
		p.Stroke("rgba(139,195,74 ,1)")
	}
	p.Line(p.PmouseX, p.PmouseY, p.MouseX, p.MouseY)
}

func main() {
	p.Setup = setup
	p.Draw = draw

	p.LaunchApp()
}
```

### Build your .go file to .js file using gopherJs
First, make sure you have [gopherJs](https://github.com/gopherjs/gopherjs), if you don't have, install with:
```
go get -u github.com/gopherjs/gopherjs
```
For build the sketch.js file only execute this line:
```
gopherjs build nameOfSketch.go -o sketch.js
```
Or:
```bash
export GOPATH=$HOME/goWork #The path of your GOPATH
$GOPATH/bin/gopherjs build nameOfSketch.go -o sketch.js
```
I usually create an .sh file with these parameters, and build in only one step.