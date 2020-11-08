package engine

import (
	"og/db"
	"og/download"
	"og/response"

	"og/pipeline"
	req "og/reqeuest"
	"og/schedule"
	"og/spider"
)

type Engine struct {
	scheduler  chan *req.Request
	downloader chan *req.Request
	pipeliner  chan *response.Response
}

func New(scheduler chan *req.Request, downloader chan *req.Request, pipeliner chan *response.Response) *Engine {
	return &Engine{
		scheduler:  scheduler,
		downloader: downloader,
		pipeliner:  pipeliner,
	}
}

func (e *Engine) PushReq(r *req.Request) {

}

// Run: engine 运行
func (e *Engine) Run() {
	// 初始化数据库
	database := db.New()

	// 初始化spider
	Spider := spider.New(e.scheduler, database)
	go Spider.Run()

	// 初始化scheduler
	Schedule := schedule.New(e.downloader)
	go Schedule.Run()

	Download := download.New(e.pipeliner)

	Pipeline := pipeline.New()
	for {
		select {
		case req := <-e.scheduler:
			Schedule.Process(req)
		case req := <-e.downloader:
			Download.Process(req)
		case resp := <-e.pipeliner:
			Pipeline.Process(resp)
			// default:
			// time.Sleep(3 * time.Second)
			// log.Println("engine process")
		}

	}
}
