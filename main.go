package main

/**
https://www.sohamkamani.com/golang/2018-06-17-golang-using-context-cancellation/ - context cancel
https://medium.com/@pinkudebnath/graceful-shutdown-of-golang-servers-using-context-and-os-signals-cc1fa2c55e97 - graceful shutdown
https://www.alexedwards.net/blog/making-and-using-middleware - middlewares
*/

import (
	"fmt"
	"go_brackets_validator/controller"
	"go_brackets_validator/middleware"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	c := &controller.BracketsController{}

	mux.Handle("/validate", middleware.IsNotPostMethodMiddleware(http.HandlerFunc(c.ValidateAction)))
	mux.Handle("/fix", middleware.IsNotPostMethodMiddleware(http.HandlerFunc(c.FixAction)))

	handler := middleware.LoggerMiddleware(mux)
	handler = middleware.PanicRecoveryMiddleware(handler)
	handler = middleware.HeadersMiddleware(handler)

	fmt.Println("Server start")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
