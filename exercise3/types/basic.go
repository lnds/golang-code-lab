package types

type Course int

const (
	North Course = iota
	East
	South
	West
)

func (c Course) String() string {
	switch c {
	case North:
		return "north"
	case East:
		return "east"
	case South:
		return "south"
	case West:
		return "west"
	default:
		return "undefined direction"
	}
}
