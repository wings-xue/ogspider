package download

import (
	"context"
	"log"
	req "og/reqeuest"
	"og/response"
	"time"
)

func (self *Download) Process(r *req.Request) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*11)
	done := make(chan *response.Response)
	go func() {
		done <- self.download(ctx, r)
	}()
	select {
	case <-ctx.Done():
		self.pipeliner <- response.New(r)
	case resp := <-done:
		self.pipeliner <- resp
	}
}

type Download struct {
	pipeliner chan *response.Response
}

func New(pipeliner chan *response.Response) *Download {
	return &Download{
		pipeliner: pipeliner,
	}
}

func (self *Download) download(ctx context.Context, r *req.Request) *response.Response {

	resp := response.New(r)

	log.Printf("[Download] fetcher url: %s\n", r.URL)
	time.Sleep(time.Second * 10)
	resp.StatusCode = 200
	return resp

}
