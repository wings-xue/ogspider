package og

import (
	"og/db"
	"og/download"
	"og/engine"
	"og/pipeline"
	req "og/reqeuest"
	"og/response"
	"og/schedule"
	scrape "og/scraper"
	"og/spider"
	"time"
)

func CreateReqTable() {

}

// LoopDispatchDB 自动从数据库提取失效的种子
func LoopDispatchDB(db *db.PgSQL, scheduler chan *req.Request) {
	for {
		for _, r := range db.SelectExpired() {
			// log.Printf("从数据库获取请求")
			// 重新开始读取数据
			r.Retry = 1
			scheduler <- r
		}
		time.Sleep(time.Minute * 10)
	}
}

// Crawl 爬虫情况通过数据库进行缓存, 每次基于缓存进行调用
func Crawl(spider ...*spider.BaseSpider) {

	database := db.New()

	scheduler := make(chan *req.Request)
	downloader := make(chan *req.Request)
	scraper := make(chan *response.Response)
	pipeliner := make(chan *response.Response)

	scrape := scrape.OpenSpider(pipeliner, scheduler, database, spider...)
	schedule := schedule.OpenSpider(downloader, database)
	pipeline := pipeline.OpenSpider(scrape.Setting, scheduler, database)
	dwonload := download.OpenSpider(scrape.Setting, scraper)
	engine := engine.OpenSpider(scheduler, downloader, pipeliner, scraper)

	go LoopDispatchDB(database, scheduler)
	go schedule.LoopDispatch()
	scrape.RunSpider(database)
	engine.RunForever(schedule, dwonload, pipeline, scrape)

}

// CrawlRPC 基于RPC调用，不影响之前爬虫爬取
func CrawlRPC(spider ...spider.BaseSpider) {
	// schedule.OpenSpiderRPC()
}
