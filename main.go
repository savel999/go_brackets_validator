package main

/**
https://www.sohamkamani.com/golang/2018-06-17-golang-using-context-cancellation/ - context cancel
https://medium.com/@pinkudebnath/graceful-shutdown-of-golang-servers-using-context-and-os-signals-cc1fa2c55e97 - graceful shutdown
https://www.alexedwards.net/blog/making-and-using-middleware - middlewares
*/

import (
	"context"
	"go_brackets_validator/controller"
	"go_brackets_validator/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func serve(ctx context.Context) error {
	mux := http.NewServeMux()
	c := &controller.BracketsController{}

	mux.Handle("/validate", middleware.IsNotPostMethodMiddleware(http.HandlerFunc(c.ValidateAction)))
	mux.Handle("/fix", middleware.IsNotPostMethodMiddleware(http.HandlerFunc(c.FixAction)))

	handler := middleware.LoggerMiddleware(mux)
	handler = middleware.PanicRecoveryMiddleware(handler)
	handler = middleware.HeadersMiddleware(handler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("server started")

	<-ctx.Done()

	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	var err error

	if err = srv.Shutdown(ctxShutDown); err != nil {
		if err == http.ErrServerClosed {
			log.Printf("server exited properly")
			err = nil
		} else {
			log.Fatalf("server Shutdown Failed:%+s", err)
		}
	}

	return err
}

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	defer close(signalChan)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-signalChan
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	if err := serve(ctx); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}
}
