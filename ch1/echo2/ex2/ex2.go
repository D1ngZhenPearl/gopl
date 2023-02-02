// 练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。
package main

import (
	"fmt"
	"os"
)

// %d: digit   (10進位的數字)
// %c: char    (字元)
// %s: string  (字串)
func main() {
	for index, arg := range os.Args[1:] {
		fmt.Printf("%d: %s\n", index, arg)
	}
}
