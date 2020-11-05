package schedule

import (
	"log"
	req "og/reqeuest"
)

type Schedule struct {
	manager []*req.Request

	// Schedule将request传入engine,后续engine获取对象，给pipeline
	downloader chan *req.Request
}

func (s *Schedule) Process(req *req.Request) {
	s.manager = append(s.manager, req)
	log.Println("[schedule] msg: add, req: Request")
}

func New(downloader chan *req.Request) *Schedule {
	manager := make([]*req.Request, 0)
	return &Schedule{
		downloader: downloader,
		manager:    manager,
	}
}

func (s *Schedule) Run() {
	for {
		// time.Sleep(3 * time.Second)
		// if len(s.manager) > 0 {
		// 	s.downloader <- s.manager[0]
		// 	time.Sleep(3 * time.Second)
		// }

	}
}
