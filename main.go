package main

import (
	"og/download"
	"og/engine"
	"og/pipeline"
	req "og/reqeuest"
	"og/schedule"
	"og/spider"
)

func crawler() {
	scheduleChan := make(chan *req.Request)

	spider.Run()

	schedule := schedule.New(scheduleChan)
	go schedule.Run()

	pipeline := pipeline.New()
	go pipeline.Run()

	download := download.New()
	go download.Run()

	engine := engine.New(scheduleChan)
	engine.Run()
}

func main() {
	crawler()
}
