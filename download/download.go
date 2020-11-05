package download

import (
	"log"
	req "og/reqeuest"
	"og/response"
)

func (d *Download) Process(*req.Request) *response.Response {
	log.Println("download process")
	return &response.Response{}
}

type Download struct {
	pipeliner chan *response.Response
	scheduler chan *req.Request
}

func New(scheduler chan *req.Request, pipeliner chan *response.Response) *Download {
	return &Download{
		scheduler: scheduler,
		pipeliner: pipeliner,
	}
}

// 运行n个下载器
func (s *Download) Run() {
	log.Println("调度reqeust给engine")
}
