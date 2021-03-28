package main

/**
https://www.sohamkamani.com/golang/2018-06-17-golang-using-context-cancellation/ - context cancel
*/

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		//fmt.Println(r.Method)
		// This prints to STDOUT to show that processing has started
		fmt.Println(os.Stdout, "processing request\n")
		// We use `select` to execute a peice of code depending on which
		// channel receives a message first
		select {
		case <-time.After(2 * time.Second):
			//w.Write([]byte("request processed"))
		case <-ctx.Done():
			fmt.Fprint(os.Stderr, "request cancelled: "+r.URL.Path)
			return
		}

		w.Write([]byte("request processed"))
	})

	fmt.Println("Server start")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
