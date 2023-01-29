// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Args的第一个元素：os.Args[0]，是命令本身的名字
	fmt.Println("Args len:", len(os.Args), os.Args)

	s, sep := "", ""
	//s[m:n]形式的切片表达式，产生从第m个元素到第n-1个元素的切片
	//省略切片表达式的m或n，会默认传入0或len(s)
	//os.Args[1:len(os.Args)] 等效于 os.Args[1:]
	//空标识符（blank identifier），即 _
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
