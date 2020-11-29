package scrape

import (
	"og/db"
	req "og/reqeuest"
	"og/response"
	"og/setting"
	"og/spider"
)

type Scrape struct {
	Spiders   []spider.OGSpider
	pipeliner chan *response.Response
	scheduler chan *req.Request
	Setting   setting.CralwerSet
}

// 注册爬虫
func (scrape *Scrape) Register(spider ...spider.OGSpider) {
	// scrape := make(chan *response.Response)

}

func New(pipeliner chan *response.Response, scheduler chan *req.Request) *Scrape {
	return &Scrape{}
}

// 注册出口chan
func (scrape *Scrape) SetPipeliner(pipeliner chan *response.Response) {
	// scrape := make(chan *response.Response)
	scrape.pipeliner = pipeliner
}

func (scrape *Scrape) SetScheduler(scheduler chan *req.Request) {
	// scrape := make(chan *response.Response)
	scrape.scheduler = scheduler
}

// func (scrape *Scrape) OpenSpider(db *db.PgSQL) {
// 	for _, spider := range scrape.spiders {
// 		spider.CheckSpider()
// 		spider.CreateTable(db)
// 		request := spider.StartRequest()
// 		log.Printf("[spider] init spider :%d", len(request))
// 		for _, r := range request {
// 			db.Update(r, true)
// 			scrape.scheduler <- r
// 		}

// 	}

// }

func OpenSpider(
	pipeliner chan *response.Response,
	scheduler chan *req.Request,
	db *db.PgSQL,
	spider ...spider.OGSpider) *Scrape {

	s := New(pipeliner, scheduler)
	if len(spider) == 0 {
		panic("请确定是否配置爬虫")
	}
	s.Register(spider...)

	for _, spider := range s.Spiders {
		spider.CheckSpider()
		spider.CreateTable(db)
		for _, startReq := range spider.StartRequest() {
			scheduler <- startReq
		}
	}

	return &Scrape{}
}

func (scrape *Scrape) Process() {

}
