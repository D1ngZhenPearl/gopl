//uniq:
//$cat kt.txt
//$uniq kt.txt

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Please input something and press Enter to continue")
	fmt.Println("Press Control+D to exit the loop - EOF")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// line := input.Text()
		// counts[line] = counts[line] + 1
		counts[input.Text()]++
	}
	fmt.Println("Dup Summary:")
	for k, v := range counts {
		if v > 1 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}
}
