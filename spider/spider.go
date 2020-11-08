package spider

import (
	"log"
	"og/db"
	req "og/reqeuest"
)

type Spider struct {
	scheduler chan *req.Request
	database  *db.PgSQL
}

func New(scheduler chan *req.Request, database *db.PgSQL) *Spider {
	return &Spider{
		scheduler: scheduler,
		database:  database,
	}
}

func ReadMemory() []*req.Request {
	return ToSpider(Zhaotoubiao())

}

func ReadDB() []*req.Request {
	return []*req.Request{req.New("")}
}

func (s *Spider) Run() {
	request := ReadMemory()
	log.Printf("[spider] init spider :%d", len(request))
	for _, r := range request {
		db.New().Update(r, true)
		s.scheduler <- r
	}
}
