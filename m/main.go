package main

import (
	"context"
	"fmt"
	"time"
)

func tree() {
	ctx1 := context.Background()
	// ctx2, _ := context.WithCancel(ctx1)
	ctx3, _ := context.WithTimeout(ctx1, time.Second*5)
	ctx4, _ := context.WithTimeout(ctx3, time.Second*3)
	ctx5, _ := context.WithTimeout(ctx3, time.Second*6)
	ctx6 := context.WithValue(ctx5, "userID", 12)

	go watch(ctx1, "ctx1")
	// go watch(ctx2, "ctx2")
	go watch(ctx3, "ctx3")
	go watch(ctx4, "ctx4")
	go watch(ctx5, "ctx5")
	go watch(ctx6, "ctx6")

	select {
	case <-ctx4.Done():
		fmt.Println("exit")
	}
	// cancel()
}

func main() {
	// 创建一个子节点的context,3秒后自动超时
	tree()

}

func find(ctx context.Context) string {

	return "string"
}

// 单独的监控协程
func watch(ctx context.Context, name string) {
	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		fmt.Println(name, "收到信号，监控退出,time=", time.Now().Unix())
	// 		return
	// 	default:
	// 		fmt.Println(name, "goroutine监控中,time=", time.Now().Unix())
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }
}
