# Pipelines

In this exercise, you will learn about channels

The program implements a pipeline of computations using channels.

A diagram of what this program does is this:


    +-----------+                   +-----------+              +-----------+
    |           |   0,1,2,3,...     |           |  0,1,4,9...  |           |
    |  Counter  |------------------>|  Squarer  |------------->|  Printer  |
    |           |    naturals       |           |  squares     |           |
    +-----------+                   +-----------+              +-----------+


Note that `Counter()` and `Squarer()` are started as go routines, but We call `Printer()` as a normal function.
This is necessary because, when the main function exits, all go routines stop. As Printer() has to wait until Squarer closes the channel, this ensures the program show the output.

## The Challenge

Implement a `Fibonacci` function that generates the Fibonacci sequence until a limit, given as parameter, i.e.:

    func Fibonacci(limit int, out chan<- int)

Then implement the function `Diff` that receives two input channels and put on an output channel the difference of the numbers received:

    func Diff(a, b <-chan int, out chan<- int)

Then determinate if the square growth faster or slower than the Fibonacci sequence.

Essentially, this exercise consist in implement this pipeline

    +-----------+                  +-----------+             
    |           |   0,1,2,3,...    |           |  0,1,4,9...  
    |  Counter  |----------------->|  Squarer  |---------------------\
    |           |   naturals       |           |  squares            |       +--------+     +-----------+
    +-----------+                  +-----------+                     |       |        |     |           |
                                                                     +------>|  Diff  |---->|  Printer  |
                                   +-----------+                     |       |        |     |           | 
                                   |           | 0,1,1,2,3,5,8,...   |       +--------+     +-----------+
                                   | Fibonacci |---------------------/       
                                   |           |   fibos
                                   +-----------+


Time expected for this exercise: 30 minutes

## Notes

(*) This example is based on chapter 8 of The Go Programming Language, by Donovan and Kernighan                                   