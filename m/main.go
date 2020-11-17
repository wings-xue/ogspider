package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func addENV() {
	var Chrome = `C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe`
	os.Setenv("rob", "show,bin="+Chrome)

}

func main() {
	l := launcher.New().
		Headless(false).
		Devtools(true)

	defer l.Cleanup() // remove user-data-dir

	url := l.MustLaunch()

	// Trace shows verbose debug information for each action executed
	// Slowmotion is a debug related function that waits 2 seconds between
	// each action, making it easier to inspect what your code is doing.
	browser := rod.New().
		ControlURL(url)
	browser.Connect()
	// defer browser.Close()
	_, err := browser.Page(proto.TargetCreateTarget{URL: "http://www.baidu.com"})
	if err != nil {
		fmt.Println(err)
	}

	// page1.Close()
	time.Sleep(time.Second * 2)

	_, err = browser.Page(proto.TargetCreateTarget{URL: "http://www.baidu.com"})
	if err != nil {
		fmt.Println(err)
	}
	// page2.Close()
	time.Sleep(time.Second * 2)

	page3, err := browser.Page(proto.TargetCreateTarget{URL: "http://www.baidu.com"})
	if err != nil {
		fmt.Println(err)
	}
	page3.Close()
	time.Sleep(time.Second * 2)
}
