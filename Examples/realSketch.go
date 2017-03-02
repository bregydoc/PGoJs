package main

import (
	p "github.com/bregydoc/PGoJs/Processing"
)

type Ball struct {
	diameter float64
	position *p.PVector
	velocity *p.PVector
}

func newBall(x, y float64, diameter float64) *Ball{
	return &Ball{diameter: diameter, position:p.CreateVector(x,y), velocity:p.Random2D().Mult(10)}
}

func (ball *Ball) updateLogic() {
	if ((ball.position.X + ball.diameter/2) > float64(p.Width)) || ((ball.position.X - ball.diameter/2) < 0){
		ball.velocity = p.CreateVector(ball.velocity.X * -1, ball.velocity.Y)
	}
	if ((ball.position.Y + ball.diameter/2) > float64(p.Height)) || ((ball.position.Y - ball.diameter/2) < 0){
		ball.velocity = p.CreateVector(ball.velocity.X, ball.velocity.Y* -1)
	}

	ball.position.Add(ball.velocity)

}

func (ball *Ball) drawBall() {
	p.NoStroke()
	p.Fill("rgba(139,195,74 ,1)")
	p.Ellipse(ball.position.X, ball.position.Y, ball.diameter, ball.diameter)
}


var balls []*Ball

func setup() {
	p.CreateCanvas(600, 600)
	p.Background(230)
	balls = append(balls, newBall(200, 200, 50))

}

func draw() {
	p.Background(230)
	for _, ball := range balls {
		ball.updateLogic()
		ball.drawBall()
	}
}

func mousePressed() {
	balls = append(balls, newBall(200, 200, 50))
}

func main() {
	p.Setup = setup
	p.Draw = draw
	p.MousePressed = mousePressed

	p.LaunchApp()
}


