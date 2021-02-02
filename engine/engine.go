package engine

import (
	"og/db"
	"og/download"
	"og/response"
	scrape "og/scraper"

	"og/pipeline"
	req "og/reqeuest"
	"og/schedule"
)

type Engine struct {
	scheduler  chan *req.Request
	downloader chan *req.Request
	pipeliner  chan *response.Response
	scraper    chan *response.Response
}

func OpenSpider(
	scheduler chan *req.Request,
	downloader chan *req.Request,
	pipeliner chan *response.Response,
	scraper chan *response.Response,
) *Engine {
	return &Engine{
		scheduler:  scheduler,
		downloader: downloader,
		pipeliner:  pipeliner,
		scraper:    scraper,
	}
}

func (e *Engine) PushReq(r *req.Request) {

}

// Run: engine 运行
func (e *Engine) Run(
	database *db.PgSQL,
	scheduler *schedule.Schedule,
	download *download.Download,
	pipeline *pipeline.Pipeline,
) {

	// 初始化scheduler

	for {
		select {
		case req := <-e.scheduler:
			scheduler.Process(req)
		case req := <-e.downloader:
			go download.Process(req)
		case resp := <-e.pipeliner:
			go pipeline.Process(resp)
			// default:
			// time.Sleep(3 * time.Second)
			// log.Println("engine process")
		}

	}
}

func (e *Engine) RunForever(
	schedule *schedule.Schedule,
	download *download.Download,
	pipeline *pipeline.Pipeline,
	scrape *scrape.Scrape,
) {
	for {

		select {
		case req := <-e.scheduler:

			schedule.Process(req)
		case req := <-e.downloader:
			go download.Process(req)
		case resp := <-e.pipeliner:
			go pipeline.Process(resp)
		case resp := <-e.scraper:
			go scrape.Process(resp)
			// default:
			// time.Sleep(3 * time.Second)
			// log.Println("engine process")
		}

	}

}
