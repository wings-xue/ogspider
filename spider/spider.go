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

func InitMemory() []*req.Request {
	return InitToSpider(Zhaotoubiao())
}

func InitDD() []*req.Request {
	return []*req.Request{req.New("")}
}

func (s *Spider) InitTable(r *req.Request, db *db.PgSQL) {
	tablename := FindKey(setting.TableName, Zhaotoubiao()).Value
	_, err := db.Conn.Exec(req.ToTableSchema(tablename, r))
	if err != nil {
		panic(err)
	}

}

func (s *Spider) Run() {
	request := InitMemory()
	log.Printf("[spider] init spider :%d", len(request))
	pgsql := db.New()
	for _, r := range request {

		pgsql.Update(r, true)
		s.InitTable(r, pgsql)
		s.scheduler <- r
	}
}
