package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"gocloud.dev/server/health"

	"gocloud.dev/server"
	"gocloud.dev/server/requestlog"
)

func main() {
	healthCheck := new(customHealthCheck)
	time.AfterFunc(5*time.Second, func() {
		healthCheck.mu.Lock()
		defer healthCheck.mu.Unlock()
		healthCheck.healthy = true
	})

	// Add request logger
	srvOptions := &server.Options{
		RequestLogger: requestlog.NewStackdriverLogger(os.Stdout, func(error) {}),
		HealthChecks:  []health.Checker{healthCheck},
	}
	srv := server.New(http.DefaultServeMux, srvOptions)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	if err := srv.ListenAndServe(":8080"); err != nil {
		log.Fatalf("%v", err)
	}
}

type customHealthCheck struct {
	mu      sync.RWMutex
	healthy bool
}

func (h *customHealthCheck) CheckHealth() error {
	fmt.Println("CheckHealth")
	h.mu.RLock()
	defer h.mu.RUnlock()
	if !h.healthy {
		return errors.New("not ready yet")
	}
	return nil
}
