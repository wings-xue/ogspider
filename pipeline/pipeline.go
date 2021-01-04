package pipeline

import (
	"log"
	"og/db"
	req "og/reqeuest"
	"og/response"
	"og/setting"
)

func (self *Pipeline) Process(resp *response.Response) {
	log.Printf("[pipeline] %s status code : %d\n", resp.Req.URL, resp.StatusCode)

	for _, r := range resp.NewReq {
		self.sendReq(r)

	}
	// 3. save Crawler RST
	self.saveResponse(resp)
	// 4. handler old request: old -> chan ; save old

	self.sendReq(resp.Req)
}

type Pipeline struct {
	scheduler chan *req.Request
	db        *db.PgSQL
	Setting   setting.CrawlerSet
}

func New(setting setting.CrawlerSet, scheduler chan *req.Request, db *db.PgSQL) *Pipeline {
	return &Pipeline{
		scheduler: scheduler,
		db:        db,
		Setting:   setting,
	}
}

func (self *Pipeline) saveResponse(resp *response.Response) {
	for key, pipelineSet := range self.Setting.PipelineSetting {
		if resp.MatchBool(key) {
			rst := req.ToCrawlerRst(resp.Req)
			rst["req_id"] = resp.Req.URL

			self.db.Save(pipelineSet.SaveTable, rst)
		}
	}
}

func (self *Pipeline) sendReq(r *req.Request) {
	self.scheduler <- r
	self.saveReq(r)
}

func (self *Pipeline) saveReq(req *req.Request) {
	if req.Seed {
		self.db.MustUpdate(*req, true)
	} else {
		self.db.MustUpdate(*req, false)
	}
}

func OpenSpider(setting setting.CrawlerSet, scheduler chan *req.Request, db *db.PgSQL) *Pipeline {
	return New(setting, scheduler, db)
}
