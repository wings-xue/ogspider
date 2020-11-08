package db

import (
	req "og/reqeuest"
	"og/spider"
	"testing"
)

func TestPgSQL_Update(t *testing.T) {
	type args struct {
		r      *req.Request
		upsert bool
	}
	request := spider.ToSpider(spider.Zhaotoubiao())
	tests := []struct {
		name string
		self *PgSQL
		args args
	}{
		// TODO: Add test cases.
		{
			self: New(),
			args: args{
				r:      request[0],
				upsert: true,
			},
		},
		{
			self: New(),
			args: args{
				r:      request[0],
				upsert: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.self.Update(tt.args.r, tt.args.upsert)
		})
	}
}
