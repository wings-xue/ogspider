package main

import (
	"fmt"
	"os"
	"time"
)

func addENV() {
	var Chrome = `C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe`
	os.Setenv("rob", "show,bin="+Chrome)

}

func main() {
	go func() {
		panic("adfaf")
	}()

	go func() {
		fmt.Println("b")
	}()
	time.Sleep(time.Second * 3)
	fmt.Println("asdfasd")
}
