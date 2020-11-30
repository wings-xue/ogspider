package schedule

import (
	"log"
	"og/context"
	"og/db"
	"og/filter"
	req "og/reqeuest"
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

func (self *Schedule) Minus() {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.WorkNum--
}

func (self *Schedule) Inc() {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.WorkNum++
}

func (self *Schedule) WorkLen() int {
	self.mu.Lock()
	defer self.mu.Unlock()
	return self.WorkNum
}

func (self *Schedule) Process(req *req.Request) {
	// s.manager = append(s.manager, req)
	if !self.filter.Contains(req.URL + strconv.Itoa(req.Retry)) {
		self.manager.Push(req)
	}
	if req.Seed {
		self.Minus()
	}
	log.Printf("[scheduler] queue request: %d, worker: %d", self.manager.Len(), self.WorkLen())
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
			self.Inc()
			req := self.manager.Pop()
			req.Seed = true
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
