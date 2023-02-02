package main

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
//           go build -o fetch main.go            //
////////////////////////////////////////////////////

////////////////////////////////////////////////////
// <===================USERAGE===================>//
//           ./fetch http://bad.gopl.io           //
//           ./fetch http://gopl.io               //
////////////////////////////////////////////////////

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		bytes, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", bytes)
	}
}
