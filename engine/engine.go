package engine

import (
	"og/download"
	"og/pipeline"
	req "og/reqeuest"
	"og/response"
	"og/schedule"
)

type Engine struct {
	scheduler  chan *req.Request
	downloader chan *req.Request
	pipeliner  chan *response.Response
}

func New(scheduler chan *req.Request) *Engine {
	return &Engine{
		scheduler: scheduler,
	}
}

func (e *Engine) PushReq(r *req.Request) {

}

// Run: engine 运行
func (e *Engine) Run() {

	for {
		select {
		case req := <-e.scheduler:
			schedule.Process(req)
		case req := <-e.downloader:
			download.Process(req)
		case resp := <-e.pipeliner:
			pipeline.Process(resp)
		default:
			// log.Println("engine process")
		}

	}
}
