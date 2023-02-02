// Server1 is a minimal "echo" server.

package main

////////////////////////////////////////////////////
// <====================RUN=====================> //
//           go build -o fetch main.go            //
////////////////////////////////////////////////////

// $ go build gopl.io/ch1/fetch
// $ ./fetch http://localhost:8000
// URL.Path = "/"
// $ ./fetch http://localhost:8000/help
// URL.Path = "/help"

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
