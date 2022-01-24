package turtles

import (
	"fmt"
	. "ediaz.dev/simulator/types"
)

type Turtle struct {
	name      string
	x, y      int
	distance  int
	direction Course
}

func NewTurtle(name string) *Turtle {
	return &Turtle{
		name:      name,
		x:         0,
		y:         0,
		distance:  0,
		direction: North,
	}
}

func (t *Turtle) Name() string {
	return t.name
}

func (t *Turtle) Reset() {
	t.x = 0
	t.y = 0
	t.distance = 0
}

func (t *Turtle) Forward(distance int) {
	t.distance += distance
	switch t.direction {
	case North:
		t.y--
	case East:
		t.x++
	case South:
		t.y++
	case West:
		t.x--
	}
}

func (t *Turtle) Left() {
	switch t.direction {
	case North:
		t.direction = West
	case East:
		t.direction = North
	case South:
		t.direction = East
	case West:
		t.direction = South
	}
}

func (t *Turtle) Right() {
	switch t.direction {
	case North:
		t.direction = East
	case East:
		t.direction = South
	case South:
		t.direction = West
	case West:
		t.direction = North
	}
}

func (t *Turtle) Odometer() int {
	return t.distance
}

func (t *Turtle) Direction() Course {
	return t.direction
}

func (t *Turtle) Status() string {
	return fmt.Sprintf("The turtle %s is head %s, and has traveled %d units", t.Name(), t.Direction(), t.Odometer())
}
