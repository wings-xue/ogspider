package download

import (
	"context"
	req "og/reqeuest"
	"og/response"
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
