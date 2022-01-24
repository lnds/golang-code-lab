# Clock Server

In this exercise, you will learn about Go Routines.

The file `clock.go` contains a simple TCP server that respond the current time of the server.

You can run it using:

    % go run clock.go

In other terminal you can do:

    % nc localhost 8000
    13:58:54
    13:58:55
    13:58:56
    13:58:57
    ^C


## The Challenge

If you note this program accept one connection at time.
Change this program to run concurrently using go routines, so it can handle many simultaneous connections.


Time expected for this exercise: 5 minutes

## Notes

(*) This example is based on chapter 8 of The Go Programming Language, by Donovan and Kernighan