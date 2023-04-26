package main

import (
	"context"
	"csv-server/domain"
	"csv-server/handler"
	"csv-server/mapper"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	host = "localhost"
	port = "8080"
	src  = "ueba.csv"
)

func main() {
	recs, err := mapper.Load(src)
	if err != nil {
		log.Fatal(err)
	}

	rs := domain.NewRecordService(recs)
	http.HandleFunc("/api/v1/get-items", handler.ServeRecords(rs))

	srv := http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: http.DefaultServeMux,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("server started on host: %v, port: %v\n", host, port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
