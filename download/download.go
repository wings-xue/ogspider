package download

import (
	"log"
	req "og/reqeuest"
	"og/response"
	"time"
)

func (self *Download) Process(r *req.Request) {

	go self.download(r)

}

type Download struct {
	pipeliner chan *response.Response
}

func New(pipeliner chan *response.Response) *Download {
	return &Download{
		pipeliner: pipeliner,
	}
}

func (self *Download) download(r *req.Request) {
	time.Sleep(time.Second * 1)
	log.Printf("[downloader] download url, %s", r.URL)
	self.pipeliner <- &response.Response{}

}
