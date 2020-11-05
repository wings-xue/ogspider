package main

import (
	"og/engine"
	req "og/reqeuest"
)

func crawler() {
	scheduler := make(chan *req.Request)
	downloader := make(chan *req.Request)

	engine := engine.New(scheduler, downloader)
	engine.Run()
}

func main() {
	crawler()
}
