package spider

import (
	"log"
	"og/db"
	req "og/reqeuest"
	"og/setting"
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
	tablename := FindKey(setting.TableName, Zhaotoubiao()).Value
	_, err := db.Conn.Exec(req.ToTableSchema(tablename, r))
	if err != nil {
		panic(err)
	}

}

func (s *Spider) Run() {
	request := ReadMemory()
	log.Printf("[spider] init spider :%d", len(request))
	pgsql := db.New()
	for _, r := range request {

		pgsql.Update(r, true)
		s.CreateTable(r, pgsql)
		s.scheduler <- r
	}
}
