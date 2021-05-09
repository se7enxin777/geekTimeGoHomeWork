package main

import (
	"context"
	"errors"
	"fmt"
	"go-project-layout/server/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	serverQuitchan := make(chan struct{})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverQuitchan <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	g.Go(func() error {
		select {
			case <-ctx.Done():
				fmt.Println("errgroup quit")
			case <-serverQuitchan:
				fmt.Println("server quit")
		}

		timeoutCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)


		fmt.Println("shutdown server")
		return server.Shutdown(timeoutCtx)
	})

	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
			case <-ctx.Done():
				return ctx.Err()
			case sig := <-quit:
				return errors.Errorf("os signal: %v", sig)
		}
	})

	fmt.Printf("errgroup quit: %+v\n", g.Wait())
}

