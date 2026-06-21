// Package health provides liveness and readiness health check handlers.
package health

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
)

// Checker validates a dependency (database, cache, etc.).
type Checker interface {
	Name() string
	Check(ctx context.Context) error
}

// Handler serves /health and /ready endpoints.
type Handler struct {
	mu       sync.RWMutex
	checkers []Checker
}

// NewHandler creates a health handler.
func NewHandler(checkers ...Checker) *Handler {
	return &Handler{checkers: checkers}
}

// Register adds a readiness checker.
func (h *Handler) Register(c Checker) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.checkers = append(h.checkers, c)
}

type response struct {
	Status string            `json:"status"`
	Checks map[string]string `json:"checks,omitempty"`
}

// Live handles liveness probes — process is running.
func (h *Handler) Live(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, response{Status: "ok"})
}

// Ready handles readiness probes — dependencies are available.
func (h *Handler) Ready(w http.ResponseWriter, r *http.Request) {
	h.mu.RLock()
	checkers := h.checkers
	h.mu.RUnlock()

	checks := make(map[string]string, len(checkers))
	allOK := true
	for _, c := range checkers {
		if err := c.Check(r.Context()); err != nil {
			checks[c.Name()] = err.Error()
			allOK = false
		} else {
			checks[c.Name()] = "ok"
		}
	}

	status := http.StatusOK
	st := "ok"
	if !allOK {
		status = http.StatusServiceUnavailable
		st = "degraded"
	}
	writeJSON(w, status, response{Status: st, Checks: checks})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
