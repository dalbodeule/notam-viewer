package main

import (
	"log"
	"net/http"
	"os"
)

const defaultAddr = ":8080"

func main() {
	addr := getAddr()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("starting NOTAM backend on %s", addr)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}

func getAddr() string {
	if v := os.Getenv("NOTAM_BACKEND_ADDR"); v != "" {
		return v
	}

	if v := os.Getenv("PORT"); v != "" {
		return ":" + v
	}

	return defaultAddr
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
