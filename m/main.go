package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()

	go func() {
		for {
			c <- 3
		}

	}()
	for {
		select {
		case i := <-c:
			fmt.Println(i)
		default:

		}
	}

}
