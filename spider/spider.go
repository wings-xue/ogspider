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

func (s *Spider) CreateTable(r *req.Request, db *db.PgSQL) {
	db.Conn.Exec(req.ToTableSchema("zhaotoubiao", r))
}

func (s *Spider) Run() {
	request := ReadMemory()
	log.Printf("[spider] init spider :%d", len(request))
	for _, r := range request {
		pgsql := db.New()
		pgsql.Update(r, true)
		s.CreateTable(r, pgsql)
		s.scheduler <- r
	}
}
