package download

import (
	"context"
	"fmt"
	"log"
	req "og/reqeuest"
	"og/response"
	"os"
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	requ "github.com/imroc/req"
)

func (self *Download) Process(r *req.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
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
	browser   *rod.Browser
	mx        sync.Mutex
	Headless  bool
}

const (
	Chrome = `C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe`
)

func AddChrome() {
	os.Setenv("rob", "bin="+Chrome)
}

func New(pipeliner chan *response.Response) *Download {

	AddChrome()
	return &Download{
		pipeliner: pipeliner,
		Headless:  true,
	}
}

func (self *Download) SetHeadless(t bool) {
	self.Headless = t
}

func (self *Download) Empty() bool {
	self.mx.Lock()
	defer self.mx.Unlock()
	if self.browser == nil {
		return true
	}
	return false
}

func (self *Download) Require() {
	self.mx.Lock()
	defer self.mx.Unlock()
	if self.browser == nil {
		if !self.Headless {
			url, _ := launcher.New().Bin(Chrome).Headless(self.Headless).Launch()
			b := rod.New().ControlURL(url).MustConnect()
			self.browser = b
		} else {
			b := rod.New().MustConnect()
			self.browser = b
		}
	}
}

func (self *Download) download(ctx context.Context, r *req.Request) *response.Response {
	log.Printf("[Download] fetcher url: %s\n", r.URL)
	// 开启headless
	self.SetHeadless(false)

	resp := self.pageDownload(ctx, r)
	return resp

}

func (self *Download) httpDownload(ctx context.Context, r *req.Request) *response.Response {

	resp := response.New(r)

	res, _ := requ.Get(r.URL)
	content, _ := res.ToString()
	resp.Page = content
	return resp

}

func (self *Download) pageDownload(ctx context.Context, r *req.Request) *response.Response {
	resp := response.New(r)

	if self.Empty() {
		self.Require()
	}
	page := self.browser.MustPage("")
	defer page.Close()
	// disable alert
	page.MustEvalOnNewDocument(`window.alert = () => {}`)

	var e proto.NetworkResponseReceived

	wait := page.WaitEvent(&e)
	navErr := page.Timeout(10 * time.Second).Navigate(r.URL)
	wait()
	page.WaitLoad()

	if navErr != nil {
		return resp
	}

	ele, err1 := page.Timeout(10 * time.Second).Element("html")
	if err1 != nil {
		resp.StatusCode = 501
		return resp
	}
	s, err := ele.HTML()
	if err != nil {
		resp.StatusCode = 502
	}
	resp.Page = s
	fmt.Println(resp.Page)
	resp.StatusCode = 200
	return resp

}
