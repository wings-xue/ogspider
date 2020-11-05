package download

import (
	"log"
	req "og/reqeuest"
)

func Process(*req.Request) {
	log.Println("download process")
}

type Download struct {
}

func New() *Download {
	return &Download{}
}
func (s *Download) Run() {
	log.Println("调度reqeust给engine")
}
