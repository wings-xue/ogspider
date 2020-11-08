package pipeline

import (
	"log"
	"og/response"
)

func (self *Pipeline) Process(*response.Response) {
	log.Println("[pipeline] pipeline run..")
}

type Pipeline struct {
}

func New() *Pipeline {
	return &Pipeline{}
}
