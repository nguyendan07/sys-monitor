package main

import (
	"log"
	"net/http"
	"time"

	"github.com/nguyendan07/sys-monitor/internal/collector"
	"github.com/nguyendan07/sys-monitor/internal/hub"
)

func main() {
	// 1. Initialize the Hub to manage WebSocket connections
	systemHub := hub.NewHub()

	go systemHub.Run()

	// 2. Initialize Goroutine to retrieve data from the system (Collector)
	go func ()  {
		log.Println("Metric Collector stated...")
		for {
			metrics, err := collector.GetMetrics()
			if err != nil {
				log.Printf("Error collecting metrics: %v", err)
				continue
			}

			systemHub.Broadcast <- metrics

			time.Sleep(2 * time.Second)
		}
	}()

	// 3. Set up Routes for the HTTP Server
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.ServeWs(systemHub, w, r)
	})

	// 4. Initialize the server
	port := ":8080"
	log.Printf("Backend server starting on http://localhost%s\n", port)
	log.Printf("WebSocket endpoint: ws://localhost:%s/ws", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
