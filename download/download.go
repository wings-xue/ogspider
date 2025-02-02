package download

import (
	"context"
	"log"
	req "og/reqeuest"
	"og/response"
	"og/setting"
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	requ "github.com/imroc/req"
)

func (self *Download) Process(r *req.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	done := make(chan *response.Response)
	// 开启headless

	go func() {

		done <- self.download(ctx, r)

	}()
	select {
	case <-ctx.Done():

		self.scraper <- response.NewFail(r)
	case resp := <-done:

		self.scraper <- resp
	}
}

type Download struct {
	scraper  chan *response.Response
	browser  *rod.Browser
	mx       sync.Mutex
	Headless bool
	Setting  setting.CrawlerSet
}

// const (
// 	Chrome = `C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe`
// )

// func addENV() {
// 	os.Setenv("rob", "bin="+Chrome)
// }

func New(scraper chan *response.Response) *Download {

	// addENV()
	return &Download{
		scraper:  scraper,
		Headless: setting.Headless,
	}
}

func (self *Download) FromCrawl() {

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
			url, err := launcher.New().
				// Proxy("127.0.0.1:28080").
				// Bin(Chrome).
				Set("--no-sandbox").
				Headless(self.Headless).
				Devtools(true).
				Launch()
			if err != nil {
				log.Panic(err)
			}
			b := rod.New().ControlURL(url).MustConnect()
			self.browser = b
		} else {
			url, err := launcher.New().
				// Proxy("127.0.0.1:28080").
				// Bin(Chrome).
				Set("--no-sandbox").
				Launch()
			if err != nil {
				log.Panic(err)
			}
			b := rod.New().ControlURL(url).MustConnect()
			self.browser = b
		}
	}

	_, err := self.browser.Page(proto.TargetCreateTarget{URL: ""})
	if err != nil {
		log.Panic(err)
	}
}

func (self *Download) ProcessMiddle(ctx context.Context, r *req.Request) {
	for key, middleware := range self.Setting.DownloadMiddleware {
		if r.MatchBool(key) {
			for _, hook := range middleware {
				r = hook.Hook(r)
			}
		}
	}
}

func (self *Download) download(ctx context.Context, r *req.Request) *response.Response {
	log.Printf("[Download] uuid: %s, fetcher url: %s, retry: %d\n", r.UUID, r.URL, r.Retry)

	self.ProcessMiddle(ctx, r)
	resp := self.pageDownload(ctx, r)
	return resp

}

func (self *Download) httpDownload(ctx context.Context, r *req.Request) *response.Response {

	resp := response.NewFail(r)

	res, _ := requ.Get(r.URL)
	content, _ := res.ToString()
	resp.Page = content
	return resp

}

func (self *Download) pageDownload(ctx context.Context, r *req.Request) *response.Response {
	resp := response.NewFail(r)

	if self.Empty() {
		self.Require()
	}
	page, err := self.browser.Page(proto.TargetCreateTarget{URL: ""})
	defer page.Close()
	if err != nil {
		log.Println(err.Error())
	}
	// disable alert
	if page == nil {
		return response.NewFail(r)
	}
	page.EvalOnNewDocument(`window.alert = () => {}`)

	var e proto.NetworkResponseReceived

	wait := page.WaitEvent(&e)
	navErr := page.Timeout(100 * time.Second).Navigate(r.URL)
	if navErr != nil {
		resp.StatusMsg = navErr.Error()
		return resp
	}
	wait()
	err = page.WaitLoad()
	if err != nil {
		return resp
	}

	ele, err1 := page.Timeout(100 * time.Second).Element("html")

	if err1 != nil {
		resp.StatusCode = 501
		return resp
	}
	s, err := ele.HTML()
	if err != nil {
		resp.StatusCode = 502
	}

	resp.Page = s
	resp.StatusCode = 200
	return resp

}

// func ErrorAccessBlock(html string) int {
// 	substr := "很抱歉，由于您访问的URL有可能对网站造成安全威胁，您的访问被阻断"
// 	if strings.Contains(html, substr) {
// 		return 405
// 	}
// 	return 200
// }

func OpenSpider(s setting.CrawlerSet, scraper chan *response.Response) *Download {
	return &Download{
		scraper:  scraper,
		Headless: setting.Headless,
		Setting:  s,
	}
}
