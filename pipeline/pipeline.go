package pipeline

import (
	"log"
	req "og/reqeuest"
	"og/response"
)

func (self *Pipeline) Process(resp *response.Response) {
	go func() {
		self.scheduler <- resp.Req
	}()

	log.Printf("[pipeline] %s status code : %d\n", resp.Req.URL, resp.StatusCode)
}

type Pipeline struct {
	scheduler chan *req.Request
}

func New(scheduler chan *req.Request) *Pipeline {
	return &Pipeline{
		scheduler: scheduler,
	}
}
