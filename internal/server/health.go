package server

import (
	"encoding/json"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var (
	ready    atomic.Bool
	lastScan atomic.Value // time.Time
	lastSync atomic.Value // time.Time
)

// writeJSON encodes the response as JSON and logs any encoding errors.
func writeJSON(w http.ResponseWriter, v any) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("failed to encode JSON response: %v", err)
	}
}

// SetReady marks the agent as ready to receive traffic.
func SetReady(r bool) {
	ready.Store(r)
}

// RecordScan records the time of the last successful scan.
func RecordScan() {
	lastScan.Store(time.Now())
}

// RecordSync records the time of the last successful sync.
func RecordSync() {
	lastSync.Store(time.Now())
}

// GetLastScan returns the time of the last successful scan.
func GetLastScan() (time.Time, bool) {
	v := lastScan.Load()
	if v == nil {
		return time.Time{}, false
	}
	t, ok := v.(time.Time)
	return t, ok
}

// GetLastSync returns the time of the last successful sync.
func GetLastSync() (time.Time, bool) {
	v := lastSync.Load()
	if v == nil {
		return time.Time{}, false
	}
	t, ok := v.(time.Time)
	return t, ok
}

// healthzHandler returns a simple health check response.
// This is a basic liveness check that always returns OK if the server is running.
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	writeJSON(w, map[string]any{
		"status":    "ok",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// readyzHandler returns whether the agent is ready to process work.
// Returns 503 if the agent is still initializing.
func readyzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if !ready.Load() {
		w.WriteHeader(http.StatusServiceUnavailable)
		writeJSON(w, map[string]any{
			"status":    "not ready",
			"reason":    "agent initializing",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	writeJSON(w, map[string]any{
		"status":    "ready",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// livezHandler returns whether the agent is alive and functioning.
// Returns 503 if no scans have occurred in the last 10 minutes.
func livezHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ls, ok := lastScan.Load().(time.Time)
	if !ok || time.Since(ls) > 10*time.Minute {
		w.WriteHeader(http.StatusServiceUnavailable)
		response := map[string]any{
			"status":    "unhealthy",
			"reason":    "no recent scans",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		}
		if ok {
			response["last_scan"] = ls.Format(time.RFC3339)
		}
		writeJSON(w, response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]any{
		"status":    "healthy",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"last_scan": ls.Format(time.RFC3339),
	}
	if lsy, ok := lastSync.Load().(time.Time); ok {
		response["last_sync"] = lsy.Format(time.RFC3339)
	}
	writeJSON(w, response)
}
