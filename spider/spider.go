package spider

import (
	"log"
	"og/job"
	req "og/reqeuest"
	"time"
)

type Spider struct {
	scheduler chan *req.Request
}

func New(scheduler chan *req.Request) *Spider {
	return &Spider{
		scheduler: scheduler,
	}
}

func ReadDB() []*job.Job {
	return []*job.Job{
		job.New(),
	}
}

func (s *Spider) Run() {
	log.Println("1. 读取数据库job")
	log.Println("2. 解析job为request")
	log.Println("3. request存入engine")
	for {
		s.scheduler <- req.New("")
		time.Sleep(time.Second * 2)
	}

}
