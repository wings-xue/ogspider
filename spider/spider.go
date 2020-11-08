package spider

import (
	"log"
	req "og/reqeuest"
	"strconv"
)

type Spider struct {
	scheduler chan *req.Request
}

func New(scheduler chan *req.Request) *Spider {
	return &Spider{
		scheduler: scheduler,
	}
}

func ReadDB() []*req.Request {
	return []*req.Request{req.New("")}
}

func (s *Spider) Run() {
	log.Println("1. 读取数据库job")
	log.Println("2. 解析job为request")
	log.Println("3. request存入engine")

	i := 0
	for {
		i++
		url := strconv.Itoa(i)
		s.scheduler <- req.New(url)
		// time.Sleep(time.Second * 2)
	}

}
