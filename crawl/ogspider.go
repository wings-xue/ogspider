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
)

// Crawl 爬虫情况通过数据库进行缓存, 每次基于缓存进行调用
func Crawl(spider ...spider.OGSpider) {

	database := db.New()

	scheduler := make(chan *req.Request)
	downloader := make(chan *req.Request)
	scraper := make(chan *response.Response)
	pipeliner := make(chan *response.Response)

	scrape := scrape.OpenSpider(pipeliner, scheduler, database, spider...)
	schedule.OpenSpider(scrape.Setting.Schedule, downloader, database)
	pipeline.OpenSpider(scrape.Setting.Pipeline, scheduler, database)
	download.OpenSpider(scrape.Setting.Download, scraper)

	engine.RunForever()
}

// CrawlRPC 基于RPC调用，不影响之前爬虫爬取
func CrawlRPC(spider ...spider.OGSpider) {
	// schedule.OpenSpiderRPC()
}
