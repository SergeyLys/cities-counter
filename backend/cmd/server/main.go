package main

import (
	"log"
	"net/http"

	"github.com/sergeylys/city-counter/backend/internal/city"
	"github.com/sergeylys/city-counter/backend/internal/config"
)

func main() {
	cfg := config.Load()

	cityHandler := city.CityModule()

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/cities/count", cityHandler.Count)

	address := ":" + cfg.Port

	log.Printf("Server listening on %s", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
