package schedule

import (
	"log"
	"og/context"
	"og/filter"
	req "og/reqeuest"
)

type Schedule struct {
	manager *context.Manager

	// Schedule将request传入engine,后续engine获取对象，给pipeline
	downloader chan *req.Request
	filter     *filter.Bloom
}

func (self *Schedule) Process(req *req.Request) {
	// s.manager = append(s.manager, req)
	if !self.filter.Contains(req.URL) {
		self.manager.Push(req)
	}
	log.Printf("[scheduler] queue have request count is: %d", self.manager.Len())
}

func New(downloader chan *req.Request) *Schedule {
	manager := context.New()
	filter := filter.New(filter.BLOOMSIZE)
	return &Schedule{
		downloader: downloader,
		manager:    manager,
		filter:     filter,
	}
}

func (self *Schedule) Run() {
	for {
		if self.manager.Len() > 0 {
			req := self.manager.Pop()
			self.downloader <- req
		}
	}
}
