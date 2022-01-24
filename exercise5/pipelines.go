package main

import "fmt"

func Counter(limit int, out chan<- int) {
	for x := 0; x < limit; x++ {
		out <- x
	}
	close(out)
}

func Squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func Printer(in  <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go Counter(100, naturals)
	go Squarer(naturals, squares)
	
	Printer(squares)
}
