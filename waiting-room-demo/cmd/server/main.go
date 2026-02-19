package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jawaracloud/waiting-room-demo/internal/queue"
	"github.com/jawaracloud/waiting-room-demo/internal/storage"
	"github.com/jawaracloud/waiting-room-demo/pkg/models"
	"github.com/nats-io/nats.go"
)

func main() {
	// Config
	redisAddr := getEnv("REDIS_ADDR", "localhost:6379")
	natsURL := getEnv("NATS_URL", nats.DefaultURL)
	jwtSecret := getEnv("JWT_SECRET", "jawaracloud-secret")
	port := getEnv("PORT", "8080")

	// Storage
	store, err := storage.NewRedisStorage(redisAddr)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// NATS
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Printf("Warning: Failed to connect to NATS: %v", err)
	} else {
		defer nc.Close()
	}

	// Service
	svc := queue.NewService(store, jwtSecret)

	// Background Cleanup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go svc.RunCleanup(ctx, 30*time.Second, 60) // Cleanup after 60s of inactivity

	// Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/enqueue", func(w http.ResponseWriter, r *http.Request) {
		token, status, err := svc.Enqueue(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("X-Queue-Token", token)
		json.NewEncoder(w).Encode(status)
	})

	r.Post("/status", func(w http.ResponseWriter, r *http.Request) {
		var req models.HeartbeatRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		status, err := svc.CheckStatus(r.Context(), req.Token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		json.NewEncoder(w).Encode(status)
	})

	// Admin endpoint to allow more users
	r.Post("/admin/allow", func(w http.ResponseWriter, r *http.Request) {
		nStr := r.URL.Query().Get("n")
		n, _ := strconv.ParseInt(nStr, 10, 64)
		if n == 0 {
			n = 10 // Default
		}

		ts, err := svc.AllowMore(r.Context(), n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if nc != nil {
			nc.Publish("waiting_room.capacity_increased", []byte(fmt.Sprintf("%d", ts)))
		}

		fmt.Fprintf(w, "Allowed up to timestamp %d", ts)
	})

	// Server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		log.Printf("Waiting Room Server starting on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")
	srv.Shutdown(ctx)
	log.Println("Server stopped")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
