package main

// 练习 1.8： 修改fetch这个范例，如果输入的url参数没有 http:// 前缀的话，
// 为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。

// EX: 顺手支持了https的情况

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
//          go build -o fetch2 main2.go           //
// <====================TEST====================> //
//          go build -o fetch2 main2.go           //
////////////////////////////////////////////////////

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	httpPrefix  = "http://"
	httpsPrefix = "https://"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, httpPrefix) && !strings.HasPrefix(url, httpsPrefix) {
			url = httpPrefix + url
		}
		fmt.Printf("Ready to fetch: %s\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//练习 1.9： 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。
		fmt.Printf("HTTP STATUS: %d\n", resp.StatusCode)
		w, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copy in %d: %v", w, err)
		}
	}
}
