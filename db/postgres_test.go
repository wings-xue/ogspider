package db

// import (
// 	req "og/reqeuest"
// 	"og/spider"
// 	"reflect"
// 	"testing"
// )

// func TestPgSQL_Update(t *testing.T) {
// 	type args struct {
// 		r      *req.Request
// 		upsert bool
// 	}
// 	request := spider.ToSpider(spider.Zhaotoubiao())
// 	r := request[0]
// 	r1 := *r
// 	r1.Retry = 1

// 	r2 := *r
// 	r2.Retry = 2

// 	tests := []struct {
// 		name string
// 		self *PgSQL
// 		args args
// 		want int
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "init data",
// 			self: New(),
// 			args: args{
// 				r:      r,
// 				upsert: true,
// 			},
// 			want: 0,
// 		},
// 		{
// 			name: "change retry",
// 			self: New(),
// 			args: args{
// 				r:      &r1,
// 				upsert: true,
// 			},
// 			want: 1,
// 		},
// 		{
// 			name: "no change retry",
// 			self: New(),
// 			args: args{
// 				r:      &r2,
// 				upsert: false,
// 			},
// 			want: 1,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.self.Update(tt.args.r, tt.args.upsert)
// 			retry := tt.self.Select(10)[0].Retry
// 			if !reflect.DeepEqual(retry, tt.want) {
// 				t.Logf("err, %s", tt.name)
// 			}

// 		})
// 	}
// }
