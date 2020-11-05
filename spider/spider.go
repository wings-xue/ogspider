package spider

import (
	"log"
	"og/job"
)

func ReadDB() []*job.Job {
	return []*job.Job{
		job.New(),
	}
}

func Run() {
	log.Println("1. 读取数据库job")
	log.Println("2. 解析job为request")
	log.Println("3. request存入engine")

}
