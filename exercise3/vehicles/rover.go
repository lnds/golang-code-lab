package vehicles

import (
	"ediaz.dev/simulator/types"
)

type Rover interface {
	Name() string
	Reset()
	Forward(distance int)
	Left()
	Right()
	Direction() types.Course
	Odometer() int
	Status() string
}
