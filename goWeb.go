package main

import (
	"net/http"
	"fmt"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,`<h1>They there</h1>
<h1>They there</h1>
<h1>They there</h1>
<h1>They there</h1>`)
	fmt.Fprintf(w, "<h1>They there</h1>")
	fmt.Fprintf(w, "<center>They there</center>")
	fmt.Fprintf(w, "<center>They %s how about %s, there is a lot to eat.</center>", "<strong> Ahmed</strong>","eating")
}
func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)

}
