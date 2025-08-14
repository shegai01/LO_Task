package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	internal "github.com/shegai01/LO_task/internal/handlers"
	"github.com/shegai01/LO_task/internal/logger"
	"github.com/shegai01/LO_task/internal/storage"
)

func main() {
	port := flag.Int("port", 8080, "port to run the server")
	flag.Parse()

	log := logger.NewAsyncLog(100)
	defer log.Close()
	log.Info("Server starting on port " + strconv.Itoa(*port))

	db := storage.NewStorage()
	server := &http.Server{
		Addr:              ":" + strconv.Itoa(*port),
		Handler:           internal.NewTimerHandler(db, log),
		ReadHeaderTimeout: time.Second,
	}

	var wg sync.WaitGroup

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err := server.ListenAndServe(); err != nil {
			return
		}
	}()

	<-signalChan

	if err := server.Shutdown(context.Background()); err != nil {
		log.Info("server.Shotdown: %v")
	}

	wg.Wait()
}
