package main

import "fmt"

func add(c chan int) {
	for {
		c <- 1
	}
}

func main() {
	c := make(chan int)
	go add(c)
	for {
		i := <-c
		fmt.Println(i)
	}
}
