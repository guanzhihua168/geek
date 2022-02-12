package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	s := &http.Server{
		Addr:    ":8081",
		Handler: http.DefaultServeMux,
	}

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting down server...")

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")

}
