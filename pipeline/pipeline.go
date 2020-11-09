package pipeline

import (
	"log"
	"og/db"
	req "og/reqeuest"
	"og/response"
	"og/setting"
	"og/spider"
)

func (self *Pipeline) Process(resp *response.Response) {
	go func() {

		rst := req.ToCrawlerRst(resp.Req)
		rst["req_id"] = resp.Req.URL
		tablename := spider.FindKey(setting.TableName, spider.Zhaotoubiao()).Value

		m := self.db.Conn.Model(&rst).TableExpr(tablename).OnConflict("(" + setting.CrawlerRstKey + ")" + " DO UPDATE")
		for key, _ := range rst {
			m.Set(key + "=EXCLUDED." + key)
		}
		_, err := m.Insert()
		if err != nil {
			panic(err)
		}
		self.scheduler <- resp.Req
	}()

	log.Printf("[pipeline] %s status code : %d\n", resp.Req.URL, resp.StatusCode)
}

type Pipeline struct {
	scheduler chan *req.Request
	db        *db.PgSQL
}

func New(scheduler chan *req.Request, db *db.PgSQL) *Pipeline {
	return &Pipeline{
		scheduler: scheduler,
		db:        db,
	}
}
