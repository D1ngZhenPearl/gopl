// Boiling prints the boiling point of water.

package main

import "fmt"

// 包一级范围声明语句声明
// 可在整个包对应的每个源文件中访问
const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
