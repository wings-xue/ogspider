package schedule

import (
	"log"
	req "og/reqeuest"
	"time"
)

type Schedule struct {
	c chan *req.Request
}

func Process(*req.Request) {
	log.Println("[schedule] msg: add, req: Request")
}

func New(c chan *req.Request) *Schedule {
	return &Schedule{
		c: c,
	}
}

func (s *Schedule) Run() {
	for {
		s.c <- req.New("")
		time.Sleep(3 * time.Second)
	}
}
