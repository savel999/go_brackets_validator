package main

/**
https://www.sohamkamani.com/golang/2018-06-17-golang-using-context-cancellation/ - context cancel
*/

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func someAction() {
	randNum := rand.Int63n(1)
	time.Sleep(time.Duration(randNum) * time.Millisecond)
}

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/validate", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// This prints to STDOUT to show that processing has started
		fmt.Println(os.Stdout, "processing request\n")
		// We use `select` to execute a peice of code depending on which
		// channel receives a message first
		select {
		case <-time.After(2 * time.Second):
			w.Write([]byte("request processed"))
		case <-ctx.Done():
			// If the request gets cancelled, log it
			// to STDERR
			fmt.Fprint(os.Stderr, "request cancelled\n")
		}
	}))
	fmt.Println("Server start")
	log.Fatal(http.ListenAndServe(":8080", handler))

}

func controller() {

}
