package main

import (
	"ediaz.dev/simulator/turtles"
	"ediaz.dev/simulator/utils"
	. "ediaz.dev/simulator/vehicles"
	"fmt"
	"os"
	"strings"
)

func Simulator(commands []string, rover Rover) {
	for _, command := range commands {
		cmd, arg := utils.ParseCommand(command)
		switch strings.ToLower(cmd) {
		case "reset":
			rover.Reset()
		case "forward":
			rover.Forward(arg)
		case "left":
			rover.Left()
		case "right":
			rover.Right()
		// uncomment this for exercise 3
		//default:
		// rover.Unknown(cmd, arg)
		}
	}
	fmt.Println(rover.Status())
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: simulator instructions_file")
	} else {
		lines := utils.ReadLines(os.Args[1])
		turtle := turtles.NewTurtle("Tortoise")
		Simulator(lines, turtle)

		// uncomment this for exercise 3:
		// uber := uber.NewUber("X")
		// Simulator(lines, uber)
	}
}
