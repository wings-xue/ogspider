package download

import (
	"context"
	"og/middle"
	req "og/reqeuest"
	"og/response"
	"og/setting"
	"reflect"
	"testing"
)

func TestDownload_pageDownload(t *testing.T) {
	type args struct {
		ctx context.Context
		r   *req.Request
	}
	pipeliner := make(chan *response.Response)
	d := New(pipeliner)
	tests := []struct {
		name string
		self *Download
		args args
		want *response.Response
	}{
		// TODO: Add test cases.
		{
			self: d,
			args: args{
				r: req.New("http://www.baidu.com"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.self.pageDownload(tt.args.ctx, tt.args.r)
			t.Log(got)
			got2 := tt.self.httpDownload(tt.args.ctx, tt.args.r)
			t.Log(got2)
		})
	}
}

func TestDownload_ProcessMiddle(t *testing.T) {
	type args struct {
		ctx context.Context
		r   *req.Request
	}
	m1 := &middle.CookieDownloadMiddle{
		Cookie: "m1",
	}

	m2 := &middle.CookieDownloadMiddle{
		Cookie: "m2",
	}
	m3 := &middle.CookieDownloadMiddle{
		Cookie: "m3",
	}
	m := make(map[string][]middle.DownloadMiddle)
	m["*"] = []middle.DownloadMiddle{m1, m2, m3}
	scrape := make(chan *response.Response)
	d := OpenSpider(setting.CralwerSet{
		DownloadMiddleware: m,
	}, scrape)

	tests := []struct {
		name string
		self *Download
		args args
	}{
		// TODO: Add test cases.
		{
			self: d,
			args: args{
				r: req.New("http://www.baidu.com"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.self.ProcessMiddle(tt.args.ctx, tt.args.r)
			if !reflect.DeepEqual(tt.args.r.Cookie, m3.Cookie) {
				t.Log(tt.args.r.Cookie)
			}

		})
	}
}
