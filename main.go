package main

import (
	"og/engine"
	req "og/reqeuest"
	"og/response"
)

func crawler() {
	scheduler := make(chan *req.Request)
	downloader := make(chan *req.Request)
	pipeliner := make(chan *response.Response)

	engine := engine.New(scheduler, downloader, pipeliner)
	engine.Run()
}

func main() {
	crawler()
}
