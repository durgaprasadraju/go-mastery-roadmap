// Command server is the entry point for the Enterprise REST API project.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/config"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/health"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/shutdown"
	"go.uber.org/zap"
)

func main() {
	cfg := config.LoadBase()
	log := logger.Must("enterprise-api")

	mux := http.NewServeMux()
	healthHandler := health.NewHandler()

	mux.HandleFunc("GET /health", healthHandler.Live)
	mux.HandleFunc("GET /ready", healthHandler.Ready)
	mux.HandleFunc("GET /api/v1/users", listUsers)
	mux.HandleFunc("POST /api/v1/users", createUser)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.HTTPPort),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Info("server starting", zap.Int("port", cfg.HTTPPort))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("server error", zap.Error(err))
		}
	}()

	shutdown.Graceful(log, cfg.ShutdownTimeout, func(ctx context.Context) error {
		return srv.Shutdown(ctx)
	})
}

func listUsers(w http.ResponseWriter, _ *http.Request) {
	users := []map[string]string{
		{"id": "1", "name": "Alice", "email": "alice@example.com"},
		{"id": "2", "name": "Bob", "email": "bob@example.com"},
	}
	writeJSON(w, http.StatusOK, users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusCreated, map[string]string{
		"id": "3", "name": req.Name, "email": req.Email,
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
