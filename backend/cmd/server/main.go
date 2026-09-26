package main

import (
	"log"
	"net/http"

	"github.com/sergeylys/city-counter/backend/internal/city"
	"github.com/sergeylys/city-counter/backend/internal/config"
)

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	cfg := config.Load()

	cityHandler := city.CityModule()

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/cities/count", cityHandler.Count)

	address := ":" + cfg.Port

	log.Printf("Server listening on %s", address)

	if err := http.ListenAndServe(address, cors(http.DefaultServeMux)); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
