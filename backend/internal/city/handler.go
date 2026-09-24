package city

import (
	"encoding/json"
	"net/http"
	"strings"
)

type CityHandler struct {
	service *CityService
}

func NewCityHandler(service *CityService) *CityHandler {
	return &CityHandler{
		service: service,
	}
}

type CountResponse struct {
	Letter string `json:"letter"`
	Count  int    `json:"count"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (handler *CityHandler) Count(w http.ResponseWriter, r *http.Request) {
	letter := strings.TrimSpace(r.URL.Query().Get("letter"))

	w.Header().Set("Content-Type", "application/json")

	if len([]rune(letter)) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "letter must contain exactly one character",
		})
		return
	}

	runeLetter := []rune(letter)[0]

	count, err := handler.service.CountByLetter(r.Context(), runeLetter)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "failed to count cities",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CountResponse{
		Letter: letter,
		Count:  count,
	})
}
