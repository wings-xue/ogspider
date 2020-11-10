package pipeline

import (
	"fmt"
	"log"
	"og/db"
	req "og/reqeuest"
	"og/response"
	"og/setting"
	"og/spider"
)

func (self *Pipeline) Process(resp *response.Response) {
	log.Printf("[pipeline] %s status code : %d\n", resp.Req.URL, resp.StatusCode)

	// 1. 解析中（1. responsn->request 2. new page）
	request := self.process(resp)
	// 2. new request -> chan
	for _, r := range request {
		self.sendReq(r)
		fmt.Println(r)
	}
	// 3. save Crawler RST
	self.saveResponse(resp)
	// 4. handler old request: old -> chan ; save old
	r := self.handleReq(resp)
	self.sendReq(r)
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

func (self *Pipeline) toItemReq(resp *response.Response) []*req.Request {
	// 1. 选择当前需要处理的item
	// 2. 提取补全item -> [][]item.itme
	// -- 需要处理（分裂css， 直接css）-> 函数处理
	// -- 不需要处理
	// 3. []item.item to request

	return []*req.Request{resp.Req}
}

func (self *Pipeline) toPageReq(resp *response.Response) []*req.Request {
	return []*req.Request{}
}

func (self *Pipeline) saveResponse(resp *response.Response) {
	rst := req.ToCrawlerRst(resp.Req)
	rst["req_id"] = resp.Req.URL
	tablename := spider.FindKey(setting.TableName, spider.Zhaotoubiao()).Value
	self.db.Save(tablename, rst)
}

func (self *Pipeline) sendReq(r *req.Request) {
	self.scheduler <- r
	self.saveReq(r)
}

func (self *Pipeline) saveReq(req *req.Request) {
	if req.Seed {
		self.db.MustUpdate(req, true)
	} else {
		self.db.MustUpdate(req, false)
	}
}

// 处理req
func (self *Pipeline) handleReq(resp *response.Response) *req.Request {
	if resp.StatusCode == 200 {
		resp.Req.Status = req.StatusSuc
	} else {
		if resp.Req.Retry == 6 {
			resp.Req.Status = req.StatusFail
		} else {
			resp.Req.Status = req.StatusRetry
			resp.Req.Retry = 1 + resp.Req.Retry
		}
	}
	return resp.Req
}

func (self *Pipeline) process(resp *response.Response) []*req.Request {

	request := make([]*req.Request, 0)
	// 解析response
	request = append(request, self.toItemReq(resp)...)
	request = append(request, self.toPageReq(resp)...)
	return request

}
