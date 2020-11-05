package pipeline

import (
	"log"
	"og/response"
)

func Process(*response.Response) {
	log.Println("[pipeline] pipeline run..")
}

type Pipeline struct {
}

func New() *Pipeline {
	return &Pipeline{}
}
func (s *Pipeline) Run() {
	log.Println("调度reqeust给engine")
}
