// Echo4 prints its command-line arguments.

package main

import (
	"flag"
	"fmt"
	"strings"
)

//包含了两个可选的命令行参数：

// -n用于忽略行尾的换行符
var n *bool = flag.Bool("n", false, "omit trailing newline")

// -s sep用于指定分隔字符（默认是空格）
var sep *string = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
