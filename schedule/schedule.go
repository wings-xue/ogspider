package schedule

import (
	"log"
	"og/context"
	"og/db"
	"og/filter"
	req "og/reqeuest"
	"og/setting"
	"strconv"
	"sync"
)

type Schedule struct {
	manager *context.Manager
	WorkNum int
	// Schedule将request传入engine,后续engine获取对象，给pipeline
	downloader chan *req.Request
	filter     *filter.Bloom
	mu         sync.Mutex
}

func (self *Schedule) Minus(req *req.Request) {
	self.mu.Lock()
	defer self.mu.Unlock()

	if req.Seed {
		self.WorkNum--
	}

}

func (self *Schedule) Inc(req *req.Request) {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.WorkNum++
	req.Seed = true
}

func (self *Schedule) WorkLen() int {
	self.mu.Lock()
	defer self.mu.Unlock()
	return self.WorkNum
}

func (self *Schedule) Process(req *req.Request) {
	// s.manager = append(s.manager, req)

	if !self.filter.Contains(req.URL + strconv.Itoa(req.Retry) + req.UpdateDate.Format("2001-01-01 01:01:01")) {

		if req.Retry < setting.Retry {
			self.manager.Push(req)
			log.Printf("[scheduler] queue request: %d, worker: %d", self.manager.Len(), self.WorkLen()-1)
		}
	}
	self.Minus(req)

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

func (self *Schedule) LoopDispatch() {
	for {

		if self.manager.Len() > 0 && self.WorkLen() < 20 {
			req := self.manager.Pop()
			self.Inc(req)

			self.downloader <- req
		}
	}
}

func OpenSpider(downloader chan *req.Request, db *db.PgSQL) *Schedule {
	manager := context.New()
	filter := filter.New(filter.BLOOMSIZE)

	return &Schedule{
		downloader: downloader,
		manager:    manager,
		filter:     filter,
	}
}
