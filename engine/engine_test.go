package engine

import (
	req "og/reqeuest"
	"og/response"
)

func newTestReqChan() chan req.Request {
	req := make(chan req.Request)
	return req
}

func newTestResChan() chan response.Response {
	response := make(chan response.Response)
	return response
}

// // debug 执行
// func TestRun(t *testing.T) {
// 	type args struct {
// 		scheduler  chan req.Request
// 		downloader chan req.Request
// 		pipeliner  chan response.Response
// 	}

// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			args: args{
// 				scheduler:  newTestReqChan(),
// 				downloader: newTestReqChan(),
// 				pipeliner:  newTestResChan(),
// 			},
// 		},
// 	}
// 	// 向chan添加数据

// 	for _, tt := range tests {
// 		go func() {
// 			tt.args.scheduler <- req.Request{}
// 			tt.args.downloader <- req.Request{}
// 			tt.args.pipeliner <- response.Response{}
// 		}()

// 		// 防止陷入死循环
// 		go func() {
// 			time.Sleep(1 * time.Microsecond)
// 			os.Exit(1)
// 		}()
// 		Run(tt.args.scheduler, tt.args.downloader, tt.args.pipeliner)

// 	}
// }
