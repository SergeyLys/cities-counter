package main

import (
	"log"
	"net/http"

	"github.com/sergeylys/city-counter/backend/internal/city"
)

func main() {
	cities := []city.City{
		{Name: "Rio de Janeiro"},
		{Name: "Cairo"},
		{Name: "Chongqing"},
		{Name: "Chengdu"},
	}

	cityRepository := city.NewMemoryCityRepository(cities)
	cityService := city.NewCityService(cityRepository)
	cityHandler := city.NewCityHandler(cityService)

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/cities/count", cityHandler.Count)

	log.Println("Server is running and listening on PORT :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
