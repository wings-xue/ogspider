package scrape

import (
	"og/db"
	req "og/reqeuest"
	"og/response"
	"og/setting"
	"og/spider"
	"time"
)

type Scrape struct {
	Spiders   []*spider.BaseSpider
	Pipeliner chan *response.Response
	Scheduler chan *req.Request
	// 所有的爬虫经过的配置
	Setting setting.CrawlerSet
	db      *db.PgSQL
}

// 注册爬虫
func (scrape *Scrape) Register(spider ...*spider.BaseSpider) {
	// 1. spider添加到scrape
	scrape.Spiders = append(scrape.Spiders, spider...)

	// 2. setting
	crawlSet := scrape.Setting
	for _, s := range spider {
		for key, val := range s.Setting.SpiderloadMiddleware {
			crawlSet.SpiderloadMiddleware[key] = val
		}
	}

	for _, s := range spider {
		for key, val := range s.Setting.DownloadMiddleware {
			crawlSet.DownloadMiddleware[key] = val
		}
	}
	for _, s := range spider {
		for key, val := range s.Setting.SpiderParse {
			crawlSet.SpiderParse[key] = val
		}
	}
	for _, s := range spider {
		for key, val := range s.Setting.PipelineSetting {
			crawlSet.PipelineSetting[key] = val
		}
	}
}

func New(db *db.PgSQL, pipeliner chan *response.Response, scheduler chan *req.Request) *Scrape {
	return &Scrape{
		db:        db,
		Pipeliner: pipeliner,
		Scheduler: scheduler,
		Setting:   setting.New(),
	}
}

func OpenSpider(
	pipeliner chan *response.Response,
	scheduler chan *req.Request,
	db *db.PgSQL,
	spider ...*spider.BaseSpider) *Scrape {

	s := New(db, pipeliner, scheduler)
	if len(spider) == 0 {
		panic("请确定是否配置爬虫")
	}
	s.Register(spider...)

	for _, spider := range s.Spiders {
		spider.CheckSpider()
		spider.CreateTable(db)
		for _, startReq := range spider.StartRequest() {
			go s.sendReq(startReq)
		}
	}
	return s
}

func (scrape *Scrape) ProcessMiddle(resp *response.Response) {
	for key, middleware := range scrape.Setting.SpiderloadMiddleware {
		if resp.MatchBool(key) {
			for _, hook := range middleware {
				resp = hook.Hook(resp)
			}
		}
	}
}

func (scrape *Scrape) ProcessParse(resp *response.Response) []*req.Request {
	out := make([]*req.Request, 0)
	for key, parses := range scrape.Setting.SpiderParse {
		if resp.MatchBool(key) {
			for _, hook := range parses {
				out = append(out, hook.Hook(resp)...)
			}
		}
	}
	return out
}

func (scrape *Scrape) handleReq(resp *response.Response) *req.Request {
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
	resp.Req.UpdateDate = time.Now()
	return resp.Req
}

func (scrape *Scrape) saveReq(req *req.Request) {
	if req.Seed {
		scrape.db.MustUpdate(*req, true)
	} else {
		scrape.db.MustUpdate(*req, false)
	}
}

func (scrape *Scrape) sendReq(r *req.Request) {
	scrape.Scheduler <- r
	scrape.saveReq(r)
}

func (scrape *Scrape) sendResp(resp *response.Response) {
	scrape.Pipeliner <- resp
}

func (scrape *Scrape) process(resp *response.Response) []*req.Request {
	// 1. baseprocess
	// 2. parse
	request := make([]*req.Request, 0)
	// 解析response
	request = append(request, resp.Extract()...)
	request = append(request, scrape.ProcessParse(resp)...)

	return request

}

func (scrape *Scrape) Process(resp *response.Response) {
	scrape.ProcessMiddle(resp)
	r := scrape.handleReq(resp)
	if resp.StatusCode != 200 {
		r.Log = resp.StatusMsg
		scrape.sendReq(r)
		return
	}

	// 1. 解析中（1. responsn->request 2. new page）
	request := scrape.process(resp)
	resp.NewReq = request
	scrape.sendResp(resp)
	// 1. 经过spidermiddle 过滤response
	// 2. 处理response
	// 3. 如果response成功，通过parse解析函数，response传入pipeline
	// 1. 如果response失败，初始失败的原始数据扔给scheduler

}
