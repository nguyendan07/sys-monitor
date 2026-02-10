# Real-time System Monitor (Go + Svelte)

A lightweight, high-performance web application designed to monitor server resources (CPU, RAM, Disk, Network) in real-time. This project leverages the concurrency power of **Golang** and the reactive efficiency of **Svelte**.

## 🚀 Key Features

* **Real-time Streaming:** Low-latency data transmission using WebSockets.
* **Low Overhead:** Efficient metric collection using `gopsutil`.
* **Reactive UI:** Fast and smooth data visualization with Svelte.
* **Modular Architecture:** Clean separation between the Go backend and Svelte frontend.

## 🛠 Tech Stack

* **Backend:** Golang, Gorilla WebSocket, gopsutil.
* **Frontend:** Svelte, Tailwind CSS, Chart.js (or LayerCake).
* **Communication:** JSON over WebSockets.

## 📂 Project Structure

```text
sys-monitor/
├── backend/    # Go API & WebSocket Hub
├── frontend/   # Svelte Web Dashboard
├── docker-compose.yml
└── README.md

```

## 🏃 Quick Start

### Prerequisites

* Go 1.20+
* Node.js (with npm or pnpm)

### Setup Backend

1. Navigate to the backend folder: `cd backend`
2. Install dependencies: `go mod tidy`
3. Start the server: `go run cmd/server/main.go`

### Setup Frontend

1. Navigate to the frontend folder: `cd frontend`
2. Install dependencies: `npm install`
3. Start the development server: `npm run dev`
