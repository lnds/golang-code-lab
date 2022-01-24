# Vehicles

In this exercise, you will learn about interfaces and duck typing

This is a more complex example, that implements a vehicle simulator.

We created this project using this command:

    % go mod init ediaz.dev/simulator

The string `"ediaz.dev/simulator"` is the "root" for all the packages.

The file structure of this project is this:

    exercise3 +
              +- main.go
              +- go.mod
              +- turtles +
              |          +- turtle.go
              |
              +- types +
              |        +- basic.go
              |
              +- utils + 
              |        +- commands.go
              |        +- scanner.go
              |
              +- vehicles +
                          +- rover.go

For example, to import the utils package you must use this syntax:

    import ""ediaz.dev/simulator/utils"

for example to use the `ReadLines()` from utils you must do this:

        import ""ediaz.dev/simulator/utils" 

        ...

        lines := utils.ReadLines(filename)


## The Rover Interface 

The file `vehicles/rover.go` define this interface

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
    
This is a contract for a simple vehicle type. And is used by the simulator to operate different types of vehicles.

## The Turtle

The file `turtles/turtle.go` implements a Turtle, like that used on the [Logo Programming Language](https://en.wikipedia.org/wiki/Logo_(programming_language)), that implements the Rover Interface.

## The Simulator

The `main.go` file define the function `Simulator()`, you can see that this function implements an interpreter of the commands received.

Every commando has this structure:

    command [arg]

For example:

    forward 10
    left
    right
    forward -5
    reset

The simulator receive a list of commands, parse them using the `ParseCommand()` function defined in `utils/commands.go`, and then call the rover received as parameter to actually execute the command.

After all commands are executed, the Simulator shows the rover status.

# Challenge

Implements an Uber object, modify 

In `main.go`file, if you look at lines 40 to 42 you will see this:

    // uncomment this for exercise 3:
    // uber := uber.NewUber("X")
    // Simulator(lines, uber)

And in lines 24 to 26 you will see:

    // uncomment this for exercise 3
    //default:
    // rover.Unknown(cmd, arg)

To solve this exercise you must:

- Implements the Uber type in package `uber`
- Include the package in main.go
- Modify the Rover Interface to handle non standard commands using the `Unknown()` method

The Uber object implements additional commands, that are passed to the object using the `Unknown()` method. This commands are:

    - driver 1: assign a driver to an uber
    - driver 0: a driver leaves the uber
    - rider n: n riders take an uber
    - rider 0: all riders leaves the uber
    - rider -n: n riders leaves an uber

An uber cannot advance without a driver.
An uber make 10 dollars every time a rider take a ride.

The `Status()` method for a Uber object must inform, how many kms the uber has traveled, how many riders it rides, and how many dollars had made.


Time expected for this exercise: 45 minutes
