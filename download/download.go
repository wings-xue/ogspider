package download

import (
	"log"
	"og/context"
	req "og/reqeuest"
	"og/response"
	"sync"
)

func (self *Download) Process(r *req.Request) *response.Response {
	self.mu.Lock()
	defer self.mu.Unlock()
	log.Printf("[downloader] download process, %d", self.manager.Len())
	self.manager.Push(r)
	if self.manager.Len() >= self.Size {
		queue := self.manager.Queue
		self.manager.Free()
		for _, request := range queue {
			self.wg.Add(1)
			go self.download(request)
		}
		self.wg.Wait()
	}
	return &response.Response{}
}

type Download struct {
	pipeliner chan *response.Response
	Size      int
	manager   *context.Manager
	wg        sync.WaitGroup
	mu        sync.Mutex
}

func New(pipeliner chan *response.Response) *Download {
	return &Download{
		Size:      10,
		pipeliner: pipeliner,
		manager:   context.New(),
	}
}

// 运行n个下载器
func (self *Download) Run() {
	log.Println("调度reqeust给engine")
}

func (self *Download) download(r *req.Request) response.Response {
	defer self.wg.Done()
	log.Println("download正在请求数据....")
	return response.Response{}
}
