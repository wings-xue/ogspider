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
	scheduler       chan *req.Request
	db              *db.PgSQL
	PipelineSetting map[string]setting.PipelineSet
}

func New(scheduler chan *req.Request, db *db.PgSQL) *Pipeline {
	return &Pipeline{
		scheduler: scheduler,
		db:        db,
	}
}

// func (self *Pipeline) toPageReq(resp *response.Response) []*req.Request {
// 	out := make([]*req.Request, 0)
// 	_total := spider.FindKey(ogconfig.PageTotal, resp.Req.Datas).Value
// 	total, _ := strconv.Atoi(_total)
// 	for _, each := range spider.FindKey(ogconfig.StartURL, resp.Req.Datas).StartURL {
// 		if resp.URL == each {
// 			for i := 1; i < total; i++ {
// 				newPage := "page=" + strconv.Itoa(i)
// 				newURL := strings.Replace(each, "page=1", newPage, -1)
// 				q := *resp.Req
// 				q.URL = newURL
// 				q.UUID = spider.HashK(newURL)
// 				q.Seed = false
// 				out = append(out, &q)
// 			}
// 		}
// 	}
// 	return out
// }

func (self *Pipeline) saveResponse(resp *response.Response) {
	for key, pipelineSet := range self.PipelineSetting {
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

func OpenSpider(setting setting.CralwerSet, scheduler chan *req.Request, db *db.PgSQL) *Pipeline {
	return New(scheduler, db)
}
